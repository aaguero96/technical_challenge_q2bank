package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/aaguero96/technical_challenge_q2bank/events/consumer/handler"
	"github.com/aaguero96/technical_challenge_q2bank/events/producer"
	"github.com/aaguero96/technical_challenge_q2bank/external_API/validator"
	"github.com/aaguero96/technical_challenge_q2bank/initializers"
	"github.com/aaguero96/technical_challenge_q2bank/repository/transaction_repository"
	"github.com/aaguero96/technical_challenge_q2bank/repository/user_repository"
	"github.com/aaguero96/technical_challenge_q2bank/repository/user_type_repository"
	"github.com/aaguero96/technical_challenge_q2bank/repository/wallet_repository"
	"github.com/aaguero96/technical_challenge_q2bank/service/external_validator_service"
	"github.com/aaguero96/technical_challenge_q2bank/service/transaction_service"
	"github.com/go-redis/redis/v7"
	uuid "github.com/satori/go.uuid"
)

var ts transaction_service.TransactionService
var evs external_validator_service.ExternalValidatorService

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.ConnectRedisClient()
	initializers.CreateConsumerGroup()
}

var (
	waitGroup    sync.WaitGroup
	consumerName string = uuid.NewV4().String()
)

func processStream(stream redis.XMessage, retry bool) {
	defer waitGroup.Done()

	typeEvent := stream.Values["type"].(string)
	dataBin := []byte(stream.Values["data"].(string))

	if typeEvent == "approving_transaction" {
		var data producer.TransactionEvent
		if err := data.UnmarshalBinary(dataBin); err != nil {
			fmt.Printf("error on unmarshal stream: %v \n", stream.ID)
			return
		}
		handler.ApprovingTransactionHandler(data, evs, ts)
	}

	if typeEvent == "canceling_transaction" {
		var data producer.TransactionCancelEvent
		if err := data.UnmarshalBinary(dataBin); err != nil {
			fmt.Printf("error on unmarshal stream: %v \n", stream.ID)
			return
		}
		handler.CancelingTransactionHandler(data, ts)
	}

	initializers.Client.XAck(os.Getenv("STREAM_REDIS_NAME"), os.Getenv("CONSUMER_GROUP_REDIS_NAME"), stream.ID)
}

func consumeEvents() {
	for {
		func() {
			streams, err := initializers.Client.XReadGroup(&redis.XReadGroupArgs{
				Streams:  []string{os.Getenv("STREAM_REDIS_NAME"), ">"},
				Group:    os.Getenv("CONSUMER_GROUP_REDIS_NAME"),
				Consumer: consumerName,
				Count:    10,
				Block:    0,
			}).Result()

			if err != nil {
				fmt.Printf("err on consume events: %v \n", err)
				return
			}

			for _, stream := range streams[0].Messages {
				waitGroup.Add(1)
				go processStream(stream, false)
			}
		}()
	}
}

// func consumePendingEvents() {
// 	ticker := time.Tick(time.Second * 30)

// 	for {
// 		select {
// 		case <- ticker:
// 			func () {

// 			}
// 		}
// 	}
// }

func main() {
	// Start Repositories
	userRepository := user_repository.NewUserRepository(initializers.DB)
	transactionRepository := transaction_repository.NewTransactionRepository(initializers.DB)
	walletRepository := wallet_repository.NewWalletRepository(initializers.DB)
	userTypeRepository := user_type_repository.NewUserTypeRepository(initializers.DB)

	// Start Services
	ts = transaction_service.NewTransactionService(transactionRepository, &userRepository, &walletRepository, userTypeRepository)
	evs = external_validator_service.NewExternalValidatorService(validator.NewValidatorExternalAPI())

	go consumeEvents()

	chanOS := make(chan os.Signal, 1)
	signal.Notify(chanOS, syscall.SIGINT, syscall.SIGTERM)
	<-chanOS

	waitGroup.Wait()
	initializers.Client.Close()
}

package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/aaguero96/technical_challenge_q2bank/events/consumer/handler"
	"github.com/aaguero96/technical_challenge_q2bank/events/producer"
	"github.com/aaguero96/technical_challenge_q2bank/externalAPI/validator"
	"github.com/aaguero96/technical_challenge_q2bank/initializers"
	"github.com/aaguero96/technical_challenge_q2bank/repository/transactionRepository"
	"github.com/aaguero96/technical_challenge_q2bank/repository/userRepository"
	"github.com/aaguero96/technical_challenge_q2bank/repository/walletRepository"
	"github.com/aaguero96/technical_challenge_q2bank/service/externalValidatorService"
	"github.com/aaguero96/technical_challenge_q2bank/service/transactionService"
	"github.com/go-redis/redis/v7"
	uuid "github.com/satori/go.uuid"
)

var ts transactionService.TransactionService
var evs externalValidatorService.ExternalValidatorService

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
	userRepository := userRepository.NewUserRepository(initializers.DB)
	transactionRepository := transactionRepository.NewTransactionRepository(initializers.DB)
	walletRepository := walletRepository.NewWalletRepository(initializers.DB)

	// Start Services
	ts = transactionService.NewTransactionService(transactionRepository, userRepository, walletRepository)
	evs = externalValidatorService.NewExternalValidatorService(validator.NewValidatorExternalAPI())

	go consumeEvents()

	chanOS := make(chan os.Signal, 1)
	signal.Notify(chanOS, syscall.SIGINT, syscall.SIGTERM)
	<-chanOS

	waitGroup.Wait()
	initializers.Client.Close()
}

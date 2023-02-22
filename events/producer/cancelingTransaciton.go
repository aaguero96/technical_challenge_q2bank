package producer

import (
	"os"

	"github.com/aaguero96/technical_challenge_q2bank/initializers"
	"github.com/go-redis/redis/v7"
	"github.com/vmihailenco/msgpack"
)

type TransactionCancelEvent struct {
	TransactionID int
	PayerID       int
	PayeeID       int
	Amount        float64
	Status        string
}

func (tce *TransactionCancelEvent) MarshalBinary() (data []byte, err error) {
	return msgpack.Marshal(tce)
}

func (tce *TransactionCancelEvent) UnmarshalBinary(data []byte) error {
	return msgpack.Unmarshal(data, tce)
}

func CancelingTransactionEvent(transaction TransactionCancelEvent) error {
	_, err := initializers.Client.XAdd(&redis.XAddArgs{
		Stream: os.Getenv("STREAM_REDIS_NAME"),
		Values: map[string]interface{}{
			"type": "canceling_transaction",
			"data": &transaction,
		},
	}).Result()
	return err
}

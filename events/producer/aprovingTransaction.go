package producer

import (
	"os"

	"github.com/aaguero96/technical_challenge_q2bank/initializers"
	"github.com/go-redis/redis/v7"
	"github.com/vmihailenco/msgpack"
)

type TransactionEvent struct {
	PayerID int
	PayeeID int
	Amount  float64
	Status  string
}

func (te *TransactionEvent) MarshalBinary() (data []byte, err error) {
	return msgpack.Marshal(te)
}

func (te *TransactionEvent) UnmarshalBinary(data []byte) error {
	return msgpack.Unmarshal(data, te)
}

func ApprovingTransactionEvent(transaction TransactionEvent) error {
	_, err := initializers.Client.XAdd(&redis.XAddArgs{
		Stream: os.Getenv("STREAM_REDIS_NAME"),
		Values: map[string]interface{}{
			"type": "approving_transaction",
			"data": &transaction,
		},
	}).Result()
	return err
}

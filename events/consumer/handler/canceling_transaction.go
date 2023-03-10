package handler

import (
	"fmt"

	"github.com/aaguero96/technical_challenge_q2bank/events/producer"
	"github.com/aaguero96/technical_challenge_q2bank/service/transaction_service"
)

func CancelingTransactionHandler(
	data producer.TransactionCancelEvent,
	ts transaction_service.TransactionService,
) error {
	fmt.Println("processing...")
	if data.Status == "canceled" {
		fmt.Println("transaction was canceled")
		return nil
	}
	err := ts.Return(data.PayerID, data.PayeeID, data.TransactionID, data.Amount, data.Status)
	if err != nil {
		fmt.Printf("error when canceling transaction with id %v \n", data.TransactionID)
		return err
	}
	fmt.Printf("It was returned the amount of %v, from %v to %v \n", data.Amount, data.PayeeID, data.PayerID)
	return nil
}

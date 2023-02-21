package handler

import (
	"fmt"

	"github.com/aaguero96/technical_challenge_q2bank/events/producer"
	"github.com/aaguero96/technical_challenge_q2bank/service/externalValidatorService"
	"github.com/aaguero96/technical_challenge_q2bank/service/transactionService"
)

func ApprovingTransactionHandler(
	data producer.TransactionEvent,
	evs externalValidatorService.ExternalValidatorService,
	ts transactionService.TransactionService,
) error {
	reponse, err := evs.Validator()
	if err != nil {
		return err
	}
	if reponse {
		err = ts.Deposit(data.PayerID, data.PayeeID, data.TransactionID, data.Amount)
		fmt.Println("processing...")
		if err != nil {
			fmt.Printf("error in transaction with id %v \n", data.TransactionID)
			return err
		}
		fmt.Printf("It was transfered the amount of %v, from %v to %v \n", data.Amount, data.PayerID, data.PayeeID)
		return nil
	}
	fmt.Println("Have problems in aprovment")
	return nil
}

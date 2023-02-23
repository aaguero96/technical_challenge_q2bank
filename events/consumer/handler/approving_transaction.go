package handler

import (
	"errors"
	"fmt"

	"github.com/aaguero96/technical_challenge_q2bank/events/producer"
	"github.com/aaguero96/technical_challenge_q2bank/service/external_validator_service"
	"github.com/aaguero96/technical_challenge_q2bank/service/transaction_service"
)

func ApprovingTransactionHandler(
	data producer.TransactionEvent,
	evs external_validator_service.ExternalValidatorService,
	ts transaction_service.TransactionService,
) error {
	reponse, err := evs.Validator()
	if err != nil {
		return err
	}

	transaction, err := ts.GetById(data.TransactionID)
	if err != nil {
		return err
	}
	fmt.Println("processing...")
	if transaction.Status != "in progress" {
		return errors.New("transaction already be canceled")
	}
	if reponse {
		err = ts.Deposit(data.PayerID, data.PayeeID, data.TransactionID, data.Amount)
		if err != nil {
			fmt.Printf("error in transaction with id %v \n", data.TransactionID)
			return err
		}
		fmt.Printf("It was transfered the amount of %v, from %v to %v \n", data.Amount, data.PayerID, data.PayeeID)
		return nil
	}
	err = ts.DenyTransfer(data.TransactionID)
	if err != nil {
		fmt.Printf("error in transaction with id %v \n", data.TransactionID)
		return err
	}
	fmt.Println("Have problems in aprovment")
	return nil
}

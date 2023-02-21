package handler

import (
	"fmt"

	"github.com/aaguero96/technical_challenge_q2bank/events/producer"
	"github.com/aaguero96/technical_challenge_q2bank/externalAPI/validator"
	"github.com/aaguero96/technical_challenge_q2bank/service/externalValidatorService"
)

func ApprovingTransactionHandler(
	data producer.TransactionEvent,
	externalValidatorAPI validator.ValidatorExternalAPI,
) error {
	evs := externalValidatorService.NewExternalValidatorService(externalValidatorAPI)
	reponse, err := evs.Validator()
	if err != nil {
		return err
	}
	if reponse {
		fmt.Printf("Foi transferido o valor de %v, de %v para %v \n", data.Amount, data.PayerID, data.PayeeID)
		return nil
	}
	fmt.Println("Transação não foi aprovada")
	return nil
}

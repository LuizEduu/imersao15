package usecase

import (
	"log"

	"github.com/luizeduu/imersao/codepix-go/domain/model"
)

type ConfirmTransactionUseCase struct {
	TransactionRepository model.TransactionRepositoryInterface
}

func (confirmTransactionUseCase *ConfirmTransactionUseCase) Execute(transactionId string) (*model.Transaction, error) {
	transaction, err := confirmTransactionUseCase.TransactionRepository.Find(transactionId)

	if err != nil {
		log.Println("Transaction not found", transactionId)
		return nil, err
	}

	transaction.Status = model.TransactionConfirmed

	err = confirmTransactionUseCase.TransactionRepository.Save(transaction)

	if err != nil {
		return nil, err
	}

	return transaction, nil

}

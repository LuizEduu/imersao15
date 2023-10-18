package usecase

import (
	"log"

	"github.com/luizeduu/imersao/codepix-go/domain/model"
)

type CompleteTransactionUseCase struct {
	TransactionRepository model.TransactionRepositoryInterface
}

func (completeTransactionUseCase *CompleteTransactionUseCase) Execute(transactionId string) (*model.Transaction, error) {
	transaction, err := completeTransactionUseCase.TransactionRepository.Find(transactionId)

	if err != nil {
		log.Println("Transaction not found", transactionId)
		return nil, err
	}

	transaction.Status = model.TransactionCompleted

	err = completeTransactionUseCase.TransactionRepository.Save(transaction)

	if err != nil {
		return nil, err
	}

	return transaction, nil

}

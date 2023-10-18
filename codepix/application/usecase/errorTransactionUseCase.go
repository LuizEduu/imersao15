package usecase

import "github.com/luizeduu/imersao/codepix-go/domain/model"

type ErrorTransactionUseCase struct {
	TransactionRepository model.TransactionRepositoryInterface
}

func (errorTransactionRepository *ErrorTransactionUseCase) Execute(transactionId string, reason string) (*model.Transaction, error) {
	transaction, err := errorTransactionRepository.TransactionRepository.Find(transactionId)

	if err != nil {
		return nil, err
	}

	transaction.Status = model.TransactionError
	transaction.CancelDescription = reason

	err = errorTransactionRepository.TransactionRepository.Save(transaction)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}

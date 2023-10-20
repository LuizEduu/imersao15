package usecase

import (
	"errors"

	"github.com/luizeduu/imersao/codepix-go/domain/model"
)

type RegisterTransactionUseCase struct {
	TransactionRepository model.TransactionRepositoryInterface
	PixKeyRepository      model.PixKeyRepositoryInterface
}

func (registerTransactionUseCase *RegisterTransactionUseCase) Execute(accountId string, amount float64, pixKeyTo string, pixKeyKindTo string, description string, id string) (*model.Transaction, error) {
	account, err := registerTransactionUseCase.PixKeyRepository.FindAccount(accountId)

	if err != nil {
		return nil, err
	}

	pixKey, err := registerTransactionUseCase.PixKeyRepository.FindKeyByKind(pixKeyKindTo, pixKeyKindTo)

	if err != nil {
		return nil, err
	}

	transaction, err := model.NewTransaction(account, amount, pixKey, description, id)

	if err != nil {
		return nil, err
	}

	registerTransactionUseCase.TransactionRepository.Save(transaction)

	if transaction.ID == "" {
		return nil, errors.New("unable to process this transaction")
	}

	return transaction, nil
}

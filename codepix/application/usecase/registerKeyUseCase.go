package usecase

import (
	"errors"

	"github.com/luizeduu/imersao/codepix-go/domain/model"
)

type RegisterKeyUseCase struct {
	PixKeyRepository model.PixKeyRepositoryInterface
}

func (registerKeyUseCase *RegisterKeyUseCase) Execute(key string, kind string, accountId string) (*model.PixKey, error) {
	account, err := registerKeyUseCase.PixKeyRepository.FindAccount(accountId)

	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(kind, account, key)

	if err != nil {
		return nil, err
	}

	registerKeyUseCase.PixKeyRepository.RegisterKey(pixKey)

	if pixKey.ID == "" {
		return nil, errors.New("unable to create new key at the moment")
	}

	return pixKey, nil

}

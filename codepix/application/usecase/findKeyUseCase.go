package usecase

import "github.com/luizeduu/imersao/codepix-go/domain/model"

type FindKeyUseCase struct {
	PixKeyRepository model.PixKeyRepositoryInterface
}

func (findKeyUseCase *FindKeyUseCase) Execute(key string, kind string) (*model.PixKey, error) {
	pixKey, err := findKeyUseCase.PixKeyRepository.FindKeyByKind(key, kind)

	if err != nil {
		return nil, err
	}

	return pixKey, err
}

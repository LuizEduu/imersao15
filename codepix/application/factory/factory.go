package factory

import (
	"github.com/jinzhu/gorm"
	"github.com/luizeduu/imersao/codepix-go/application/usecase"
	"github.com/luizeduu/imersao/codepix-go/infrastructure/repository"
)

func CompleteTransactionUseCaseFactory(database *gorm.DB) usecase.CompleteTransactionUseCase {
	transactionRepository := repository.TransactionRepositoryDb{Db: database}
	completeTransactionUseCase := usecase.CompleteTransactionUseCase{TransactionRepository: &transactionRepository}

	return completeTransactionUseCase
}

func ConfirmTransactionUseCaseFactory(database *gorm.DB) usecase.ConfirmTransactionUseCase {
	transactionRepository := repository.TransactionRepositoryDb{Db: database}
	confirmTransactionUseCase := usecase.ConfirmTransactionUseCase{TransactionRepository: &transactionRepository}

	return confirmTransactionUseCase
}
func ErrorTransactionUseCaseFactory(database *gorm.DB) usecase.ErrorTransactionUseCase {
	transactionRepository := repository.TransactionRepositoryDb{Db: database}
	errorTransactionUseCase := usecase.ErrorTransactionUseCase{TransactionRepository: &transactionRepository}

	return errorTransactionUseCase
}
func FindKeyUseCaseFactory(database *gorm.DB) usecase.FindKeyUseCase {
	pixRepository := repository.PixKeyRepositoryDb{Db: database}
	findKeyUseCase := usecase.FindKeyUseCase{PixKeyRepository: &pixRepository}

	return findKeyUseCase
}
func RegisterKeyUseCaseFactory(database *gorm.DB) usecase.RegisterKeyUseCase {
	pixRepository := repository.PixKeyRepositoryDb{Db: database}
	registerKeyUseCase := usecase.RegisterKeyUseCase{PixKeyRepository: &pixRepository}

	return registerKeyUseCase
}
func RegisterTransactionUseCaseFactory(database *gorm.DB) usecase.RegisterTransactionUseCase {
	transactionRepository := repository.TransactionRepositoryDb{Db: database}
	pixRepository := repository.PixKeyRepositoryDb{Db: database}

	RegisterTransactionUseCase := usecase.RegisterTransactionUseCase{
		TransactionRepository: &transactionRepository,
		PixKeyRepository:      &pixRepository}

	return RegisterTransactionUseCase
}

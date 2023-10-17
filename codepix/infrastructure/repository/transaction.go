package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/luizeduu/imersao/codepix-go/domain/model"
)

type TransactionRepositoryDb struct {
	Db *gorm.DB
}

func (repository *TransactionRepositoryDb) Register(tranction *model.Transaction) error {
	err := repository.Db.Create(tranction).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository *TransactionRepositoryDb) Save(transaction *model.Transaction) error {
	err := repository.Db.Save(transaction).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository *TransactionRepositoryDb) Find(id string) (*model.Transaction, error) {
	var transaction model.Transaction

	repository.Db.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("transaction not found")
	}

	return &transaction, nil
}

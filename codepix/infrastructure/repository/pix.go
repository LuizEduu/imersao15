package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/luizeduu/imersao/codepix-go/domain/model"
)

type PixKeyRepositoryDb struct {
	Db *gorm.DB
}

func (repository PixKeyRepositoryDb) AddBank(bank *model.Bank) error {
	err := repository.Db.Create(bank).Error

	if err != nil {
		return err
	}

	return nil
}

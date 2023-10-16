package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

type PixKeyRepositoryInterface interface {
	RegisterKey(pixKey *PixKey) (*PixKey, error)
	FindKeyByKind(key string, kind string) (*PixKey, error)
	AddBank(bank *Bank) error
	AddAccount(account *Account) error
	FindAccount(id string) (*Account, error)
}

type PixKey struct {
	Base      `valid:"required"`
	Kind      string   `json:"kind" valid:"notnull"`
	Key       string   `json:"key" valid:"notnull"`
	AccountID string   `json:"account_id" valid:"notnull"`
	Account   *Account `valid:"-"`
	Status    string   `json:"status" valid:"notnull"`
}

func (pixKey *PixKey) isValid() error {
	_, err := govalidator.ValidateStruct(pixKey)
	if err != nil {
		return err
	}

	err = pixKey.validateKindType()

	if err != nil {
		return err
	}

	err = pixKey.validateStatus()

	if err != nil {
		return err
	}

	return nil
}

func (pixKey *PixKey) validateKindType() error {
	if pixKey.Kind != "email" && pixKey.Kind != "cpf" {
		return errors.New("invalid type of key")
	}
	return nil
}

func (pixKey *PixKey) validateStatus() error {
	if pixKey.Status != "active" && pixKey.Status != "inactive" {
		return errors.New("invalid status")
	}
	return nil
}

func NewPixKey(kind string, key string, Account *Account) (*PixKey, error) {
	pixKey := PixKey{
		Kind:    kind,
		Key:     key,
		Account: Account,
		Status:  "active",
	}

	pixKey.ID = uuid.New().String()
	pixKey.CreatedAt = time.Now()

	err := pixKey.isValid()

	if err != nil {
		return nil, err
	}

	return &pixKey, nil
}

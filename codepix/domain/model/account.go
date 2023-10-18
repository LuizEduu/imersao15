package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/google/uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Account struct {
	ID        string    `json:"id" valid:"required"`
	OwnerName string    `gorm:"column:owner_name;type:varchar(255);not null" json:"owner_name" valid:"notnull"`
	Bank      *Bank     `valid:"-"`
	BankID    string    `gorm:"column:bank_id;type:uuid;not null" valid:"-"`
	Number    string    `json:"number" valid:"notnull"`
	PixKeys   []*PixKey `gorm:"ForeignKey: AccountID" valid:"-"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
}

func (account *Account) isValid() error {
	_, err := govalidator.ValidateStruct(account)

	if err != nil {
		return err
	}

	return nil
}

func NewAccount(bank *Bank, number string, ownerName string) (*Account, error) {
	account := Account{
		OwnerName: ownerName,
		Bank:      bank,
		BankID:    bank.ID,
		Number:    number,
	}

	account.ID = uuid.New().String()
	account.CreatedAt = time.Now()

	err := account.isValid()

	if err != nil {
		return nil, err
	}

	return &account, nil
}

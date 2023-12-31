package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/google/uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Bank struct {
	ID        string     `json:"id" valid:"required"`
	Code      string     `json:"code" gorm:"type:varchar(20)" valid:"notnull"`
	Name      string     `json:"name" gorm:"type:varchar(20)" valid:"notnull"`
	Accounts  []*Account `gorm:"ForeignKey:BankID" valid:"-"`
	CreatedAt time.Time  `json:"created_at" valid:"-"`
	UpdatedAt time.Time  `json:"updated_at" valid:"-"`
}

func (bank *Bank) isValid() error {
	_, err := govalidator.ValidateStruct(bank)

	if err != nil {
		return err
	}

	return nil
}

func NewBank(code string, name string) (*Bank, error) {
	bank := Bank{
		Code: code,
		Name: name,
	}

	bank.ID = uuid.New().String()
	bank.CreatedAt = time.Now()

	err := bank.isValid()

	if err != nil {
		return nil, err
	}

	return &bank, nil
}

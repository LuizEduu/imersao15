package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/google/uuid"
)

const (
	TransactionPending   string = "pending"
	TransactionCompleted string = "completed"
	TransactionError     string = "error"
	TransactionConfirmed string = "confirmed"
)

type Transactions struct {
	Transaction []*Transaction
}

type Transaction struct {
	Base              `valid:"required"`
	AccountFrom       *Account `valid:"-"`
	Amount            float64  `json:"amount" valid:"notnull"`
	PixKeyTo          *PixKey  `valid:"-"`
	Status            string   `json:"status" valid:"notnull"`
	Description       string   `json:"description" valid:"notnull"`
	CancelDescription string   `json:"cancel_description" valid:"notnull"`
}

func (transaction *Transaction) validateAmount() error {
	if transaction.Amount <= 0 {
		return errors.New("the amount must be greater than 0")
	}

	return nil
}

func (transaction *Transaction) validateStatus() error {

	status := []string{TransactionPending, TransactionCompleted, TransactionError, TransactionConfirmed}
	isContain := false

	for _, s := range status {
		if s == transaction.Status {
			isContain = true
		}
	}

	if !isContain {
		return errors.New("invalid status of transaction")
	}

	return nil
}

func (transaction *Transaction) validatePixKeyToIsEqualFrom() error {

	if transaction.AccountFrom.ID == transaction.PixKeyTo.AccountID {
		return errors.New("the source and destination account cannot be the same")
	}

	return nil
}

func (transaction *Transaction) isValid() error {
	_, err := govalidator.ValidateStruct(transaction)
	if err != nil {
		return err
	}

	err = transaction.validateAmount()

	if err != nil {
		return err
	}

	err = transaction.validateStatus()

	if err != nil {
		return err
	}

	err = transaction.validatePixKeyToIsEqualFrom()

	if err != nil {
		return err
	}

	return nil
}

func NewTransaction(accountFrom *Account, amount float64, pixKeyTo *PixKey, description string) (*Transaction, error) {
	transaction := Transaction{
		AccountFrom: accountFrom,
		Amount:      amount,
		PixKeyTo:    pixKeyTo,
		Status:      TransactionPending,
		Description: description,
	}

	transaction.ID = uuid.New().String()
	transaction.CreatedAt = time.Now()

	err := transaction.isValid()

	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (transaction *Transaction) Complete() error {
	transaction.Status = TransactionCompleted
	transaction.UpdatedAt = time.Now()

	err := transaction.isValid()

	return err
}

func (transaction *Transaction) Confirm() error {
	transaction.Status = TransactionConfirmed
	transaction.UpdatedAt = time.Now()

	err := transaction.isValid()

	return err
}

func (transaction *Transaction) Cancel(description string) error {
	transaction.Status = TransactionError
	transaction.UpdatedAt = time.Now()
	transaction.CancelDescription = description

	err := transaction.isValid()

	return err
}

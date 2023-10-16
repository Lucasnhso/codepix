package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

const (
	TransactionPending string = "pending"
	TransactionCompleted string = "completed"
	TransactionError string = "error"
	TransctionConfirmed string = "confirmed"
)

type TransactionRepositoryInterface interface {
	Register(transaction *Transaction) error
	Save(transaction *Transaction) error
	Find(id string) (*Transaction, error)
}

type Transactions struct {
	Transaction []Transaction
}

type Transaction struct {
	Base `valid:"required"`
	AccountFrom *Account `valid:"-"`
	Amount float64 `json:"amount" valid:"notnull"`
	PixKeyTo *PixKey `valid:"-"`
	Status string `json:"status" valid:"notnull"`
	Description string `json:"description" valid:"notnull"`
	CancelDescription string `json:"cancel_description" valid:"-"`
}

func (t *Transaction) isValid() error {
	_, err := govalidator.ValidateStruct(t)
	
	if t.Amount <= 0 {
		return errors.New("The amount must be greater than 0")
	}
	
	if t.Status!= TransactionPending && t.Status != TransactionCompleted && t.Status != TransactionError {
		return errors.New("invalid status")
	}

	if t.PixKeyTo.AccountId == t.AccountFrom.ID {
		return errors.New("the souce and destination account cannot be the same")
	}

	if err!= nil {
    return err
  }
	return nil
}

func NewTransaction(AccountFrom *Account, amount float64, pixKeyTo *PixKey, description string) (*Transaction, error) {
	transaction := Transaction{
		AccountFrom: AccountFrom,
    Amount: amount,
    PixKeyTo: pixKeyTo,
    Status: TransactionPending,
    Description: description,
	}

	transaction.ID = string(uuid.NewV4().String())
	transaction.CreatedAt = time.Now()

	err:= transaction.isValid()

	if err!= nil {
    return nil, err
  }

	return &transaction, nil
}

func (t *Transaction) Complete() error {
	t.Status = TransactionCompleted
  t.UpdatedAt = time.Now()
	err := t.isValid()
	return err
}
func (t *Transaction) Confirm() error {
	t.Status = TransctionConfirmed
  t.UpdatedAt = time.Now()
	err := t.isValid()
	return err
}

func (t *Transaction) Cancel(description string) error {
	t.Status = TransactionError
  t.UpdatedAt = time.Now()
	t.Description = description
	err := t.isValid()
	return err
}
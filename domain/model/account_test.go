package model_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"

	"github.com/Lucasnhso/codepix-go/domain/model"
	"github.com/stretchr/testify/require"
)

func TestModel_NewAccount(t *testing.T) {
	code := "001"
	name := "test bank"
	bank, err := model.NewBank(code, name)

	accountNumber := "abcnumber"
	ownerName := "Lucas"
	account, err := model.NewAccount(bank, accountNumber, ownerName)

	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil(account.ID))
	require.Equal(t, account.Number, accountNumber)
	require.Equal(t, account.Bank.ID, bank.ID)

	_, err = model.NewAccount(bank, "", ownerName)
	require.NotNil(t, err)
}
package db

import (
	"context"
	"database/sql"
	"github.com/simpleBank/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"

	_ "github.com/stretchr/testify"
)

// Create account
func CreateAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

// Test crate account
func TestCreateAccount(t *testing.T) {
	CreateAccount(t)
}

// Test get account
func TestGetAccount(t *testing.T) {
	account1 := CreateAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Currency, account2.Currency)
	require.Equal(t, account1.Balance, account2.Balance)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

// Test update account
//func TestUpdateAccount(t *testing.T) {
//	account1 := CreateAccount(t)
//
//	arg := AddAccountBalanceParams{
//		ID:     account1.ID,
//		Amount: util.RandomMoney(),
//	}
//
//	account2, err := testQueries.AddAccountBalance(context.Background(), arg)
//	require.NoError(t, err)
//	require.NotEmpty(t, account2)
//
//	require.Equal(t, account1.ID, account2.ID)
//	require.Equal(t, account1.Owner, account2.Owner)
//	require.Equal(t, arg.Amount, account2.Balance)
//	require.Equal(t, account1.Currency, account2.Currency)
//	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
//
//}

// Test delete account
func TestDeleteAccount(t *testing.T) {
	account1 := CreateAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.Equal(t, err.Error(), sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

// Test list account
func TestListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateAccount(t)
	}

	arg := ListAccountParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)

	}
}

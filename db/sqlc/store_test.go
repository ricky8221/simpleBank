package db

//import (
//	"context"
//	"fmt"
//	"github.com/stretchr/testify/require"
//	"testing"
//)
//
//func TestTransferTxDeadlock(t *testing.T) {
//	store := NewStore(testDB)
//
//	account1 := CreateAccount(t)
//	account2 := CreateAccount(t)
//	fmt.Println(">> before:", account1.Balance, account2.Balance)
//
//	// run c concurrent transfer transaction
//	n := 10
//	amount := int64(10)
//
//	errs := make(chan error)
//	results := make(chan TransferTxResult)
//	for i := 0; i < n; i++ {
//		formAccountId := account1.ID
//		toAccountId := account2.ID
//
//		if 1%2 == 1 {
//			formAccountId = account2.ID
//			toAccountId = account1.ID
//
//		}
//		go func() {
//			result, err := store.TransferTx(context.Background(), TransferTxParams{
//				FromAccountID: formAccountId,
//				ToAccountID:   toAccountId,
//				Amount:        amount,
//			})
//
//			errs <- err
//			results <- result
//		}()
//	}
//
//	// Validate result
//	for i := 0; i < n; i++ {
//		err := <-errs
//		require.NoError(t, err)
//	}
//
//	// check the final updated balances
//	updatedAccount1, err := store.GetAccount(context.Background(), account1.ID)
//	require.NoError(t, err)
//
//	updatedAccount2, err := store.GetAccount(context.Background(), account2.ID)
//	require.NoError(t, err)
//
//	fmt.Println(">> after:", account1.Balance, account2.Balance)
//
//	require.Equal(t, account1.Balance, updatedAccount1.Balance)
//	require.Equal(t, account2.Balance, updatedAccount2.Balance)
//
//}

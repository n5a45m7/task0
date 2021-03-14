package memory

import (
	"app/storage"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransaction(t *testing.T) {
	accID := 4
	st := NewTransactionStorage()

	// test Creation
	for i := 0; i < 10; i++ {
		_, err := st.Create(storage.CreateTransactionDTO{AccountID: accID, Amount: 100})
		assert.NoError(t, err)
	}

	// test correct amount value
	lastTransaction, ok, err := st.GetLastAccountTx(accID)
	assert.True(t, ok)
	assert.NoError(t, err)
	assert.Equal(t, 10.0*100, lastTransaction.AccountAmount)

	// test receive correct number of txs
	txs, err := st.GetByAccount(accID, storage.WithOffsetLimit(4, 4))
	assert.NoError(t, err)
	assert.Equal(t, 4, len(txs))

	// test receive correct number of txs 2
	// here correct number is 3 because we try to receive 4 txs from total 10 with offset 7
	// (not enough elements to return 4 txs)
	txs, err = st.GetByAccount(accID, storage.WithOffsetLimit(7, 4))
	assert.NoError(t, err)
	assert.Equal(t, 3, len(txs))

	// test receive tx from unexisted account
	txs, err = st.GetByAccount(5)
	assert.NoError(t, err)
	assert.Empty(t, txs)

	// test outbound limit offset
	txs, err = st.GetByAccount(accID, storage.WithOffsetLimit(100, 10))
	assert.NoError(t, err)
	assert.Empty(t, txs)

	// test outbound limit offset 2
	txs, err = st.GetByAccount(accID, storage.WithOffsetLimit(100, 10), storage.WithDir(storage.DIR_ASC))
	assert.NoError(t, err)
	assert.Empty(t, txs)

	// test negative amount
	_, err = st.Create(storage.CreateTransactionDTO{AccountID: accID, Amount: -10000})
	assert.Error(t, err)
	assert.Equal(t, storage.ErrStorTrCrNegativeAccountAmount, err)
}

func TestConcurency(t *testing.T) {
	accID := 4
	st := NewTransactionStorage()

	var wg sync.WaitGroup

	// lots of writes
	wg.Add(1)
	go func() {
		defer wg.Done()
		var wg0 sync.WaitGroup
		for i := 0; i < 10000; i++ {
			wg0.Add(1)
			go func() {
				defer wg0.Done()
				_, err := st.Create(storage.CreateTransactionDTO{AccountID: accID, Amount: 100})
				assert.NoError(t, err)
			}()
		}
		wg0.Wait()
	}()

	// lots of reads
	wg.Add(1)
	go func() {
		defer wg.Done()
		var wg0 sync.WaitGroup
		for i := 0; i < 10000; i++ {
			wg0.Add(1)
			go func() {
				defer wg0.Done()
				_, _, err := st.GetLastAccountTx(accID)
				assert.NoError(t, err)
			}()
		}
		wg0.Wait()
	}()

	wg.Wait()
}

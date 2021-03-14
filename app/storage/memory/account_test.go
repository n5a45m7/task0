package memory

import (
	"app/storage"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccount(t *testing.T) {
	st := NewAccountStorage()
	userID := 4
	accNum := 10

	for i := 0; i < accNum; i++ {
		_, err := st.Create(storage.CreateAccountDTO{CustomerID: userID})
		assert.NoError(t, err)
	}

	// correct number of accounts
	accounts, err := st.GetByUser(userID)
	assert.NoError(t, err)
	assert.Equal(t, accNum, len(accounts))

	// correct number of accounts for non existed user
	accounts, err = st.GetByUser(765)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(accounts))

}

package memory

import (
	"app/storage"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	st := NewUserStorage()
	entity, err := st.Create(storage.CreateUserDTO{Name: "Alexa", Surname: "Amazon"})
	assert.NoError(t, err)

	entity1, ok, err := st.GetUser(entity.ID)
	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, entity.Name, entity1.Name)
	assert.Equal(t, entity.Surname, entity1.Surname)

	_, ok, err = st.GetUser(765)
	assert.NoError(t, err)
	assert.False(t, ok)
}

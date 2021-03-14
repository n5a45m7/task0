package memory

import (
	"app"
	"app/storage"
	"sync"
)

type accountStorage struct {
	// [userID] => acc entity
	data map[int][]app.Account
	// id auto increment
	idInc int

	mu sync.RWMutex
}

func NewAccountStorage() interface {
	storage.AccountCreator
	storage.AccountReceiver
} {
	return &accountStorage{
		data:  make(map[int][]app.Account),
		idInc: 1,
	}
}

func (s *accountStorage) Create(dto storage.CreateAccountDTO) (app.Account, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	entity := app.Account{
		ID:         s.idInc,
		CustomerID: dto.CustomerID,
	}

	s.data[entity.CustomerID] = append(s.data[entity.CustomerID], entity)
	s.idInc++

	return entity, nil
}

func (s *accountStorage) GetByUser(userID int) ([]app.Account, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	entity := s.data[userID]
	return entity, nil
}

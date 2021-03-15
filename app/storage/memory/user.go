package memory

import (
	"app"
	"app/storage"
	"sync"
)

type userStorage struct {
	// [userID] => user entity
	data map[int]app.User
	// id auto increment
	idInc int

	mu sync.RWMutex
}

func NewUserStorage() interface {
	storage.UserCreator
	storage.UserReceiver
} {
	return &userStorage{
		data:  make(map[int]app.User),
		idInc: 1,
	}
}

func (s *userStorage) Create(dto storage.CreateUserDTO) (app.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	entity := app.User{
		ID:      s.idInc,
		Name:    dto.Name,
		Surname: dto.Surname,
	}
	s.data[entity.ID] = entity
	s.idInc++

	return entity, nil
}

func (s *userStorage) GetUser(userID int) (app.User, bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	entity, ok := s.data[userID]
	return entity, ok, nil
}

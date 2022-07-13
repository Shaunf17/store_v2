package store

import (
	"fmt"
	"store/auth"
	"sync"
)

type Store struct {
	Entities    map[string]Entity
	DataChannel chan []Entity
	sync.Mutex
}

var (
	ErrKeyAlreadyExists error
	ErrKeyDoesntExists  error
)

var store = NewStore()

func (s *Store) Add(key string, value string, user *auth.User) (string, error) {
	s.Lock()
	defer s.Unlock()
	if _, exists := s.Entities[key]; exists {
		ErrKeyAlreadyExists = fmt.Errorf("The key {%v} already exists", key)
		return "", ErrKeyAlreadyExists
	}

	s.Entities[key] = *NewEntity(key, value, user)

	return fmt.Sprintf("Key {%v} successfully added", key), nil
}

func (s *Store) Update(key string, value string, user *auth.User) (string, error) {
	s.Lock()
	defer s.Unlock()

	return fmt.Sprintf("Key {%v} successfully updated", key), nil
}

func (s *Store) Find(key string) (Entity, error) {
	s.Lock()
	defer s.Unlock()
	entity, exists := s.Entities[key]
	if !exists {
		ErrKeyDoesntExists = fmt.Errorf("The key {%v} does not exist", key)
		return Entity{}, ErrKeyDoesntExists
	}

	return entity, nil
}

func (s *Store) GetAll() ([]Entity, error) {
	s.Lock()
	defer s.Unlock()

	var entities []Entity
	for _, entity := range s.Entities {
		entities = append(entities, entity)
	}
	return entities, nil
}

func (s *Store) Delete(key string) (string, error) {
	s.Lock()
	defer s.Unlock()
	_, exists := s.Entities[key]
	if !exists {
		ErrKeyDoesntExists = fmt.Errorf("The key {%v} does not exist", key)
		return "", ErrKeyDoesntExists
	}

	delete(s.Entities, key)
	return fmt.Sprintf("Key {%v} successfully deleted", key), nil
}

func NewStore() *Store {
	s := Store{
		Entities: make(map[string]Entity),
	}
	return &s
}

func Connect() *Store {
	return store
}

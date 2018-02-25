package store

import (
	"github.com/MEDIGO/laika/models"
)

type MemoryStore struct {
	state *models.State
}

func NewMemoryStore(s *models.State) (*MemoryStore, error) {
	return &MemoryStore{
		state: s,
	}, nil
}

func (*MemoryStore) Persist(eventType string, data string) (int64, error) {
	panic("memory store does not support Persist()")
}

func (s *MemoryStore) State() (*models.State, error) {
	return s.state, nil
}

func (*MemoryStore) Migrate() error {
	return nil
}

func (*MemoryStore) Ping() error {
	return nil
}

func (*MemoryStore) Reset() error {
	return nil
}

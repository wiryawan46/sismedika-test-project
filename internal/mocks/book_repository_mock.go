package mocks

import (
	"sismedika-test-project/internal/book/infrastructure/storage"
	"sync"
)

// MockStore adalah mock dari store yang digunakan di repository
type MockStore struct {
	Books map[string]storage.BookData
	Mu    sync.RWMutex
}

// NewMockStore constructor untuk MockStore
func NewMockStore() *MockStore {
	return &MockStore{
		Books: make(map[string]storage.BookData),
		Mu:    sync.RWMutex{},
	}
}

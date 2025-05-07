// internal/book/infrastructure/bookstore.go
package storage

import (
	"sync"
)

// BookData adalah struktur penyimpanan internal buku (bukan domain).
type BookData struct {
	ID            string
	Title         string
	Author        string
	PublishedYear int
}

// BookStore adalah penyimpanan data buku di memory.
type BookStore struct {
	Books map[string]BookData
	Mu    sync.RWMutex
}

var (
	instance *BookStore
	once     sync.Once
)

// GetBookStore mengembalikan instance Singleton BookStore.
func GetBookStore() *BookStore {
	once.Do(func() {
		instance = &BookStore{
			Mu:    sync.RWMutex{},
			Books: make(map[string]BookData),
		}
	})
	return instance
}

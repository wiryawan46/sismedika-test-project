// internal/book/infrastructure/bookstore.go
package infrastructure

import (
	"errors"
	"sync"

	"sismedika-test-project/internal/book/domain"
)

// BookStore adalah penyimpanan data buku di memory.
type BookStore struct {
	mu    sync.RWMutex
	books map[string]domain.Book
}

var (
	instance *BookStore
	once     sync.Once
)

// GetBookStore mengembalikan instance Singleton BookStore.
func GetBookStore() *BookStore {
	once.Do(func() {
		instance = &BookStore{
			books: make(map[string]domain.Book),
		}
	})
	return instance
}

// GetAll mengembalikan semua buku.
func (s *BookStore) GetAll() ([]domain.Book, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	books := make([]domain.Book, 0, len(s.books))
	for _, b := range s.books {
		books = append(books, b)
	}
	return books, nil
}

// GetByID mengambil buku berdasarkan ID.
func (s *BookStore) GetByID(id string) (domain.Book, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	book, ok := s.books[id]
	if !ok {
		return domain.Book{}, errors.New("buku tidak ditemukan")
	}
	return book, nil
}

// Create menambahkan buku baru.
func (s *BookStore) Create(book domain.Book) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.books[book.ID]; exists {
		return errors.New("buku dengan ID ini sudah ada")
	}
	s.books[book.ID] = book
	return nil
}

// Update memperbarui data buku.
func (s *BookStore) Update(id string, book domain.Book) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.books[id]; !ok {
		return errors.New("buku tidak ditemukan")
	}
	book.ID = id // pastikan ID tetap
	s.books[id] = book
	return nil
}

// Delete menghapus buku.
func (s *BookStore) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.books[id]; !ok {
		return errors.New("buku tidak ditemukan")
	}
	delete(s.books, id)
	return nil
}

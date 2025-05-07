package repository

import (
	"context"
	"errors"
	"sismedika-test-project/internal/book/domain"
	"sismedika-test-project/internal/book/infrastructure/storage"
)

// bookRepository mengimplementasikan domain.BookRepository
type bookRepository struct {
	store *storage.BookStore
}

// NewBookRepository membuat instance baru dari bookRepository
func NewBookRepository(store *storage.BookStore) BookRepository {
	return &bookRepository{
		store: store,
	}
}

// GetAll mengembalikan semua buku dari penyimpanan
func (r *bookRepository) GetAll(ctx context.Context) ([]domain.Book, error) {
	r.store.Mu.RLock()
	defer r.store.Mu.RUnlock()

	var books []domain.Book
	for _, b := range r.store.Books {
		books = append(books, domain.Book{
			ID:            b.ID,
			Title:         b.Title,
			Author:        b.Author,
			PublishedYear: b.PublishedYear,
		})
	}
	return books, nil
}

// GetByID mengembalikan satu buku berdasarkan ID
func (r *bookRepository) GetByID(ctx context.Context, id string) (*domain.Book, error) {
	r.store.Mu.RLock()
	defer r.store.Mu.RUnlock()

	b, ok := r.store.Books[id]
	if !ok {
		return nil, errors.New("book not found")
	}
	return &domain.Book{
		ID:            b.ID,
		Title:         b.Title,
		Author:        b.Author,
		PublishedYear: b.PublishedYear,
	}, nil
}

// Create menambahkan buku baru ke penyimpanan
func (r *bookRepository) Create(ctx context.Context, book domain.Book) (*domain.Book, error) {
	r.store.Mu.Lock()
	defer r.store.Mu.Unlock()

	r.store.Books[book.ID] = storage.BookData{
		ID:            book.ID,
		Title:         book.Title,
		Author:        book.Author,
		PublishedYear: book.PublishedYear,
	}
	return &book, nil
}

// Update memperbarui informasi buku berdasarkan ID
func (r *bookRepository) Update(ctx context.Context, id string, book domain.Book) (*domain.Book, error) {
	r.store.Mu.Lock()
	defer r.store.Mu.Unlock()

	if _, ok := r.store.Books[id]; !ok {
		return nil, errors.New("book not found")
	}

	book.ID = id // pastikan ID tetap
	r.store.Books[id] = storage.BookData{
		ID:            id,
		Title:         book.Title,
		Author:        book.Author,
		PublishedYear: book.PublishedYear,
	}
	return &book, nil
}

// Delete menghapus buku berdasarkan ID
func (r *bookRepository) Delete(ctx context.Context, id string) error {
	r.store.Mu.Lock()
	defer r.store.Mu.Unlock()

	if _, ok := r.store.Books[id]; !ok {
		return errors.New("book not found")
	}
	delete(r.store.Books, id)
	return nil
}

package repository

import (
	"context"
	"sismedika-test-project/internal/book/domain"
)

// BookRepository adalah interface abstraksi untuk operasi data buku.
type BookRepository interface {
	GetAll(ctx context.Context) ([]domain.Book, error)
	GetByID(ctx context.Context, id string) (*domain.Book, error)
	Create(ctx context.Context, book domain.Book) (*domain.Book, error)
	Update(ctx context.Context, id string, book domain.Book) (*domain.Book, error)
	Delete(ctx context.Context, id string) error
}

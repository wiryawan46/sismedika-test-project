package service

import (
	"context"
	"sismedika-test-project/internal/book/domain"
)

// BookRepository adalah interface abstraksi untuk operasi terhadap entitas Book.
type BookService interface {
	GetAllBooks(ctx context.Context) ([]domain.Book, error)
	GetBookByID(ctx context.Context, id string) (*domain.Book, error)
	CreateBook(ctx context.Context, input domain.Book) (*domain.Book, error)
	UpdateBook(ctx context.Context, id string, input domain.Book) (*domain.Book, error)
	DeleteBook(ctx context.Context, id string) error
}

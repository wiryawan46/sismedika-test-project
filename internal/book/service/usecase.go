package service

import (
	"context"
	"errors"
	"sismedika-test-project/internal/book/repository"

	"github.com/google/uuid"
	"sismedika-test-project/internal/book/domain"
)

// BookService adalah implementasi usecase untuk manajemen buku.
// bookService mengimplementasikan BookService
type bookService struct {
	repo repository.BookRepository
}

// NewBookService membuat instance BookService
func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{repo: repo}
}

// GetAllBooks memanggil repository untuk mengambil semua buku
func (s *bookService) GetAllBooks(ctx context.Context) ([]domain.Book, error) {
	return s.repo.GetAll(ctx)
}

// GetBookByID memanggil repository untuk mengambil detail buku
func (s *bookService) GetBookByID(ctx context.Context, id string) (*domain.Book, error) {
	return s.repo.GetByID(ctx, id)
}

// CreateBook menambahkan buku baru, sekaligus memvalidasi input
func (s *bookService) CreateBook(ctx context.Context, input domain.Book) (*domain.Book, error) {
	if input.Title == "" || input.Author == "" || input.PublishedYear <= 0 {
		return nil, errors.New("invalid input")
	}
	input.ID = uuid.New().String()
	return s.repo.Create(ctx, input)
}

// UpdateBook memperbarui data buku berdasarkan ID
func (s *bookService) UpdateBook(ctx context.Context, id string, input domain.Book) (*domain.Book, error) {
	if input.Title == "" || input.Author == "" || input.PublishedYear <= 0 {
		return nil, errors.New("invalid input")
	}
	return s.repo.Update(ctx, id, input)
}

// DeleteBook menghapus buku berdasarkan ID
func (s *bookService) DeleteBook(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

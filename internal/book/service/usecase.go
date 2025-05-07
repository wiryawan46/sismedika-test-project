package service

import (
	"log"
	"sismedika-test-project/internal/shared"

	"github.com/google/uuid"
	"sismedika-test-project/internal/book/domain"
)

// BookService adalah implementasi usecase untuk manajemen buku.
type BookService struct {
	repo domain.BookRepository
}

// NewBookService adalah constructor untuk membuat BookService baru.
func NewBookService(repo domain.BookRepository) *BookService {
	return &BookService{repo: repo}
}

// GetAllBooks mengambil semua buku dari repository.
func (s *BookService) GetAllBooks() ([]domain.Book, error) {
	log.Println("Usecase: Mengambil semua buku")
	return s.repo.GetAll()
}

// GetBookByID mengambil satu buku berdasarkan ID.
func (s *BookService) GetBookByID(id string) (domain.Book, error) {
	log.Printf("Usecase: Mengambil buku dengan ID %s\n", id)
	return s.repo.GetByID(id)
}

// CreateBook membuat buku baru dan mengatur ID-nya.
func (s *BookService) CreateBook(book domain.Book) error {
	log.Println("Usecase: Menambahkan buku baru")
	if err := shared.ValidateBookInput(book); err != nil {
		return err
	}
	book.ID = uuid.NewString()
	return s.repo.Create(book)
}

// UpdateBook memperbarui data buku berdasarkan ID.
func (s *BookService) UpdateBook(id string, book domain.Book) error {
	log.Printf("Usecase: Memperbarui buku dengan ID %s\n", id)
	if err := shared.ValidateBookInput(book); err != nil {
		return err
	}
	return s.repo.Update(id, book)
}

// DeleteBook menghapus buku berdasarkan ID.
func (s *BookService) DeleteBook(id string) error {
	log.Printf("Usecase: Menghapus buku dengan ID %s\n", id)
	return s.repo.Delete(id)
}

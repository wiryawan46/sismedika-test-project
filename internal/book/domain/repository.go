package domain

// BookRepository adalah interface abstraksi untuk operasi terhadap entitas Book.
type BookRepository interface {
	GetAll() ([]Book, error)
	GetByID(id string) (Book, error)
	Create(book Book) error
	Update(id string, book Book) error
	Delete(id string) error
}

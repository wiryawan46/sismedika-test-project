package shared

import (
	"errors"
	"sismedika-test-project/internal/book/domain"
)

// ValidateBookInput mengecek kebutuhan field harus terpenuhi
func ValidateBookInput(book domain.Book) error {
	if book.Title == "" {
		return errors.New("title is required")
	}
	if book.Author == "" {
		return errors.New("author is required")
	}
	if book.PublishedYear <= 0 {
		return errors.New("published_year must be a positive integer")
	}
	return nil
}

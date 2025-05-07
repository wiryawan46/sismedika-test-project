package repository

import (
	"context"
	"github.com/stretchr/testify/assert"
	"sismedika-test-project/internal/book/domain"
	"sismedika-test-project/internal/book/infrastructure/storage"
	"sismedika-test-project/internal/mocks"
	"testing"
)

func TestGetAll(t *testing.T) {
	mockStore := mocks.NewMockStore()
	mockRepo := NewBookRepository((*storage.BookStore)(mockStore))

	// Menambahkan buku ke store
	mockStore.Books["1"] = storage.BookData{
		ID:            "1",
		Title:         "Book 1",
		Author:        "Author 1",
		PublishedYear: 2021,
	}

	// Test GetAll
	books, err := mockRepo.GetAll(context.Background())
	assert.NoError(t, err)
	assert.Len(t, books, 1)
	assert.Equal(t, "Book 1", books[0].Title)
}

func TestGetByID_Found(t *testing.T) {
	mockStore := mocks.NewMockStore()
	mockRepo := NewBookRepository((*storage.BookStore)(mockStore))

	// Menambahkan buku ke store
	mockStore.Books["1"] = storage.BookData{
		ID:            "1",
		Title:         "Book 1",
		Author:        "Author 1",
		PublishedYear: 2021,
	}

	// Test GetByID with valid ID
	book, err := mockRepo.GetByID(context.Background(), "1")
	assert.NoError(t, err)
	assert.NotNil(t, book)
	assert.Equal(t, "Book 1", book.Title)
}

func TestGetByID_NotFound(t *testing.T) {
	mockStore := mocks.NewMockStore()
	mockRepo := NewBookRepository((*storage.BookStore)(mockStore))

	// Test GetByID with invalid ID
	book, err := mockRepo.GetByID(context.Background(), "999")
	assert.Error(t, err)
	assert.Nil(t, book)
	assert.Equal(t, "book not found", err.Error())
}

func TestCreate(t *testing.T) {
	mockStore := mocks.NewMockStore()
	mockRepo := NewBookRepository((*storage.BookStore)(mockStore))

	// Test Create
	book := domain.Book{
		ID:            "1",
		Title:         "New Book",
		Author:        "Author 1",
		PublishedYear: 2021,
	}

	createdBook, err := mockRepo.Create(context.Background(), book)
	assert.NoError(t, err)
	assert.Equal(t, "New Book", createdBook.Title)
	assert.Contains(t, mockStore.Books, "1")
}

func TestUpdate_Found(t *testing.T) {
	mockStore := mocks.NewMockStore()
	mockRepo := NewBookRepository((*storage.BookStore)(mockStore))

	// Menambahkan buku ke store
	mockStore.Books["1"] = storage.BookData{
		ID:            "1",
		Title:         "Old Title",
		Author:        "Old Author",
		PublishedYear: 2020,
	}

	// Test Update
	book := domain.Book{
		Title:         "Updated Title",
		Author:        "Updated Author",
		PublishedYear: 2023,
	}

	updatedBook, err := mockRepo.Update(context.Background(), "1", book)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Title", updatedBook.Title)
	assert.Equal(t, "Updated Author", updatedBook.Author)
	assert.Equal(t, 2023, updatedBook.PublishedYear)
}

func TestUpdate_NotFound(t *testing.T) {
	mockStore := mocks.NewMockStore()
	mockRepo := NewBookRepository((*storage.BookStore)(mockStore))

	// Test Update with non-existent ID
	book := domain.Book{
		Title:         "Updated Title",
		Author:        "Updated Author",
		PublishedYear: 2023,
	}

	updatedBook, err := mockRepo.Update(context.Background(), "999", book)
	assert.Error(t, err)
	assert.Nil(t, updatedBook)
	assert.Equal(t, "book not found", err.Error())
}

func TestDelete_Found(t *testing.T) {
	mockStore := mocks.NewMockStore()
	mockRepo := NewBookRepository((*storage.BookStore)(mockStore))

	// Menambahkan buku ke store
	mockStore.Books["1"] = storage.BookData{
		ID:            "1",
		Title:         "Book to delete",
		Author:        "Author",
		PublishedYear: 2021,
	}

	// Test Delete
	err := mockRepo.Delete(context.Background(), "1")
	assert.NoError(t, err)
	assert.NotContains(t, mockStore.Books, "1")
}

func TestDelete_NotFound(t *testing.T) {
	mockStore := mocks.NewMockStore()
	mockRepo := NewBookRepository((*storage.BookStore)(mockStore))

	// Test Delete with non-existent ID
	err := mockRepo.Delete(context.Background(), "999")
	assert.Error(t, err)
	assert.Equal(t, "book not found", err.Error())
}

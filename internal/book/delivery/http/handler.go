package http

import (
	"encoding/json"
	"net/http"
	logger "sismedika-test-project/internal/middleware/logger"

	"sismedika-test-project/internal/book/domain"
	usecase "sismedika-test-project/internal/book/service"

	"github.com/go-chi/chi/v5"
)

// BookHandler adalah handler untuk endpoint buku
type BookHandler struct {
	Service usecase.BookService
}

// NewBookHandler membuat handler baru
func NewBookHandler(service usecase.BookService) *BookHandler {
	return &BookHandler{Service: service}
}

// GetAllBooks menangani GET /books
func (h *BookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.Service.GetAllBooks(r.Context())
	if err != nil {
		logger.Error("failed to get books:", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(books)
}

// GetBookByID menangani GET /books/{id}
func (h *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	book, err := h.Service.GetBookByID(r.Context(), id)
	if err != nil {
		logger.Error("failed to get book:", err)
		http.Error(w, "book not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(book)
}

// CreateBook menangani POST /books
func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book domain.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		logger.Error("invalid input:", err)
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}
	created, err := h.Service.CreateBook(r.Context(), book)
	if err != nil {
		logger.Error("failed to create book:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

// UpdateBook menangani PUT /books/{id}
func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var book domain.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		logger.Error("invalid input:", err)
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}
	updated, err := h.Service.UpdateBook(r.Context(), id, book)
	if err != nil {
		logger.Error("failed to update book:", err)
		http.Error(w, "book not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(updated)
}

// DeleteBook menangani DELETE /books/{id}
func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := h.Service.DeleteBook(r.Context(), id)
	if err != nil {
		logger.Error("failed to delete book:", err)
		http.Error(w, "book not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

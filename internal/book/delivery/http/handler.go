package http

import (
	"encoding/json"
	"net/http"
	logger "sismedika-test-project/internal/middleware/logger"

	"sismedika-test-project/internal/book/domain"
	usecase "sismedika-test-project/internal/book/service"
	response "sismedika-test-project/internal/shared"

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
		response.ErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, books, "success")
}

// GetBookByID menangani GET /books/{id}
func (h *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	book, err := h.Service.GetBookByID(r.Context(), id)
	if err != nil {
		logger.Error("failed to get book:", err)
		response.ErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, book, "success")
}

// CreateBook menangani POST /books
func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book domain.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		logger.Error("invalid input:", err)
		response.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	created, err := h.Service.CreateBook(r.Context(), book)
	if err != nil {
		logger.Error("failed to create book:", err)
		response.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	response.JSON(w, http.StatusCreated, created, "Data Created!")
}

// UpdateBook menangani PUT /books/{id}
func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var book domain.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		logger.Error("invalid input:", err)
		response.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	updated, err := h.Service.UpdateBook(r.Context(), id, book)
	if err != nil {
		logger.Error("failed to update book:", err)
		response.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	response.JSON(w, http.StatusCreated, updated, "Data Updated!")
}

// DeleteBook menangani DELETE /books/{id}
func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := h.Service.DeleteBook(r.Context(), id)
	if err != nil {
		logger.Error("failed to delete book:", err)
		response.ErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, nil, "Data Deleted Success!")
}

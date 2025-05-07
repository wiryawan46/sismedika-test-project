package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	bookHttp "sismedika-test-project/internal/book/delivery/http"
	infrastructure "sismedika-test-project/internal/book/infrastructure/storage"
	bookRepo "sismedika-test-project/internal/book/repository"
	bookService "sismedika-test-project/internal/book/service"
	"sismedika-test-project/internal/middleware/logger"
)

func main() {
	// Inisialisasi dependency untuk buku
	bookStore := infrastructure.GetBookStore()                // Singleton store
	bookRepository := bookRepo.NewBookRepository(bookStore)   // Repository
	bookUsecase := bookService.NewBookService(bookRepository) // Service (Usecase)
	bookHandler := bookHttp.NewBookHandler(bookUsecase)       // Handler

	// Router utama
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)    // Recover from panics
	r.Use(logger.LoggerMiddleware) // Custom logger middleware

	// Register route book
	bookHttp.RegisterBookRoutes(r, bookHandler)

	logger.Info("Starting server on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		logger.Error("Server failed: %v", err)
	}
}

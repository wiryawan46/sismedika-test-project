package http

import (
	"github.com/go-chi/chi/v5"
)

// RegisterBookRoutes mendaftarkan semua endpoint untuk resource buku
func RegisterBookRoutes(r chi.Router, h *BookHandler) {
	r.Route("/books", func(r chi.Router) {
		r.Get("/", h.GetAllBooks)       // GET /books
		r.Get("/{id}", h.GetBookByID)   // GET /books/{id}
		r.Post("/", h.CreateBook)       // POST /books
		r.Put("/{id}", h.UpdateBook)    // PUT /books/{id}
		r.Delete("/{id}", h.DeleteBook) // DELETE /books/{id}
	})
}

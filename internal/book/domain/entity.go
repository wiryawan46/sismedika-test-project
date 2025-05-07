package domain

// Book adalah entitas utama untuk sistem manajemen buku.
type Book struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedYear int    `json:"published_year"`
}

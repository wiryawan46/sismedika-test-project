# 📚 Book Management API - Clean Architecture (Go + Chi)

Sistem RESTful API sederhana untuk manajemen data buku, dibangun dengan bahasa Go dan menggunakan framework `go-chi`. Project ini menggunakan arsitektur **Clean Architecture**, penyimpanan data **in-memory**, serta logger middleware custom untuk mencatat semua request.

---

## ✨ Fitur

- ✅ GET `/books` – Mendapatkan semua data buku
- ✅ GET `/books/{id}` – Mendapatkan detail buku berdasarkan ID
- ✅ POST `/books` – Menambahkan buku baru
- ✅ PUT `/books/{id}` – Memperbarui buku
- ✅ DELETE `/books/{id}` – Menghapus buku
- ✅ Middleware Logger – Mencatat method, path, dan waktu eksekusi
- ✅ Penyimpanan menggunakan **singleton in-memory**
- ✅ Struktur **Clean Architecture** modular dan testable
- ✅ Unit test untuk service dan repository

---

## 🗂 Struktur Folder (Clean Architecture)

```
your_project/
├── cmd/                  # Main Application Entry Point
│   └── main.go
├── internal/
│   ├── book/           # Domain
│       ├── domain/           # Entity & Interface (Core)
│       ├── delivery/         # Use Case (Business Logic)
│       ├── infrastructure/       # Adapter (Storage Layer)
│       ├── repository/   # In-Memory Storage (Singleton)
│       └── service/          # HTTP Routing and Handler
│   ├── middleware/           # Mock for Test
|       ├── logger/           # Custom Logger Middleware
│   ├── mocks/           # Mock for Test
│   ├── shared/          # Custom Validator and Response
├── go.mod
├── Sismedika Project Test.postman_collection.json
├── start.txt            # The start of development project
├── end.txt              # The end of development project
└── README.md
```

---

## 🚀 Cara Menjalankan Aplikasi

### 1. Clone Repository

```bash
git clone https://github.com/wiryawan46/sismedika-test-project.git
cd sismedika-test-project
```

### 2. Inisialisasi Go Modules (jika perlu)

```bash
go mod tidy
```

### 3. Jalankan Aplikasi

```bash
go run cmd/main.go
```

Server akan berjalan di `http://localhost:8080`.

---

## ✅ Contoh Request

### POST `/books`

```json
{
  "title": "Clean Architecture",
  "author": "Robert C. Martin",
  "published_year": 2017
}
```

---

## 🧪 Menjalankan Unit Test

Untuk saat ini unit test hanay tersedia untuk `repository` saja

```bash
go test ./...
```

Jika ingin menjalankan unit test dengan coverage:

```bash
go test -v -cover ./...
```

---

## 🧱 Teknologi & Library

- [Golang](https://golang.org/)
- [go-chi v5](https://github.com/go-chi/chi)
- Standard library (log, net/http, sync)
- Clean Architecture Pattern
- In-memory singleton storage

---

## 📁 File Penanda Proses

Untuk menandakan proses pengerjaan:

- ⏳ `start.txt` → dibuat saat mulai
- ✅ `end.txt` → dibuat saat selesai

---

## 🧑‍💻 Author

Made with ❤️ by Manggala Pramuditya Wiryawan

---

## 📝 License

MIT License

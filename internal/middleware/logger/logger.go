package logger

import (
	"log"
	"time"
)

// Info mencetak pesan informasi
func Info(msg string, args ...any) {
	log.Printf("[INFO] "+msg, args...)
}

// Error mencetak pesan kesalahan
func Error(msg string, args ...any) {
	log.Printf("[ERROR] "+msg, args...)
}

// WithDuration mencetak log dengan durasi
func WithDuration(start time.Time, msg string, args ...any) {
	elapsed := time.Since(start).Milliseconds()
	log.Printf("[DURATION] %s - %dms", msg, elapsed)
}

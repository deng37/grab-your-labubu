package middleware

import (
	"net/http"
	"os"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-KEY")
		expectedKey := os.Getenv("LABUBU_API_KEY")

		if apiKey == "" || apiKey != expectedKey {
			http.Error(w, "Unauthorized: Invalid or missing API Key", http.StatusUnauthorized)
			return
		}

		// Jika valid, lanjut ke handler berikutnya
		next.ServeHTTP(w, r)
	})
}
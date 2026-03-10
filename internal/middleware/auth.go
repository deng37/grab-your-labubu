package middleware

import (
	"net/http"
	"os"
	"strings"
	"fmt"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		publicRoutes := []string{"/", "/favicon.ico"}		// Checking public routes
		isPublic := false
		for _, path := range publicRoutes {
			if r.URL.Path == path {
				isPublic = true
				break
			}
		}

		if strings.HasPrefix(r.URL.Path, "/assets/") {		// Setting public for assets
			isPublic = true
		}

		if !isPublic {		// Not public, checking API key
			key := r.Header.Get("X-API-KEY")
			if key != os.Getenv("LABUBU_API_KEY") {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"error": "Unauthorized", "message": "Missing API Key"}`)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
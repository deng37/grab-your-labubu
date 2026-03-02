package util

import (
	"net/http"
)

func UpdateHeaderJson(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
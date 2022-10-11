// GetAllBooks Endpoint

package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/mstomar698/go-REST-api/package/mocks"
)

// func GetAllBooks(w http.ResponseWriter, r *http.Request) {
func (h handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Books)
}
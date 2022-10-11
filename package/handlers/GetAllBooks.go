// GetAllBooks Endpoint

package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	// "github.com/mstomar698/go-REST-api/package/mocks"
	"github.com/mstomar698/go-REST-api/package/models"
)

// func GetAllBooks(w http.ResponseWriter, r *http.Request) {
func (h handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book

	if result := h.DB.Find(&books); result.Error != nil {
		fmt.Println(result.Error)
	}
	
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(mocks.Books)
	json.NewEncoder(w).Encode(books)
}
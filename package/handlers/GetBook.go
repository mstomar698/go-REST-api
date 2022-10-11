package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	// "github.com/mstomar698/go-REST-api/package/mocks"
	"github.com/mstomar698/go-REST-api/package/models"
)

// func GetBook(w http.ResponseWriter, r * http.Request) {
func (h handler) GetBook(w http.ResponseWriter, r *http.Request) {
	// Check dynamic parameter {id}
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// for db it will we writtena s
	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)

	// check in all mocks for the {id}
	// for _, book := range mocks.Books {
	// 	if book.Id == id {
	// 		w.Header().Add("Content-Type", "application/json")
	// 		w.WriteHeader(http.StatusOK)
	// 		json.NewEncoder(w).Encode(book)
	// 		break
	// 	}
	// }
}

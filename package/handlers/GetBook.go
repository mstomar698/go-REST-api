package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/mstomar698/go-REST-api/package/mocks"
)

func GetBook(w http.ResponseWriter, r * http.Request) {
	// Check dynamic parameter {id}
	vars := mux.vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// check in all mocks for the {id}
	for _, book := range mocks.Books {
		if book.Id == id {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(book)
			break
		}
	}
}
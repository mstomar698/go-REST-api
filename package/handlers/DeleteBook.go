package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mstomar698/go-REST-api/package/mocks"
)

// func DeleteBook(w http.ResponseWriter, r *http.Request) {
func (h handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	// searching param
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	//iterating
	for index, book := range mocks.Books {
		if book.Id == id {
			// delete book at {id}
			mocks.Books = append(mocks.Books[:index], mocks.Books[index+1:]...)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Book Deleted")
			break
		}
	}
}

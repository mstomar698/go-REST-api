package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	// "github.com/mstomar698/go-REST-api/package/mocks"
	"github.com/mstomar698/go-REST-api/package/models"
)

// func UpdateBook(w http.ResponseWriter, r * http.Request) {
func (h handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	// Check dynamic parameter {id}
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Request body for {id}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updateBook models.Book
	json.Unmarshal(body, &updateBook)
	// for db we will use this
	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	book.Title = updateBook.Title
	book.Author = updateBook.Author
	book.Desc = updateBook.Desc

	h.DB.Save(&book)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Book Updated")

	// check in all mocks for the {id}
	// for index, book := range mocks.Books {
	// 	if book.Id == id {
	// 		// Update book on matching
	// 		book.Title = updateBook.Title
	// 		book.Author = updateBook.Author
	// 		book.Desc = updateBook.Desc

	// 		mocks.Books[index] = book
	// 		w.Header().Add("Content-Type", "application/json")
	// 		w.WriteHeader(http.StatusOK)
	// 		json.NewEncoder(w).Encode("Book Updated")
	// 		break
	// 	}
	// }
}

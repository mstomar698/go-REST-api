package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	// "math/rand"
	"fmt"
	"net/http"

	// "github.com/mstomar698/go-REST-api/package/mocks"
	"github.com/mstomar698/go-REST-api/package/models"
)

// func AddBook(w http.ResponseWriter, r *http.Request) {
func (h handler) AddBook(w http.ResponseWriter, r *http.Request) {
	// Requesting a body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var book models.Book
	json.Unmarshal(body, &book)

	// Add books to the mock db
	// book.Id = rand.Intn(100)
	// mocks.Books = append(mocks.Books, book)
	// In place of adding them via mock we will now add them using db
	if  result := h.DB.Create(&book); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Sending confirmation message
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created new book")
}
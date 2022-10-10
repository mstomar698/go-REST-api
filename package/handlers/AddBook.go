package handlers

import (
	"encoding/josn"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	"github.com/mstomar698/go-REST-api/package/mocks"
	"github.com/mstomar698/go-REST-api/package/models"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	// Requesting a body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var book models.Book
	json.Unmarshal(body, &book)

	// Add books to the mock db
	book.Id = rand.Intn(100)
	mocks.Books = append(mocks.Books, book)

	// Sending confirmation message
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created new books")
}
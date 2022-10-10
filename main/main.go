package main

import (
	"log"
	"net/http"

	// "encoding/json"
	"github.com/gorilla/mux"
	"github.com/mstomar698/go-REST-api/package/handlers"
)

func main() {
	router := mux.NewRouter()

	// Api endpoints will come here
	// router.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
	// 	json.NewEncoder(w).Encode("Hello World")
	// })
	// First Endpoint
	router.HandleFunc("/books", handlers.GetAllBooks).Methods(http.MethodGet)
	// Second Endpoint
	router.HandleFunc("/books", handlers.AddBook).Methods(http.MethodPost)
	// Third Endpoint
	router.HandleFunc("/books/{id}", handlers.GetBook).Methods(http.MethodGet)
	// Fourth Endpoint
	router.HandleFunc("/books/{id}", handlers.UpdateBook).Methods(http.MethodPut)
	
	log.Println("Api is running on port 4000")
	http.ListenAndServe(":4000", router)
}

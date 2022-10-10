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
	router.HandleFunc("/books", handlers.GetAllBooks).Methods(http.MethodGet)

	log.Println("Api is running on port 4000")
	http.ListenAndServe(":4000", router)
}
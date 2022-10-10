package main

import (
	"log"
	"net/http"
	// "encoding/json"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Api endpoints will come here
	// router.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
	// 	json.NewEncoder(w).Encode("Hello World")
	// })

	log.Println("Api is running on port 4000")
	http.ListenAndServe(":4000", router)
}
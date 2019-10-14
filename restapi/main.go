package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var books []Book

func main() {
	// Init router
	r := mux.NewRouter()

	books = append(books, Book{ID: "1", Isbn: "111111", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "222222", Title: "Book Two", Author: &Author{Firstname: "Tom", Lastname: "Wong"}})

	// Router handlers
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// Http server
	log.Fatal(http.ListenAndServe(":8888", r))
}

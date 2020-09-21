package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/este001/restapi/Model"
	"github.com/gorilla/mux"
)

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r) // Ger params
	// Loop through books and find with id
	for _, item := range books {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Model.Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var book Model.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.Id = strconv.Itoa(rand.Intn(1000000)) //Mock ID - not safe
	books = append(books, book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.Id == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Model.Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.Id = params["id"] //Mock ID - not safe
			books = append(books, book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.Id == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

//Init books variable as a slice Book struct
var books []Model.Book

func main() {
	//Init Router
	r := mux.NewRouter()

	// Mock data - @todo - implement DB
	books = append(books, Model.Book{Id: "1", Isbn: "44312312", Title: "Book one", Author: &Model.Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Model.Book{Id: "2", Isbn: "12346556", Title: "Book two", Author: &Model.Author{Firstname: "Jane", Lastname: "Doe"}})

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	
	//curl -i -X POST -H 'Content-Type: application/json' -d '{"isbn": "213213", "title": "Book three", "author" : {"firstname": "Carol", "lastname": "Williams"} }' http://localhost:8000/api/books
	r.HandleFunc("/api/books", createBook).Methods("POST")
	
	//curl -i -X PUT -H 'Content-Type: application/json' -d '{"isbn": "213213", "title": "Book three", "author" : {"firstname": "Carol", "lastname": "Williams"} }' http://localhost:8000/api/books/2
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	
	//curl -i -X DELETE http://localhost:8000/api/books/1
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

	fmt.Print(strconv.Itoa(1))
	rand.Int()

}

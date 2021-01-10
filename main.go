package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Book struct : model
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

var books []Book

//Author struct : model
type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// GET all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "hello world!!!")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //map
	// fmt.Fprint(w, params["id"])
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(Book{})
}
func createBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000)) // r.Body hasn't field ID

	books = append(books, book)

	json.NewEncoder(w).Encode(book)
}
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(Book{})
}
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	//data post
	book := Book{}
	json.NewDecoder(r.Body).Decode(&book)

	// fmt.Fprint(w, book.Author)

	for index, item := range books {
		if item.ID == params["id"] {
			if book.Isbn != "" {
				books[index].Isbn = book.Isbn
			}
			if book.Title != "" {
				books[index].Title = book.Title
			}
			if book.Author != nil {
				books[index].Author = book.Author
			}
			json.NewEncoder(w).Encode(books[index])
			return
		}
	}
	json.NewEncoder(w).Encode(Book{})

}
func main() {
	// fmt.Println("nghiahsgs")
	r := mux.NewRouter()

	//mock data
	author1 := Author{FirstName: "nghia", LastName: "hsgs"}
	author2 := Author{FirstName: "nghia", LastName: "nguyen"}
	books = append(books, Book{ID: "123", Isbn: "123abc", Title: "this is title one", Author: &author1})
	books = append(books, Book{ID: "456", Isbn: "456abc", Title: "this is title two", Author: &author1})
	books = append(books, Book{ID: "789", Isbn: "789abc", Title: "this is title three", Author: &author2})

	// Route handler
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBooks).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3000", r))

}

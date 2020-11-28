package api

import (
	"GoRestAPI/data"
	"GoRestAPI/model"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetBooks ...
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.Books)
}

// GetBook ...
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range data.Books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&model.Book{})
}

// CreateBook ...
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book model.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000000))
	data.Books = append(data.Books, book)
	json.NewEncoder(w).Encode(book)
}

// UpdateBook ...
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range data.Books {
		if item.ID == params["id"] {
			data.Books = append(data.Books[:index], data.Books[index+1:]...)
			var book model.Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			data.Books = append(data.Books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

// DeleteBook ...
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range data.Books {
		if item.ID == params["id"] {
			data.Books = append(data.Books[:index], data.Books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(data.Books)
}

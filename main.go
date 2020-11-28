package main

import (
	"GoRestAPI/data"
	"GoRestAPI/model"
	"GoRestAPI/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := routes.Router()

	data.Books = []model.Book{
		model.Book{ID: "1", Isbn: "123", Title: "Book 1", Author: &model.Author{Firstname: "sakib", Lastname: "alamin"}},
		model.Book{ID: "2", Isbn: "234", Title: "Book 2", Author: &model.Author{Firstname: "nova", Lastname: "khan"}},
		model.Book{ID: "3", Isbn: "345", Title: "Book 3", Author: &model.Author{Firstname: "lionel", Lastname: "messi"}},
		model.Book{ID: "4", Isbn: "456", Title: "Book 4", Author: &model.Author{Firstname: "diego", Lastname: "maradona"}},
	}

	fmt.Println("server running at port :8080")
	log.Fatal(http.ListenAndServe(":8080", &router))
}

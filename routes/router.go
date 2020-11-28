package routes

import (
	"GoRestAPI/api"

	"github.com/gorilla/mux"
)

// Router ...
func Router() mux.Router {
	Router := mux.NewRouter()

	Router.HandleFunc("/api/books", api.GetBooks).Methods("GET")
	Router.HandleFunc("/api/books/{id}", api.GetBook).Methods("GET")
	Router.HandleFunc("/api/books", api.CreateBook).Methods("POST")
	Router.HandleFunc("/api/books/{id}", api.UpdateBook).Methods("PUT")
	Router.HandleFunc("/api/books/{id}", api.DeleteBook).Methods("DELETE")

	return *Router
}

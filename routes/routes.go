package routes

import (
	"book_store/controllers"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
    // Book Routes
    r.HandleFunc("/books", controllers.GetBooks).Methods("GET")
    r.HandleFunc("/books/{id}", controllers.GetBook).Methods("GET")
    r.HandleFunc("/books", controllers.CreateBook).Methods("POST")
    r.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
    r.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")

    // Employee Routes
    r.HandleFunc("/employees", controllers.GetEmployees).Methods("GET")
    r.HandleFunc("/employees", controllers.CreateEmployee).Methods("POST")
}

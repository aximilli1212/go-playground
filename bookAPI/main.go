package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//Book Struct(Model)

type Book struct{
	ID string `json:"id"`
}

func main(){
	//Init mux router

	r := mux.NewRouter() // : type inference

	//Router handler / Endpoints

	r.HandleFunc("/api/books",getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}",getBook).Methods("GET")
	r.HandleFunc("/api/books",createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}",updateBooks).Methods("PUT")
	r.HandleFunc("/api/books",deleteBooks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000",r))


}
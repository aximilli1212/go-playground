package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

//Book Struct(Model)

type Book struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
  type Author struct {
  	Firstname string `json:"firstname"`
  	Lastname string `json:"lastname"`
  }

// init books var as slice Book struct
var books []Book

// Get all Books
func getBooks(w http.ResponseWriter, r *http.Request){
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(books)
}

// Get all Books
func getBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //get params
	//Loop through books and find with id
	for _,item := range books{
		if item.ID == params["id"]{
		json.NewEncoder(w).Encode(item)
		return
	}
	}

	json.NewEncoder(w).Encode(&Book{})

}// Get all Books
func createBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
  var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000)) //Mock ID - not Safe
	books = append(books,book)
	json.NewEncoder(w).Encode(&Book{})

}// Get all Books
func updateBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r);

	for index, item := range books{
		if item.ID == params["id"]{
			books = append(books[:index], books[index+1:]... )
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = strconv.Itoa(rand.Intn(1000000)) // MOck ID

			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}


}// Get all Books
func deleteBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params :=mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"]{
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
              json.NewEncoder(w).Encode(books)

}

func main(){
	//Init mux router

	r := mux.NewRouter() // : type inference



	// Mock Data - implement DB
	books = append(books, Book{ID:"1",Isbn:"338383", Title:"Sweet Victory", Author: &Author{Firstname:"John", Lastname: "Doe"}})
	books = append(books, Book{ID:"2",Isbn:"8383", Title:"Honey Victory", Author: &Author{Firstname:"Jane", Lastname: "Drax"}})

	//Router handler / Endpoints

	r.HandleFunc("/api/books",getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}",getBook).Methods("GET")
	r.HandleFunc("/api/books",createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}",updateBooks).Methods("PUT")
	r.HandleFunc("/api/books/{id}",deleteBooks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000",r))


}
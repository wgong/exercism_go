package main

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/zenazn/goji"
)

type book struct {
	ISBN    string "json:isbn"
	Title   string "json:name"
	Authors string "json:author"
	Price   string "json:price"
}

var bookStore = []book{
	book{
		ISBN:    "0321774639",
		Title:   "Programming in Go: Creating Applications for the 21st Century (Developer's Library)",
		Authors: "Mark Summerfield",
		Price:   "$34.57",
	},
	book{
		ISBN:    "0134190440",
		Title:   "The Go Programming Language",
		Authors: "Alan A. A. Donovan, Brian W. Kernighan",
		Price:   "$34.57",
	},
}

func main() {
	goji.Get("/", hello)
	goji.Get("/books", allBooks)
	//goji.Get("/books/:isbn", bookByISBN)
	goji.Serve()
}

func allBooks(w http.ResponseWriter, r *http.Request) {
	jsonOut, _ := json.Marshal(bookStore)
	fmt.Fprintf(w, string(jsonOut))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Goji")
}

/*
func bookByISBN(w http.ResponseWriter, r *http.Request) {
	isbn := pat.Param(r, "isbn")
	for _, b := range bookStore {
		if b.ISBN == isbn {
			jsonOut, _ := json.Marshal(b)
			fmt.Fprintf(w, string(jsonOut))
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}
*/

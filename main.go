package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"example.com/basicServer/database"
)

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./static/index.html"))
	tmpl.Execute(w, nil)
}

func aboutPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./static/about.html"))
	tmpl.Execute(w, nil)
}

func booksPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./static/books.html"))
	tmpl.Execute(w, database.Books)
}

func createBookHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./static/add_book.html"))
	tmpl.Execute(w, nil)
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homePageHandler(w, r)

	case "/about":
		aboutPageHandler(w, r)

	case "/books":
		booksPageHandler(w, r)

	case "/add-book":
		createBookHandler(w, r)

	default:
		http.ServeFile(w, r, "./static/error.html")
	}
}

func main() {
	http.HandleFunc("/", requestHandler)

	fmt.Println("Starting the server at port:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

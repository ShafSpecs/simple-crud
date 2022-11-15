package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("static/index.html"))

	_ = tpl.Execute(w, nil)
}

func about(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("static/about.html"))

	_ = tpl.Execute(w, nil)
}

func main() {
	mux := http.NewServeMux()
	err := godotenv.Load()

	if err != nil {
		log.Fatal(".env file not found")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// register assets
	fs := http.FileServer(http.Dir("assets"))

	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Interesting...
	mux.HandleFunc("/", index)
	//http.HandleFunc("/about", about)
	//mux.HandleFunc("/about", about)

	fmt.Println("Server Starting on port " + port + "...")
	http.ListenAndServe(":"+port, mux)
}

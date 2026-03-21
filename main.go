package main

import (
	"log"
	"net/http"
	"html/template"
)

func main() { 
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", handlerCreateInterface)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func handlerCreateInterface(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}
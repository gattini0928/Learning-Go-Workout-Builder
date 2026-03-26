package main

import (
	"log"
	"net/http"
	"github.com/gattini0928/Learning-Go-Workout-Builder/internal/routes"
)

func main() {
	mux := routes.SetupRoutes()
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
		return
	}
}
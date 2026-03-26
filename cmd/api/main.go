package main

import (
	"log"
	"net/http"
	"github.com/gattini0928/Learning-Go-Workout-Builder/internal/routes"
	"github.com/gattini0928/Learning-Go-Workout-Builder/internal/configs"
	"github.com/joho/godotenv"

)

func main() {

	godotenv.Load("../../.env")
	configs.Load()
	log.Println(configs.GetDb())
	mux := routes.SetupRoutes()
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
		return
	}
}
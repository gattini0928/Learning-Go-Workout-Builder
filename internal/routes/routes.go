package routes

import (
	"github.com/gattini0928/Learning-Go-Workout-Builder/internal/handlers"
	"net/http"
)

func SetupRoutes() *http.ServeMux{
	mux := http.NewServeMux()
	mux.HandleFunc("POST /signup", handlers.HandleSignup)
	return mux
}

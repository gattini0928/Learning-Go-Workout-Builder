package routes

import (
	"github.com/gattini0928/Learning-Go-Workout-Builder/internal/handlers"
	"net/http"
)

func SetupRoutes() *http.ServeMux{
	mux := http.NewServeMux()
	mux.HandleFunc("POST /signup", handlers.HandleSignup)
	mux.HandleFunc("POST /login", handlers.HandleLogin)
	mux.HandleFunc("POST /exercises", handlers.HandleExercise)
	mux.HandleFunc("POST /workouts", handlers.HandleWorkouts)
	mux.HandleFunc("POST /users/{id}/workouts/{workout_id}/{exercise_id}", handlers.HandleUpdateRepsAndSets)
	mux.HandleFunc("GET /exercises", handlers.HandleShowExercises)
	mux.HandleFunc("GET /users/{id}/workouts", handlers.HandleShowWorkouts)
	mux.HandleFunc("GET /workouts/{workout_id}", handlers.HandleShowWorkoutById)
	mux.HandleFunc("DELETE /users/{id}/workouts/{workout_id}", handlers.HandleDeleteWorkout)
	mux.HandleFunc("DELETE /exercises/{exercise_id}", handlers.HandleDeleteExercise)
	return mux
}


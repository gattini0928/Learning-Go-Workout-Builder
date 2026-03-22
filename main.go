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
	mux.HandleFunc("/workout-detail", handlerWorkoutDetail)
	mux.HandleFunc("/exercises", handlerExercises)
	mux.HandleFunc("/packages", handlerWorkoutsPackages)
	mux.HandleFunc("/build-workouts", handlerBuildWorkouts)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
 
func handlerCreateInterface(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func handlerWorkoutDetail(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/workout_detail.html"))
	tmpl.Execute(w, nil)
}

func handlerWorkoutsPackages(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/packages.html"))
	tmpl.Execute(w, nil)
}

func handlerBuildWorkouts(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/build_workout.html"))
	tmpl.Execute(w, nil)
}

func handlerExercises(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/exercises.html"))
	tmpl.Execute(w, nil)
}

package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gattini0928/Learning-Go-Workout-Builder/internal/models"
)

func HandleWorkouts(w http.ResponseWriter, r *http.Request) {
	var workout models.Workout
	var response Message

	err := json.NewDecoder(r.Body).Decode(&workout)
	if err != nil {
		log.Println(err)
		response.Content = err.Error()
		response.Status = 400
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(response)
		return
	}

	_, err = models.InsertWorkout(workout)
	if err != nil {
		log.Println(err)
		response.Content = err.Error()
		response.Status = 500
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Content = "Workout Criado com Sucesso"
	response.Status = 201
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

	log.Println(workout)
}
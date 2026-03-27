package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gattini0928/Learning-Go-Workout-Builder/internal/models"
)

func HandleExercise(w http.ResponseWriter, r *http.Request) {
	var exercise models.Exercise
	var response Message

	err := json.NewDecoder(r.Body).Decode(&exercise)
	if err != nil {
		log.Println(err)
		response.Content = err.Error()
		response.Status = 400
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(response)
		return
	}

	_, err = models.InsertExercice(exercise)
	if err != nil {
		log.Println(err)
		response.Content = err.Error()
		response.Status = 500
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Content = "Exercício Criado com Sucesso"
	response.Status = 201
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

	log.Println(exercise)
}
package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gattini0928/Learning-Go-Workout-Builder/internal/models"
)

func HandleShowExercises(w http.ResponseWriter, r *http.Request) {
	var response Message

	categorie := r.URL.Query().Get("categorie")

	var exercises []models.Exercise
	var err error

	if categorie != "" {
		exercises, err = models.GetExercicesByCategory(categorie)
	} else {
		exercises, err = models.GetAllExercices()
	}

	if err != nil {
		log.Println(err)
		response.Content = err.Error()
		response.Status = 500
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(response)
		return
	}

	json.NewEncoder(w).Encode(exercises)
}
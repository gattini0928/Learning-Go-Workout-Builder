package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"github.com/gattini0928/Learning-Go-Workout-Builder/internal/models"
)

func HandleShowWorkouts(w http.ResponseWriter, r *http.Request) {
	var response Message

	idStr := r.PathValue("id")
	if idStr == "" {
		response.Content = "id é obrigatório"
		response.Status = 400
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(response)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Content = "id inválido"
		response.Status = 400
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(response)
		return
	}

	workouts, err := models.GetAllWorkouts(id)
	if err != nil {
		log.Println(err)
		response.Content = err.Error()
		response.Status = 500
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(response)
		return
	}

	json.NewEncoder(w).Encode(workouts)
}
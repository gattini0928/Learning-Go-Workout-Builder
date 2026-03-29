package handlers


import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"github.com/gattini0928/Learning-Go-Workout-Builder/internal/models"
)

func HandleDeleteWorkout(w http.ResponseWriter, r *http.Request) {
	var response Message

	userIdStr := r.PathValue("id")
	if userIdStr == "" {
		response.Content = "id é obrigatório"
		response.Status = 400
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(response)
		return
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		response.Content = "id é inválido"
		response.Status = 400
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(response)
		return
	}

	workoutIdStr := r.PathValue("workout_id")
	if workoutIdStr == "" {
		response.Content = "workout_id é obrigatório"
		response.Status = 400
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(response)
		return
	}

	workoutId, err := strconv.Atoi(workoutIdStr)
	if err != nil {
		response.Content = "workout_id é inválido"
		response.Status = 400
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(response)
		return
	}

	rows, err := models.DeleteWorkout(workoutId, userId)
	if err != nil {
		log.Println(err)
		response.Content = err.Error()
		response.Status = 500
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(response)
		return
	}

	if rows == 0 {
		response.Content = "Treino não encontrado"
		response.Status = 404
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Content = "Treino excluído com sucesso"
	response.Status = 200
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response)

}
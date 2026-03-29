package handlers


import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"github.com/gattini0928/Learning-Go-Workout-Builder/internal/models"
)

func HandleDeleteExerciseFromWorkout(w http.ResponseWriter, r *http.Request) {
	var response Message

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

	exerciseIdStr := r.PathValue("exercise_id")
	if exerciseIdStr == "" {
		response.Content = "exercise_id é obrigatório"
		response.Status = 400
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(response)
		return
	}

	exerciseId, err := strconv.Atoi(exerciseIdStr)
	if err != nil {
		response.Content = "exercise_id é inválido"
		response.Status = 400
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(response)
		return
	}

	rows, err := models.DeleteExerciseFromWorkout(workoutId, exerciseId)
	if err != nil {
		log.Println(err)
		response.Content = err.Error()
		response.Status = 500
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(response)
		return
	}

	if rows == 0 {
		response.Content = "Exercício não encontrado"
		response.Status = 404
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Content = "Exercício excluído do treino"
	response.Status = 200
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response)
}
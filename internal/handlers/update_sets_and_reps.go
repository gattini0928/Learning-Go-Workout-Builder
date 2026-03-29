package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gattini0928/Learning-Go-Workout-Builder/internal/models"
)

type UpdatePayload struct {
	Reps *int `json:"reps"`
	Sets *int `json:"sets"`
}

func HandleUpdateRepsAndSets(w http.ResponseWriter, r *http.Request) {
	var response Message
	var payload UpdatePayload

	workoutID, err := strconv.Atoi(r.PathValue("workout_id"))
	if err != nil {
		response.Content = "workout_id inválido"
		response.Status = 400
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(response)
		return
	}

	exerciseID, err := strconv.Atoi(r.PathValue("exercise_id"))
	if err != nil {
		response.Content = "exercise_id inválido"
		response.Status = 400
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		response.Content = err.Error()
		response.Status = 400
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(response)
		return
	}

	if payload.Reps == nil && payload.Sets == nil {
		w.WriteHeader(204)
		return
	}

	current, err := models.GetWorkoutExercise(workoutID, exerciseID)
	if err != nil {
		response.Content = err.Error()
		response.Status = 500
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(response)
		return
	}

	reps := current.Reps
	sets := current.Sets

	if payload.Reps != nil {
		reps = *payload.Reps
	}

	if payload.Sets != nil {
		sets = *payload.Sets
	}

	rows, err := models.UpdateWorkoutExerciceRepsAndSets(workoutID, exerciseID, reps, sets)
	if err != nil {
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

	response.Content = "Atualizado com sucesso"
	response.Status = 200
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response)
}
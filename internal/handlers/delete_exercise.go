package handlers


import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"github.com/gattini0928/Learning-Go-Workout-Builder/internal/models"
)

func HandleDeleteExercise(w http.ResponseWriter, r *http.Request) {
	var response Message

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

	rows, err := models.DeleteExercice(exerciseId)
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

	response.Content = "Exercício excluído com sucesso"
	response.Status = 200
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response)

}
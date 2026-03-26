package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gattini0928/Learning-Go-Workout-Builder/internal/models"
	"github.com/gattini0928/Learning-Go-Workout-Builder/internal/validators"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	var requestUser models.User
	var response Message

	err := json.NewDecoder(r.Body).Decode(&requestUser)
	if err != nil {
		log.Println(err)
		response.Content = err.Error()
		response.Status = 400
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(response)
		return
	}

	userFromDb, err := models.GetUserByEmail(requestUser.Email)
	if err != nil {
		log.Println(err)
		response.Content = "Usuário não encontrado"
		response.Status = 404
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(response)
		return
	}

	ok := validators.CheckPasswordHash(requestUser.Password, userFromDb.Password)
	if !ok {
		response.Content = "Senha Ínvalida"
		response.Status = 401
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(response)
		return
	}
		
	response.Content = "Login realizado com sucesso"
	response.Status = 200
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(response)
	
}
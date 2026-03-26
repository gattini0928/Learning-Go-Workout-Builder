package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gattini0928/Learning-Go-Workout-Builder/internal/models"
	"github.com/gattini0928/Learning-Go-Workout-Builder/internal/validators"
)

type Message struct {
    Content string `json:"content"`
    Status  int `json:"status"`
}

func HandleSignup(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var response Message

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		response.Content = "Erro ao ler dados"
		response.Status = 400
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = validators.ValidateName(user.Name)
	if err != nil {
		log.Println(err)
		response.Content = err.Error()
		response.Status = 400
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = validators.ValidateEmail(user.Email)
	if err != nil {
		log.Println(err)
		response.Content = err.Error()
		response.Status = 400
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = validators.ValidatePassword(user.Password)
	if err != nil {
		log.Println(err)
		response.Content = err.Error()
		response.Status = 400
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(response)
		return
	}

	passwordHashed, err := validators.HashPassword(user.Password)
	if err != nil {
		log.Println(err)
		response.Content = err.Error()
		response.Status = 500
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(response)
		return
	}

	user.Password = passwordHashed

	_, err = models.InsertUser(user)
	if err != nil {
		log.Println(err)
		response.Content = err.Error()
		response.Status = 500
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(response)
		// http.Error(w, "Erro ao inserir usuário", http.StatusInternalServerError)
		return
	}
	
	response.Content = "Usuário criado com sucesso"
	response.Status = 201
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}
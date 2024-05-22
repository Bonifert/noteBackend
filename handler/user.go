package handler

import (
	"awesomeProject/dto"
	"awesomeProject/model"
	"awesomeProject/service"
	"awesomeProject/validator"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var loginInfo model.User
	if err := json.NewDecoder(r.Body).Decode(&loginInfo); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	if err := validator.Validate.Struct(loginInfo); err != nil {
		http.Error(w, "Invalid user", http.StatusBadRequest)
		return
	}

	jwt, err := service.Authenticate(loginInfo.Username, loginInfo.Password)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusBadRequest)
		return
	}

	w.Header().Set("Authorization", "Bearer "+jwt)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser dto.NewUser
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	if err := validator.Validate.Struct(newUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := service.CreateUser(&newUser)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	m := make(map[string]uint)
	m["id"] = id
	SendJSONResponse(w, m)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := service.GetUserById(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	SendJSONResponse(w, user)
}

func GetMe(w http.ResponseWriter, r *http.Request) {
	idStr, ok := r.Context().Value("id").(string)
	if !ok {
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return
	}
	fmt.Println(id)
	user, err := service.GetUserById(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	SendJSONResponse(w, user)
}

func SendJSONResponse(w http.ResponseWriter, data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
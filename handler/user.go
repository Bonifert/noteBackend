package handler

import (
	"awesomeProject/dto"
	"awesomeProject/service"
	"awesomeProject/validator"
	"encoding/json"
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
	"net/http"
	"strconv"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser dto.UsernameAndPassword
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	if err := validator.ValidateStruct(newUser); len(err) != 0 {
		w.WriteHeader(http.StatusBadRequest)
		sendJSONResponse(w, err)
		return
	}

	id, err := service.CreateUser(&newUser)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrDuplicated):
			http.Error(w, "Username is already taken", http.StatusConflict)
		default:
			http.Error(w, "failed to create the user", http.StatusInternalServerError)
		}
		return
	}
	m := make(map[string]uint)
	m["id"] = id
	w.WriteHeader(http.StatusCreated)
	sendJSONResponse(w, m)
}

func DeleteMe(w http.ResponseWriter, r *http.Request) {
	idStr, ok := r.Context().Value("id").(string)
	if !ok {
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return
	}
	id, _ := strconv.Atoi(idStr)

	err := service.DeleteUserById(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
	}
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

	sendJSONResponse(w, user)
}

func GetMe(w http.ResponseWriter, r *http.Request) {
	idStr, ok := r.Context().Value("id").(string)
	if !ok {
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return
	}
	id, _ := strconv.Atoi(idStr)

	user, err := service.GetUserById(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	sendJSONResponse(w, user)
}

func sendJSONResponse(w http.ResponseWriter, data interface{}) {
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

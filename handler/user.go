package handler

import (
	"awesomeProject/dto"
	"awesomeProject/service"
	"awesomeProject/validator"
	"encoding/json"
	"errors"
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
		errBody := dto.ErrorMessage{Errors: err, Status: http.StatusBadRequest}
		sendJSONResponse(w, errBody, http.StatusBadRequest)
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
	sendJSONResponse(w, m, http.StatusCreated)
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
		switch {
		case errors.Is(err, service.ErrNotFound):
			http.Error(w, "user not found", http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusConflict)
		}
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

	sendJSONResponse(w, user, http.StatusOK)
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

	sendJSONResponse(w, user, http.StatusOK)
}

func EditUsername(w http.ResponseWriter, r *http.Request) {
	idStr, ok := r.Context().Value("id").(string)
	if !ok {
		http.Error(w, "invalid token", http.StatusUnauthorized)
	}
	id, _ := strconv.Atoi(idStr)

	var editUsername dto.NewUsername
	if err := json.NewDecoder(r.Body).Decode(&editUsername); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	if err := validator.ValidateStruct(editUsername); len(err) != 0 {
		errBody := dto.ErrorMessage{Errors: err, Status: http.StatusBadRequest}
		sendJSONResponse(w, errBody, http.StatusBadRequest)
		return
	}

	err := service.EditUsernameById(uint(id), &editUsername)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrUnauthorized):
			http.Error(w, "unauthorized", http.StatusUnauthorized)
		case errors.Is(err, service.ErrNotFound):
			http.Error(w, "user not found", http.StatusNotFound)
		case errors.Is(err, service.ErrDuplicated):
			http.Error(w, "username is already taken", http.StatusConflict)
		default:
			http.Error(w, "failed to edit username", http.StatusInternalServerError)
		}
		return
	}
}

func EditPassword(w http.ResponseWriter, r *http.Request) {
	idStr, ok := r.Context().Value("id").(string)
	if !ok {
		http.Error(w, "invalid token", http.StatusUnauthorized)
	}
	id, _ := strconv.Atoi(idStr)
	var editPassword dto.NewPassword
	if err := json.NewDecoder(r.Body).Decode(&editPassword); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
	}
	if err := validator.ValidateStruct(editPassword); len(err) != 0 {
		errBody := dto.ErrorMessage{Errors: err, Status: http.StatusBadRequest}
		sendJSONResponse(w, errBody, http.StatusBadRequest)
	}

	err := service.EditPasswordById(uint(id), editPassword)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrUnauthorized):
			http.Error(w, "unauthorized", http.StatusUnauthorized)
		case errors.Is(err, service.ErrNotFound):
			http.Error(w, "user not found", http.StatusNotFound)
		default:
			http.Error(w, "failed to edit password", http.StatusInternalServerError)
		}
	}
}

func sendJSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

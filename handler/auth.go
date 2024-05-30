package handler

import (
	"awesomeProject/dto"
	"awesomeProject/service"
	"awesomeProject/validator"
	"encoding/json"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var loginInfo dto.UsernameAndPassword
	if err := json.NewDecoder(r.Body).Decode(&loginInfo); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	if err := validator.ValidateStruct(loginInfo); len(err) != 0 {
		errBody := dto.ErrorMessage{Errors: err, Status: http.StatusBadRequest}
		sendJSONResponse(w, errBody, http.StatusBadRequest)
		return
	}

	jwt, err := service.Authenticate(loginInfo.Username, loginInfo.Password)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Authorization", "Bearer "+jwt)
}

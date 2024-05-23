package handler

import (
	"awesomeProject/model"
	"awesomeProject/service"
	"awesomeProject/validator"
	"encoding/json"
	"net/http"
	"strconv"
)

func CreateNote(w http.ResponseWriter, r *http.Request) {
	var newNote model.Note
	if err := json.NewDecoder(r.Body).Decode(&newNote); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
	}
	idStr, ok := r.Context().Value("id").(string)
	if !ok {
		http.Error(w, "invalid token", http.StatusUnauthorized)
	}
	id, _ := strconv.Atoi(idStr)

	newNote.UserID = uint(id)

	if err := validator.Validate.Struct(newNote); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	noteId, err := service.CreateNote(&newNote)
	if err != nil {
		http.Error(w, "failed to create note", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
	m := make(map[string]uint)
	m["id"] = noteId
	SendJSONResponse(w, m)
}

func GetNote(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid note id", http.StatusBadRequest)
	}
	note, err := service.GetNote(uint(id))
	if err != nil {
		http.Error(w, "failed to get note", http.StatusInternalServerError)
	}
	SendJSONResponse(w, note)
}

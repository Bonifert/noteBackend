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

	if err := validator.ValidateStruct(newNote); len(err) != 0 {
		w.WriteHeader(http.StatusBadRequest)
		sendJSONResponse(w, err) // TODO content-type is text instead of application/json
		return
	}

	noteId, err := service.CreateNote(&newNote)
	if err != nil {
		http.Error(w, "failed to create note", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
	m := make(map[string]uint)
	m["id"] = noteId
	sendJSONResponse(w, m)
}

func GetNote(w http.ResponseWriter, r *http.Request) {
	userIdStr := r.PathValue("id")
	id, err := strconv.Atoi(userIdStr)
	if err != nil {
		http.Error(w, "invalid parameter", http.StatusBadRequest)
		return
	}
	note, err := service.GetNote(uint(id))
	if err != nil {
		http.Error(w, "note not found", http.StatusNotFound)
		return
	}
	userIdStr, ok := r.Context().Value("id").(string)
	if !ok {
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return
	}
	userId, _ := strconv.Atoi(userIdStr)

	if note.UserID != uint(userId) {
		http.Error(w, "access denied", http.StatusForbidden)
		return
	}
	sendJSONResponse(w, note)
}

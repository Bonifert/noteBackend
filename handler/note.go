package handler

import (
	"awesomeProject/dto"
	"awesomeProject/model"
	"awesomeProject/service"
	"awesomeProject/validator"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func CreateNote(w http.ResponseWriter, r *http.Request) {
	var newNote model.Note
	if err := json.NewDecoder(r.Body).Decode(&newNote); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}
	idStr, ok := r.Context().Value("id").(string)
	if !ok {
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return
	}
	id, _ := strconv.Atoi(idStr)
	newNote.UserID = uint(id)

	if err := validator.ValidateStruct(newNote); len(err) != 0 {
		w.WriteHeader(http.StatusBadRequest)
		sendJSONResponse(w, err)
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

func EditNoteContent(w http.ResponseWriter, r *http.Request) {
	var editNoteContent dto.EditNoteContent
	if err := json.NewDecoder(r.Body).Decode(&editNoteContent); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	if err := validator.ValidateStruct(editNoteContent); len(err) != 0 {
		w.WriteHeader(http.StatusBadRequest)
		sendJSONResponse(w, err)
		return
	}

	userIdStr, ok := r.Context().Value("id").(string)
	if !ok {
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return
	}
	userId, _ := strconv.Atoi(userIdStr)

	noteIdStr := r.PathValue("id")
	noteId, err := strconv.Atoi(noteIdStr)
	if err != nil {
		http.Error(w, "invalid note id", http.StatusBadRequest)
		return
	}

	err = service.EditNoteContent(uint(userId), uint(noteId), &editNoteContent)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrNotFound):
			http.Error(w, "note not found", http.StatusNotFound)
		case errors.Is(err, service.ErrForbidden):
			http.Error(w, "permission denied, the user don't have access to this resource", http.StatusForbidden)
		default:
			http.Error(w, "failed to edit note", http.StatusInternalServerError)
		}
		return
	}

}

func EditNoteTitle(w http.ResponseWriter, r *http.Request) {
	var editNoteTitle dto.EditNoteTitle
	if err := json.NewDecoder(r.Body).Decode(&editNoteTitle); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	if err := validator.ValidateStruct(editNoteTitle); len(err) != 0 {
		w.WriteHeader(http.StatusBadRequest)
		sendJSONResponse(w, err)
		return
	}

	userIdStr, ok := r.Context().Value("id").(string)
	if !ok {
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return
	}
	userId, _ := strconv.Atoi(userIdStr)

	noteIdStr := r.PathValue("id")
	noteId, err := strconv.Atoi(noteIdStr)
	if err != nil {
		http.Error(w, "invalid note id", http.StatusBadRequest)
		return
	}

	err = service.EditNoteTitle(uint(userId), uint(noteId), &editNoteTitle)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrNotFound):
			http.Error(w, "note not found", http.StatusNotFound)
		case errors.Is(err, service.ErrForbidden):
			http.Error(w, "permission denied", http.StatusForbidden)
		default:
			http.Error(w, "failed to edit note", http.StatusInternalServerError)
		}
		return
	}
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	userIdStr, ok := r.Context().Value("id").(string)
	if !ok {
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return
	}
	userId, _ := strconv.Atoi(userIdStr)

	noteIdStr := r.PathValue("id")
	noteId, err := strconv.Atoi(noteIdStr)
	if err != nil {
		http.Error(w, "invalid note id", http.StatusBadRequest)
		return
	}

	err = service.DeleteNote(uint(noteId), uint(userId))
	if err != nil {
		switch {
		case errors.Is(err, service.ErrForbidden):
			http.Error(w, "permission denied", http.StatusForbidden)
		case errors.Is(err, service.ErrNotFound):
			http.Error(w, "note not found", http.StatusNotFound)
		default:
			http.Error(w, "failed to delete note", http.StatusInternalServerError)
		}
		return
	}
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

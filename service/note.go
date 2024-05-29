package service

import (
	"awesomeProject/database"
	"awesomeProject/model"
	"errors"
)

var (
	ErrForbidden    = errors.New("access forbidden")
	ErrNotFound     = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
	ErrDuplicated   = errors.New("duplicated")
)

func CreateNote(newNote *model.Note) (uint, error) {
	result := database.DB.Create(&newNote)
	if result.Error != nil {
		return 0, result.Error
	}

	return newNote.ID, nil
}

func GetNote(noteId uint) (*model.Note, error) {
	var note model.Note
	result := database.DB.Where("ID = ?", noteId).First(&note)
	if result.Error != nil {
		return nil, result.Error
	}
	return &note, nil
}

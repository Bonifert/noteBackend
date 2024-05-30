package service

import (
	"awesomeProject/database"
	"awesomeProject/dto"
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

func DeleteNote(nodeId uint, userId uint) error {
	note, err := getNoteById(nodeId)
	if err != nil {
		return err
	}
	if note.UserID != userId {
		return ErrForbidden
	}
	database.DB.Delete(&note)
	return nil
}

func EditNoteContent(userId uint, noteId uint, editNoteContent *dto.EditNoteContent) error {
	note, err := getNoteById(noteId)
	if err != nil {
		return err
	}
	if note.UserID != userId {
		return ErrForbidden
	}
	note.Content = editNoteContent.NewContent
	result := database.DB.Save(&note)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func EditNoteTitle(userId uint, noteId uint, editNoteTitle *dto.EditNoteTitle) error {
	note, err := getNoteById(noteId)
	if err != nil {
		return err
	}
	if note.UserID != userId {
		return ErrForbidden
	}
	note.Title = editNoteTitle.NewTitle
	result := database.DB.Save(note)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetNote(noteId uint, userId uint) (*model.Note, error) {
	note, err := getNoteById(noteId)
	if err != nil {
		return nil, err
	}
	if note.UserID != userId {
		return nil, ErrForbidden
	}
	return note, nil
}

func getNoteById(noteId uint) (*model.Note, error) {
	var note model.Note
	result := database.DB.Where("id = ?", noteId).First(&note)
	if result.Error != nil {
		return nil, ErrNotFound
	}
	return &note, nil
}

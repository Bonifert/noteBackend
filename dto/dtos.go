package dto

type UsernameAndPassword struct {
	Username string `json:"username" validate:"required,gte=2,lte=20"`
	Password string `json:"password" validate:"required,gte=2,lte=20"`
}

type UserInfo struct {
	Username   string `json:"username"`
	NumOfNotes int    `json:"numOfNotes"`
}

type ErrorMessage struct {
	Errors any    `json:"errors"`
	Status string `json:"status"`
}

type NewUsername struct {
	NewUsername string `json:"username" validate:"required,gte=1,lte=20"`
	Password    string `json:"password" validate:"required"`
}

type NewPassword struct {
	NewPassword string `json:"newPassword" validate:"required,gte=1,lte=20"`
	Password    string `json:"password" validate:"required"`
}

type EditNoteTitle struct {
	NewTitle string `json:"newTitle" validate:"required,gte=1,lte=20"`
}

type EditNoteContent struct {
	NewContent string `json:"newContent" validate:"required"`
}

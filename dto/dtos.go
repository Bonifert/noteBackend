package dto

type NewUser struct {
	Username string `json:"username" validate:"required,gte=2,lte=255"`
	Password string `json:"password" validate:"required,gte=2,lte=255"`
}

type UserInfo struct {
	Username   string `json:"username"`
	NumOfNotes int    `json:"numOfNotes"`
}

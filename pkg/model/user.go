package model

type User struct {
	GormModel
	Username string `json:"username" validate:"required,gte=1,lte=255" gorm:"unique"`
	Password string `json:"-" validate:"required,gte=1,lte=255" gorm:"not null"`
	Notes    []Note `json:"-" gorm:"foreignKey:UserID"`
}

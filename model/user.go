package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username  string    `json:"username" validate:"required,gte=1,lte=255" gorm:"unique"`
	Password  string    `json:"password" validate:"required,gte=1,lte=255" gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime,notnull" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime,notnull" json:"updated_at"`
	Notes     []Note    `gorm:"foreignKey:UserID" json:"notes"`
	DeletedAt gorm.DeletedAt
}

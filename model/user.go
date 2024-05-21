package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username  string    `json:"username" validate:"required,gte=1,lte=255"`
	Password  string    `json:"password" validate:"required,gte=1,lte=255"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Notes     []Note    `gorm:"foreignKey:UserID" json:"notes"`
	DeletedAt gorm.DeletedAt
}

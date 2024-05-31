package model

import (
	"gorm.io/gorm"
	"time"
)

type GormModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime,notnull"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime,notnull"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

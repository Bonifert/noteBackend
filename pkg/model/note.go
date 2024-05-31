package model

type Note struct {
	GormModel
	Title   string `gorm:"size:40;not null" json:"title" validate:"required,gte=1,lte=40"`
	Content string `gorm:"type:text" json:"content" validate:"required,lte=255"`
	UserID  uint   `validate:"omitempty" json:"-"`
}

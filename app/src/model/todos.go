package model

import (
	"gorm.io/gorm"
	"time"
)

type Todos struct {
	ID        uint
	TITLE     string `gorm:"column:title;type:varchar(255)"`
	CONTENT   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

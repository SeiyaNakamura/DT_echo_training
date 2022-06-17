package models

import "time"

type SampleData struct {
	Name string
	Age  int
}

//article構造体
type Article struct {
	ID        uint `gorm:"primary_key" json:"id"`
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

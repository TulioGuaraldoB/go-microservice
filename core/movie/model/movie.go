package model

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	Title     string  `gorm:"size:255" json:"title"`
	Budget    float64 `json:"budget"`
	Genre     string  `gorm:"size:75" json:"genre"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

package models

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	ID          uint `gorm:"primaryKey"`
	FirstName   string
	LastName    string
	Photo       string
	DateOfBirth time.Time
}

package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName   string
	LastName    string
	Photo       string
	DateOfBirth time.Time
}

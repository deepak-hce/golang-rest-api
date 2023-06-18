package interfaces

import (
	"time"
)

type User struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Photo       string    `json:"photo"`
	DateOfBirth time.Time `json:"dateOfBirth"`
}

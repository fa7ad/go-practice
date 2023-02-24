package entities

import (
	"gorm.io/gorm"
	"time"
)

// Book Constructs your Book model under entities.
type Book struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primarykey"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// DeleteRequest struct is used to parse Delete Requests for Books
type DeleteRequest struct {
	gorm.Model
}

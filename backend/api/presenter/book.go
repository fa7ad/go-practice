package presenter

import (
	"github.com/gofiber/fiber/v2"
	"go-practice/pkg/entities"
)

// Book is the presenter object which will be passed in the response by Handler
type Book struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// BookSuccessResponse is the singular SuccessResponse that will be passed in the response by
// Handler
func BookSuccessResponse(data *entities.Book) *fiber.Map {
	book := Book{
		ID:     data.ID,
		Title:  data.Title,
		Author: data.Author,
	}
	return &fiber.Map{
		"success": true,
		"data":    book,
		"error":   nil,
	}
}

// BooksSuccessResponse is the list SuccessResponse that will be passed in the response by Handler
func BooksSuccessResponse(data *[]Book) *fiber.Map {
	return &fiber.Map{
		"success": true,
		"data":    data,
		"error":   nil,
	}
}

// BookErrorResponse is the ErrorResponse that will be passed in the response by Handler
func BookErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"success": false,
		"data":    "",
		"error":   err.Error(),
	}
}

package book

import (
	"fmt"
	"go-practice/api/presenter"
	"go-practice/pkg/entities"
	"gorm.io/gorm"
)

// Repository interface allows us to access the CRUD Operations in db.
type Repository interface {
	CreateBook(book *entities.Book) (*entities.Book, error)
	ReadBook() (*[]presenter.Book, error)
	UpdateBook(ID uint, book *entities.Book) (*entities.Book, error)
	DeleteBook(ID uint) error
	HandleError(res *gorm.DB) error

	AutoMigrate() error
}

type AutoMigrate func() error

type repository struct {
	DB *gorm.DB
}

// NewRepo is the single instance repo that is being created.
func NewRepo(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) AutoMigrate() error {
	return r.DB.AutoMigrate(&entities.Book{})
}

// CreateBook is a function that helps to create books
func (r *repository) CreateBook(book *entities.Book) (*entities.Book, error) {
	res := r.DB.Create(&book)
	err := r.HandleError(res)
	return book, err
}

// ReadBook is a function that helps to fetch books
func (r *repository) ReadBook() (*[]presenter.Book, error) {
	var bookEntities []entities.Book
	var books []presenter.Book
	res := r.DB.Find(&bookEntities)

	for _, entity := range bookEntities {
		book := presenter.Book{
			ID:     entity.ID,
			Title:  entity.Title,
			Author: entity.Author,
		}
		books = append(books, book)
	}

	err := r.HandleError(res)
	if err != nil {
		return nil, err
	}

	return &books, nil
}

// UpdateBook is a function that helps to update books
func (r *repository) UpdateBook(ID uint, book *entities.Book) (*entities.Book, error) {
	var existingBook entities.Book
	findRes := r.DB.Find(&existingBook, ID)
	err := r.HandleError(findRes)

	if err != nil {
		return nil, err
	}

	res := r.DB.Model(&existingBook).Updates(book)

	err = r.HandleError(res)
	if err != nil {
		return nil, err
	}

	return &existingBook, nil
}

// DeleteBook is a function that helps to delete books
func (r *repository) DeleteBook(ID uint) error {
	var book entities.Book

	findResponse := r.DB.Find(&book, ID)
	err := r.HandleError(findResponse)

	if err != nil {
		return err
	}

	deleteResponse := r.DB.Delete(&book)
	err = r.HandleError(deleteResponse)

	if err != nil {
		return err
	}
	return nil
}

// HandleError is a function for dealing with error responses from DB
func (r *repository) HandleError(res *gorm.DB) error {
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		err := fmt.Errorf("Error: %w", res.Error)
		return err
	}
	return nil
}

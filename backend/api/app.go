package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go-practice/api/routes"
	"go-practice/pkg/book"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, cancel, err := databaseConnection()
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}
	fmt.Println("Database connection success!")

	bookRepo := book.NewRepo(db)
	migrationError := bookRepo.AutoMigrate()
	if migrationError != nil {
		cancel()
		log.Fatal("Migration Failed, Error: $s", migrationError)
	}
	bookService := book.NewService(bookRepo)

	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to the go-practice mongo book shop!"))
	})
	api := app.Group("/api")
	routes.BookRouter(api, bookService)
	defer cancel()

	port := os.Getenv("PORT")
	listenUri := fmt.Sprintf(":%s", port)
	log.Fatal(app.Listen(listenUri))
}

func databaseConnection() (*gorm.DB, context.CancelFunc, error) {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	dbPath := os.Getenv("DB_PATH")
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		cancel()
		return nil, nil, err
	}

	return db, cancel, nil
}

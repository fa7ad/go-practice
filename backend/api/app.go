package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-practice/api/routes"
	"go-practice/pkg/book"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	db, cancel, err := databaseConnection()
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}
	fmt.Println("Database connection success!")
	bookCollection := db.Collection("books")
	bookRepo := book.NewRepo(bookCollection)
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

func databaseConnection() (*mongo.Database, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	bin, passErr := os.ReadFile("/run/secrets/db-password")
	if passErr != nil {
		cancel()
		return nil, nil, passErr
	}

	rawUser := os.Getenv("DB_USERNAME")
	rawPass := string(bin)

	user := url.QueryEscape(rawUser)
	pass := url.QueryEscape(rawPass)

	connectionUri := fmt.Sprintf("mongodb://%s:%s@db:27017/fiber", user, pass)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionUri).SetServerSelectionTimeout(5*time.Second))
	if err != nil {
		cancel()
		return nil, nil, err
	}
	db := client.Database("books")
	return db, cancel, nil
}

package main

import (
	"insider-case-study/internal/application"
	"insider-case-study/internal/domain/message"
	"insider-case-study/internal/infrastructure"
	"insider-case-study/internal/interfaces"
	"log"

	"insider-case-study/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	if err := config.LoadEnv(); err != nil {
		log.Fatal("cannot load .env file")
	}
	db, err := infrastructure.ConnectDatabase()
	if err != nil {
		log.Fatal("cannot connect database", err.Error())
	}
	//redis := infrastructure.ConnectRedis()
	// dependency
	messageRepo := message.NewMessageRepository(db)
	messageService := application.NewMessageService(messageRepo)

	app := fiber.New()
	interfaces.SetupRoutes(app, messageService)
	log.Println("Server is running on port 9005")
	if err := app.Listen(":9005"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

package main

import (
	"insider-case-study/internal/application"
	"insider-case-study/internal/domain/message"
	"insider-case-study/internal/infrastructure"
	"insider-case-study/internal/interfaces"
	"insider-case-study/pkg"
	"log"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"insider-case-study/config"

	"github.com/gofiber/fiber/v2"
)

// @title Message Service API
// @version 1.0
// @description This is a message service server with job management.
// @host localhost:9005
// @BasePath /api
func main() {
	if err := config.LoadEnv(); err != nil {
		log.Fatal("cannot load .env file")
	}
	db, err := infrastructure.ConnectDatabase()
	if err != nil {
		log.Fatal("cannot connect database", err.Error())
	}
	infrastructure.AutoMigrate(db)
	infrastructure.FillData("./internal/infrastructure/createdata.sql", db)
	redis, err := infrastructure.ConnectRedis()
	if err != nil {
		log.Println("Failed to connect to Redis:", err.Error())
	}
	// dependency
	messageRepo := message.NewMessageRepository(db)
	messageService := application.NewMessageService(messageRepo)

	jobService := pkg.NewJobService(messageService, redis)
	jobService.SendMessagesJob()

	app := fiber.New(fiber.Config{
		BodyLimit: 10 * 1024 * 1024,
	})
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${method} ${path} | Latency: ${latency} | IP: ${ip}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "UTC",
	}))

	app.Use(recover.New())
	app.Use(cors.New())
	interfaces.SetupRoutes(app, messageService, redis, jobService)

	log.Println("Server is running on port 9005")
	if err := app.Listen(":" + config.GetVar("PORT")); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

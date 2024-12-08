package interfaces

import (
	"insider-case-study/internal/application"
	"insider-case-study/internal/infrastructure"
	"insider-case-study/internal/interfaces/handler"
	"insider-case-study/pkg"

	_ "insider-case-study/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SetupRoutes(app *fiber.App, messageService application.MessageService, redisClient *infrastructure.RedisClient, jobService *pkg.JobService) {

	app.Get("/swagger/*", swagger.HandlerDefault)

	api := app.Group("/api")

	messageHandler := handler.NewMessageHandler(messageService, redisClient, jobService)

	api.Get("/messages", messageHandler.GetSentMessages())
	api.Post("/job-status", messageHandler.ChangeJobStatus())
	api.Get("/message/:id", messageHandler.GetSentMessage())
}

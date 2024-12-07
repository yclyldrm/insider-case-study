package interfaces

import (
	"insider-case-study/internal/application"
	"insider-case-study/internal/interfaces/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, messageService application.MessageService) {

	messageHandler := handler.NewMessageHandler(messageService)

	app.Get("/job-status", messageHandler.ChangeJobStatus())
	app.Get("/messages", messageHandler.GetSentMessages())
}

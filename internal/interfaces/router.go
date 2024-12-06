package interfaces

import (
	"insider-case-study/internal/application"
	"insider-case-study/internal/domain/message"
	"insider-case-study/internal/interfaces/handler"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	messageRepo := message.NewMessageRepository(db)
	messageService := application.NewMessageService(messageRepo)

	messageHandler := handler.NewMessageHandler(messageService)

	app.Get("/job-status", messageHandler.ChangeJobStatus())
	app.Get("/messages", messageHandler.GetSentMessages())
}

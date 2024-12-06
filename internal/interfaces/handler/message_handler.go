package handler

import (
	"insider-case-study/internal/application"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type MessageHandler struct {
	messageService application.MessageService
}

func NewMessageHandler(ms application.MessageService) *MessageHandler {
	return &MessageHandler{
		messageService: ms,
	}
}

// GetSentMessages retrieves all sent messages.
// @Summary Get sent messages
// @Description Retrieve all sent messages from the database.
// @Tags Messages
// @Success 200 {array} domain.Message
// @Failure 500 {object} map[string]string
// @Router /messages [get]
func (mh *MessageHandler) GetSentMessages() fiber.Handler {
	return func(c *fiber.Ctx) error {
		messages, err := mh.messageService.GetSentMessages()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(messages)
	}
}

// StartStopJob handles starting or stopping the message-sending job.
// @Summary Start or stop the message-sending job
// @Description Start or stop the periodic job that processes unsent messages every 2 minutes.
// @Tags Job
// @Param action query string true "Action (start/stop)"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /job [get]
func (mh *MessageHandler) ChangeJobStatus() fiber.Handler {
	return func(c *fiber.Ctx) error {

		jobStatus, err := strconv.ParseBool(c.Query("status"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		// after this, job should be run according to jobStatus
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"jobStatus": jobStatus})
	}
}

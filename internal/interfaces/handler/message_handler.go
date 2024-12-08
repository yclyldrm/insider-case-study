package handler

import (
	"insider-case-study/internal/application"
	"insider-case-study/internal/infrastructure"
	"insider-case-study/pkg"

	"github.com/gofiber/fiber/v2"
)

type MessageHandler struct {
	messageService application.MessageService
	redisClient    *infrastructure.RedisClient
	jobService     *pkg.JobService
}

func NewMessageHandler(ms application.MessageService, redisClient *infrastructure.RedisClient, js *pkg.JobService) *MessageHandler {
	return &MessageHandler{
		messageService: ms,
		redisClient:    redisClient,
		jobService:     js,
	}
}

// GetSentMessages godoc
// @Summary Get all sent messages
// @Description Retrieves all messages that have been sent
// @Tags messages
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "Returns sent messages"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /messages [get]
func (mh *MessageHandler) GetSentMessages() fiber.Handler {
	return func(c *fiber.Ctx) error {
		messages, err := mh.messageService.GetSentMessages()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(fiber.Map{"data": messages})
	}
}

// ChangeJobStatus godoc
// @Summary Change job status
// @Description Enable or disable the message sending job
// @Tags jobs
// @Accept json
// @Produce json
// @Param status body boolean true "Job status (true/false)"
// @Success 200 {object} map[string]string "Returns job status"
// @Failure 400 {object} map[string]string "Bad request"
// @Router /job-status [post]
func (mh *MessageHandler) ChangeJobStatus() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body struct {
			Status bool `json:"status"`
		}
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		mh.jobService.SetJobStatus(body.Status)

		go mh.jobService.SendMessagesJob()

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"jobStatus": body.Status})
	}
}

// GetSentMessage godoc
// @Summary Get a specific sent message
// @Description Retrieves a specific message by its ID
// @Tags messages
// @Accept json
// @Produce json
// @Param id path string true "Message ID"
// @Success 200 {object} message.MessageResponse "Returns the message"
// @Failure 400 {object} map[string]string "Bad request - missing ID"
// @Failure 404 {object} map[string]string "Message not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /messages/{id} [get]
func (mh *MessageHandler) GetSentMessage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		messageID := c.Params("id")
		if messageID == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "message ID is required"})
		}

		message, err := mh.redisClient.Get(messageID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		if message == nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "message not found"})
		}

		return c.JSON(fiber.Map{"data": message})
	}
}

package message

import (
	"fmt"

	"gorm.io/gorm"
)

type MessageRepository interface {
	GetUnsentMessages() ([]*Message, error)
	GetSentMessages() ([]*Message, error)
	UpdateMessage(*Message) error
}

type messageRepository struct {
	DB *gorm.DB
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &messageRepository{DB: db}
}

func (mr *messageRepository) GetUnsentMessages() ([]*Message, error) {
	var messages []*Message
	err := mr.DB.Where("status=?", false).Limit(2).Find(&messages).Error

	return messages, err
}

func (mr *messageRepository) GetSentMessages() ([]*Message, error) {
	var messages []*Message
	err := mr.DB.Where("status=?", true).Find(&messages).Error

	return messages, err
}

func (mr *messageRepository) UpdateMessage(message *Message) error {
	if message == nil {
		return fmt.Errorf("message cannot be nil")
	}

	return mr.DB.Model(message).Updates(map[string]interface{}{
		"status":     true,
		"sent_at":    message.SentAt,
		"message_id": message.MessageID,
	}).Error
}

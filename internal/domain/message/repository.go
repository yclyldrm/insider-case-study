package message

import (
	"time"

	"gorm.io/gorm"
)

type MessageRepository interface {
	GetUnsentMessages() ([]*Message, error)
	GetSentMessages() ([]*Message, error)
	ChangeStatusandSentTime(*Message) error
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

func (mr *messageRepository) ChangeStatusandSentTime(message *Message) error {
	return mr.DB.Model(message).Updates(map[string]interface{}{
		"status":  true,
		"sent_at": time.Now,
	}).Error
}

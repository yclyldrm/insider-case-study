package application

import (
	"insider-case-study/internal/domain/message"
)

type MessageService interface {
	GetSentMessages() ([]*message.Message, error)
	GetUnsentMessages() ([]*message.Message, error)
	ChangeStatusandSentTime(*message.Message) error
}

type messageService struct {
	messageRepo message.MessageRepository
}

func NewMessageService(repo message.MessageRepository) MessageService {
	return &messageService{messageRepo: repo}
}

func (ms *messageService) GetSentMessages() ([]*message.Message, error) {
	return ms.messageRepo.GetSentMessages()
}

func (ms *messageService) GetUnsentMessages() ([]*message.Message, error) {
	return ms.messageRepo.GetUnsentMessages()
}

func (ms *messageService) ChangeStatusandSentTime(message *message.Message) error {
	return ms.messageRepo.ChangeStatusandSentTime(message)
}

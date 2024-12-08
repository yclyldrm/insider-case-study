package message

import (
	"fmt"
	"time"
)

type Message struct {
	ID        uint      `gorm:"primaryKey"`
	Content   string    `gorm:"size:100"`
	Receiver  string    `json:"to"`
	MessageID string    `json:"messageID"`
	Status    bool      `json:"status"`
	SentAt    time.Time `json:"sentAt"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *Message) Validate() error {
	if len(m.Content) == 0 || len(m.Content) > 100 {
		return fmt.Errorf("message content must be between 1 and 100 characters")
	}
	if len(m.Receiver) == 0 {
		return fmt.Errorf("recipient phone number cannot be empty")
	}
	return nil
}

package message

import "time"

// MessageResponse represents the response structure for messages
type MessageResponse struct {
	ID        uint      `json:"id" example:"1"`
	Content   string    `json:"content" example:"Hello World"`
	Receiver  string    `json:"to" example:"+123545"`
	MessageID string    `json:"messageId,omitempty" example:"msg123"`
	Status    bool      `json:"status" example:"true"`
	SentAt    time.Time `json:"sentAt,omitempty"`
}

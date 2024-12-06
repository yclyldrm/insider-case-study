package api

import (
	"fmt"
	"insider-case-study/internal/domain/message"
)

func (c *Client) SendMessage(message message.Message) (string, error) {
	params := map[string]string{
		"message": message.Content,
		"to":      message.To,
	}
	resp := make(map[string]string)
	if err := c.sendRequest("POST", params, resp); err != nil {
		return "", fmt.Errorf("error occured while sending messages")
	}

	return resp["messageID"], nil
}

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"insider-case-study/config"
	"net/http"
	"time"
)

type Client struct {
	url        string
	httpClient *http.Client
}

func NewClient() *Client {
	url := config.GetVar("WEBHOOK_URL")

	client := &Client{
		url: url,
		httpClient: &http.Client{
			Timeout: 15 * time.Second,
		},
	}

	return client
}

func (c *Client) sendRequest(method string, params, response map[string]string) error {
	if c.url == "" {
		return fmt.Errorf("webhook URL is not configured")
	}

	payload, err := json.Marshal(params)
	if err != nil {
		return fmt.Errorf("failed to marshal params: %w", err)
	}

	request, err := http.NewRequest(method, c.url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		return fmt.Errorf("failed to send message, status code: %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("failed to read response; %s", err.Error())
	}

	return nil
}

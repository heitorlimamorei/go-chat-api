package services

import "fmt"

type ChatRequest struct {
	Name    string `json:"name"`
	OwnerId string `json:"ownerId"`
}

type MessageRequest struct {
	ChatId string `json:"chatId"`
	Text   string `json:"text"`
	Author string `json:"author"`
}

func (cr *ChatRequest) Validate() error {
	if cr.Name == "" {
		return fmt.Errorf("malformed request body (name): %v", cr.Name)
	}
	if cr.OwnerId == "" {
		return fmt.Errorf("malformed request body (ownerId): %v", cr.OwnerId)
	}
	return nil
}

func (cr *MessageRequest) Validate() error {
	if cr.ChatId == "" {
		return fmt.Errorf("malformed request body (chatId): %v", cr.ChatId)
	}
	if cr.Text == "" {
		return fmt.Errorf("malformed request body (text): %v", cr.Text)
	}
	if cr.Author == "" {
		return fmt.Errorf("malformed request body (author): %v", cr.Author)
	}
	return nil
}

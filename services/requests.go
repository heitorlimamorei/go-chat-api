package services

import "fmt"

type ChatRequest struct {
	Name    string `json:"name"`
	OwnerId string `json:"ownerId"`
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

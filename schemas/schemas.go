package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	Name    string `json:"name"`
	OwnerId string `json:"ownerId"`
}

type ChatResp struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	Name      string    `json:"name"`
	OwnerId   string    `json:"ownerId"`
}

type Message struct {
	gorm.Model
	ChatId string `json:"chatId"`
	Text   string `json:"text"`
	Author string `json:"author"`
}

type MessageResp struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	ChatId    string    `json:"chatId"`
	Text      string    `json:"text"`
	Author    string    `json:"author"`
}

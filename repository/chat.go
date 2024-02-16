package repository

import (
	"github.com/heitorlimamorei/go-chat-api/config"
	"github.com/heitorlimamorei/go-chat-api/schemas"
)

func GetChats() []schemas.Chat {
	db := config.GetPostgres()
	var chats []schemas.Chat
	db.Find(&chats)
	return chats

}

func CreateChat(newChat *schemas.Chat) error {
	err := db.Create(newChat).Error

	if err != nil {
		logger.ErrorF("Error while creating new chat: %v", err.Error())
		return err
	}

	return nil
}

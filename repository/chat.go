package repository

import (
	"fmt"

	"github.com/heitorlimamorei/go-chat-api/schemas"
)

func GetChats() []schemas.Chat {
	var chats []schemas.Chat
	db.Find(&chats)
	return chats

}

func GetChatById(id string) (res *schemas.Chat, err error) {
	chat := schemas.Chat{}

	resp := db.First(&chat, id)

	if resp.Error != nil {
		return nil, fmt.Errorf("user not found with id: %v", id)
	}

	return &chat, nil
}

func DeleteChat(id string) (rerr error) {
	chat := schemas.Chat{}

	resp := db.Delete(&chat, id)

	if resp.Error != nil {
		return fmt.Errorf("user not found with id: %v", id)
	}

	return nil
}

func CreateChat(newChat *schemas.Chat) error {
	err := db.Create(newChat).Error

	if err != nil {
		logger.ErrorF("Error while creating new chat: %v", err.Error())
		return err
	}

	return nil
}

func UpdateChat(updatedChat *schemas.Chat, id string) (updated *schemas.Chat, err error) {
	chat := schemas.Chat{}

	resp := db.First(&chat, id)

	if resp.Error != nil {
		return nil, fmt.Errorf("user not found with id: %v", id)
	}

	chat.Name = updatedChat.Name
	chat.OwnerId = updatedChat.OwnerId

	db.Save(&chat)

	return &chat, nil
}

package repository

import (
	"fmt"

	"github.com/heitorlimamorei/go-chat-api/schemas"
)

func CreateMessage(newMessage *schemas.Message) error {
	err := db.Create(newMessage).Error

	if err != nil {
		logger.ErrorF("Error while creating new message: %v", err.Error())
		return err
	}
	return nil
}

func GetMessages(chatId string) ([]*schemas.Message, error) {
	var messages []*schemas.Message
	resp := db.Where("chat_id =?", chatId).Find(&messages)

	if resp.Error != nil {
		return nil, fmt.Errorf("messages not found with this chatId: %v", chatId)
	}

	return messages, nil
}

func UpadateMessage(updatedMessage *schemas.Message, id string) (message *schemas.Message, err error) {
	message = &schemas.Message{}

	resp := db.First(message, id)

	if resp.Error != nil {
		return nil, fmt.Errorf("message not found with id: %v", id)
	}

	message.Text = updatedMessage.Text

	db.Save(message)

	return message, nil

}

func DeleteMessage(id string) error {
	message := &schemas.Message{}

	resp := db.First(message, id)

	if resp.Error != nil {
		return fmt.Errorf("message not found with id: %v", id)
	}

	db.Delete(message, id)

	return nil

}

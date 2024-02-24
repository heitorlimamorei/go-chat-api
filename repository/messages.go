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

func getQuery(author string, chatId string) map[string]interface{} {
	fmt.Println(chatId)
	if chatId != "" {
		return map[string]interface{}{"author": author, "chat_id": chatId}
	}
	return map[string]interface{}{"author": author}
}

func GetMessagesByAuthor(author string, chatId string) ([]*schemas.Message, error) {
	var messages []*schemas.Message

	resp := db.Where(getQuery(author, chatId)).Find(&messages)

	if resp.Error != nil {
		return nil, fmt.Errorf("error while finding for messages with the author: %v; error: %v", author, resp.Error.Error())
	}

	if len(messages) == 0 {
		return nil, fmt.Errorf("no messages found with the author: %v", author)
	}

	return messages, nil
}

package services

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/heitorlimamorei/go-chat-api/repository"
	"github.com/heitorlimamorei/go-chat-api/schemas"
)

func CheckIfChatExists(chatId string) bool {
	_, err := repository.GetChatById(chatId)

	return err == nil
}

func GetMessages(c *gin.Context) {
	chatId := c.Param("chat_id")

	if chatId == "" {
		sendError(c, http.StatusBadRequest, "Chat id is required")
		return
	}

	messages, err := repository.GetMessages(chatId)

	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSucess(c, "Get Messages by chat_id", messages)

}

func GetMessagesByQueries(c *gin.Context) {
	email := c.Query("author")
	chat_id := c.Query("chat_id")

	if email != "" {
		messages, err := repository.GetMessagesByAuthor(email, chat_id)

		if err != nil {
			sendError(c, http.StatusInternalServerError, err.Error())
			return
		}

		sendSucess(c, "Messages retrieved successfully", messages)
		return
	}

	sendError(c, http.StatusInternalServerError, "Invalid email recived(query params)!")
}

func CreateMessage(c *gin.Context) {
	req := MessageRequest{}
	c.BindJSON(&req)
	fmt.Println(req)
	if !CheckIfChatExists(req.ChatId) {
		sendError(c, http.StatusBadRequest, "Invalid chatId provided")
		return
	}

	if err := req.Validate(); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	newMessage := schemas.Message{
		ChatId: req.ChatId,
		Text:   req.Text,
		Author: req.Author,
	}

	err := repository.CreateMessage(&newMessage)

	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSucess(c, "Create Message", nil)
}

package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/heitorlimamorei/go-chat-api/repository"
	"github.com/heitorlimamorei/go-chat-api/schemas"
)

func GetChats(c *gin.Context) {
	chats := repository.GetChats()
	sendSucess(c, "Get chats", chats)
}

func GetChat(c *gin.Context) {
	chatId := c.Param("id")

	if chatId == "" {
		sendError(c, http.StatusBadRequest, "Chat id is required")
		return
	}

	chat, err := repository.GetChatById(chatId)

	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSucess(c, "Get Chat by Id", *chat)
}

func CreateChat(c *gin.Context) {
	req := ChatRequest{}

	c.BindJSON(&req)

	if err2 := req.Validate(); err2 != nil {
		logger.ErrorF("Error during the validation: %v", err2.Error())
		sendError(c, http.StatusBadRequest, err2.Error())
		return
	}

	newChat := schemas.Chat{
		Name:    req.Name,
		OwnerId: req.OwnerId,
	}

	err := repository.CreateChat(&newChat)

	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSucess(c, "Salving chat", newChat)
}

func UpdateChat(c *gin.Context) {
	req := ChatRequest{}
	chatId := c.Param("id")

	if chatId == "" {
		sendError(c, http.StatusBadRequest, "Chat id is required")
		return
	}

	c.BindJSON(&req)

	if err2 := req.Validate(); err2 != nil {
		logger.ErrorF("Error during the validation: %v", err2.Error())
		sendError(c, http.StatusBadRequest, err2.Error())
		return
	}

	chatToUpdated := schemas.Chat{
		Name:    req.Name,
		OwnerId: req.OwnerId,
	}

	updatedChat, err := repository.UpdateChat(&chatToUpdated, chatId)

	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSucess(c, "Salving chat", updatedChat)
}

func DeleteChat(c *gin.Context) {
	chatId := c.Param("id")

	if chatId == "" {
		sendError(c, http.StatusBadRequest, "Chat id is required")
		return
	}

	err := repository.DeleteChat(chatId)

	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSucess(c, "Chat deleted sucefully", nil)
}

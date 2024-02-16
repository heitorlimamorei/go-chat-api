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

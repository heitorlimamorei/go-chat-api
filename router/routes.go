package router

import (
	"github.com/gin-gonic/gin"
	"github.com/heitorlimamorei/go-chat-api/services"
)

func initializeRoutes(r *gin.Engine) {
	r.GET("/chat", services.GetChats)
	r.POST("/chat", services.CreateChat)
}

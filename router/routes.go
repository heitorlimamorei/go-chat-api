package router

import (
	"github.com/gin-gonic/gin"
	"github.com/heitorlimamorei/go-chat-api/services"
)

func initializeRoutes(r *gin.Engine) {
	r.GET("/chat", services.GetChats)
	r.GET("/chat/:id", services.GetChat)
	r.DELETE("/chat/:id", services.DeleteChat)
	r.POST("/chat", services.CreateChat)
	r.PUT("/chat/:id", services.UpdateChat)
}

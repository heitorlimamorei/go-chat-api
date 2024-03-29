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
	r.GET("/messages/:chat_id", services.GetMessages)
	r.GET("/messages", services.GetMessagesByQueries)
	r.POST("/messages", services.CreateMessage)
	r.GET("/users", services.GetUsers)
	r.GET("/users/:user_id", services.GetUserById)
	r.PUT("/users/:user_id", services.UpdateUser)
	r.POST("/users", services.CreateUser)
	r.DELETE("/users/:user_id", services.DeleteUser)
}

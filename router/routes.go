package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {
	r.GET("/chat", func(req *gin.Context) {
		req.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})
}

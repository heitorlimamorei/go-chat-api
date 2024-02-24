package router

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.Default()
	initializeRoutes(router)
	port := os.Getenv("PORT")

	if port == "" {
		port = ":8080"
	}

	router.Run(port)
}

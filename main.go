package main

import (
	"github.com/heitorlimamorei/go-chat-api/config"
	"github.com/heitorlimamorei/go-chat-api/repository"
	"github.com/heitorlimamorei/go-chat-api/router"
	"github.com/heitorlimamorei/go-chat-api/services"
)

func main() {
	config.InitDB()
	repository.Init()
	services.Init()
	router.Init()
}

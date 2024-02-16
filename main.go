package main

import (
	"github.com/heitorlimamorei/go-chat-api/config"
	"github.com/heitorlimamorei/go-chat-api/router"
)

func main() {
	go config.InitDB()
	router.Initialize()
}

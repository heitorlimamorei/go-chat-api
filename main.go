package main

import (
	"fmt"

	"github.com/heitorlimamorei/go-chat-api/config"
	"github.com/heitorlimamorei/go-chat-api/router"
)

func main() {
	fmt.Println("Starting")
	config.InitDB()
	router.Initialize()
}

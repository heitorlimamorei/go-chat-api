package services

import (
	"github.com/heitorlimamorei/go-chat-api/config"
)

var (
	logger *config.Logger
)

func Init() {
	logger = config.GetLogger("SERVICES")
}

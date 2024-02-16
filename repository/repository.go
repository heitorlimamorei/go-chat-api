package repository

import (
	"github.com/heitorlimamorei/go-chat-api/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func Init() {
	db = config.GetPostgres()
	logger = config.GetLogger("REPOSITORY")
}

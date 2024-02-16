package config

import "gorm.io/gorm"

var (
	DB *gorm.DB
)

func InitDB() {
	err := initPostgres()
	if err != nil {
		panic(err)
	}
}

func GetLogger(prefix string) *Logger {
	logger := NewLogger(prefix)
	return logger
}

func GetPostgres() *gorm.DB {
	return DB
}

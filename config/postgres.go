package config

import (
	"fmt"

	"github.com/heitorlimamorei/go-chat-api/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDsn() string {
	host := "kesavan.db.elephantsql.com"
	user := "ekpxyioz"
	password := "kplEnPdhnvQL7eWeWVF89nOEnq_pLVpM"

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v  sslmode=disable", host, user, password, user)

	return dsn
}
func initPostgres() error {
	var err error

	logger := GetLogger("POSTGRES")

	dsn := getDsn()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.ErrorF("Faling to connect to the database: %v", err)
		return err
	}

	logger.Info("Connected to database")

	DB = db

	err = DB.AutoMigrate(&schemas.Chat{}, &schemas.Message{})

	if err != nil {
		logger.ErrorF("Faling to migrate the database: %v", err)
		return err
	}

	return nil
}

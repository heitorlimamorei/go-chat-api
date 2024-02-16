package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDsn() string {
	host := "kesavan.db.elephantsql.com"
	user := "ekpxyioz"
	password := "kplEnPdhnvQL7eWeWVF89nOEnq_pLVpM"

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=9920 sslmode=disable TimeZone=amazon-web-services::sa-east-1", host, user, password, user)

	return dsn
}
func initPostgres() {

	dsn := getDsn()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	DB = db
}

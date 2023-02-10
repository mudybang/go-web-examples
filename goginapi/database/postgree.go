package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConnectPostgree() {
	var err error
	host := os.Getenv("DB_HOST_POSTGRE")
	username := os.Getenv("DB_USER_POSTGRE")
	password := os.Getenv("DB_PASSWORD_POSTGRE")
	databaseName := os.Getenv("DB_NAME_POSTGRE")
	port := os.Getenv("DB_PORT_POSTGRE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, username, password, databaseName, port)
	Database1, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}
}

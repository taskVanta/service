package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	host := os.Getenv("POSTGRESHOST")
	user := os.Getenv("POSTGRESUSER")
	password := os.Getenv("POSTGRESPASSWORD")
	db := os.Getenv("POSTGRESDB")
	port := os.Getenv("POSTGRESPORT")
	sslmode := os.Getenv("SSLMODE")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, db, port, sslmode,
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	DB = database
}

package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// This function will handle the database connection logic
	// For example, you can use GORM or any other ORM to connect to your database
	// Here is a placeholder for the actual implementatio

	dsn := "host=dpg-d3kf0q49c44c73ae5vlg-a.singapore-postgres.render.com user=taskvanta password=ucVevp8si0yE75kK5uzYwA0jtgcNy7AP dbname=taskvanta_1bi6 port=5432 sslmode=require"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	DB = database
}

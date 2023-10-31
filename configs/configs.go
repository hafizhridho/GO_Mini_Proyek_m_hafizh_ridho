package configs

import (
	"latihan/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Loadenv() {
	err := godotenv.Load()
	if err != nil {
		panic("error loading .env")
	}
}

func InitDb() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		panic("DATABASE_URL is not set in your environment variables")
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}

	Migration()
}

func Migration() {
	DB.AutoMigrate(&models.User{}, &models.List{}, &models.Tugas{})
}

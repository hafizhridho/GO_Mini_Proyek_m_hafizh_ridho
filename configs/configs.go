package configs

import (
	"fmt"
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
	host := os.Getenv("PGHOST")
    user := os.Getenv("PGUSER")
    password := os.Getenv("PGPASSWORD")
    database := os.Getenv("PGDATABASE")
    port := os.Getenv("PGPORT")
    sslmode := os.Getenv("PGSSLMODE")
    timezone := os.Getenv("PGTIMEZONE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
	host, user, password, database, port, sslmode, timezone)
	
	var err error
	DB, err =	gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed connect to database")
	}
	Migration()
}

func Migration() {
	DB.AutoMigrate(&models.User{}, &models.List{}, &models.Tugas{})
	
 }


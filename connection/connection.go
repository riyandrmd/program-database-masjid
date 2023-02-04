package connection

import (
	"fmt"
	"masjiddotexe/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"log"
	"os"
)

func ConnectToDb() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		log.Fatalln("Can't Load Env")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable", dbUser, dbPass, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Can't Connect To Database")
	}

	db.AutoMigrate(&models.JadwalSholat{}, &models.Ktg_Inventaris{}, &models.Inventaris{}, &models.InfaqMasuk{})
	return db

}

package database

import (
	"fmt"
	"galaxy-weather/model"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var db *gorm.DB

func Client() *gorm.DB {
	if db != nil {
		return db
	}
	var err error

	_ = godotenv.Load()

	dbUsername := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbUsername, dbName, dbPassword)

	db, err = gorm.Open("postgres", dbURI)
	if err != nil {
		logrus.Errorf("ERROR: %s", err.Error())
		panic("failed to connect database")
	}
	// defer db.Close()

	db.LogMode(false)
	db.DB().SetMaxIdleConns(1000)
	Migrate(db)
	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Planet{})
	db.AutoMigrate(&model.Weather{})
	db.AutoMigrate(&model.Position{})
}

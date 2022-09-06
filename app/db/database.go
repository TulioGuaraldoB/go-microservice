package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

var db *gorm.DB

func StartDb(dbCfg DbConfig) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbCfg.User,
		dbCfg.Password,
		dbCfg.Host,
		dbCfg.Port,
		dbCfg.Name,
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		errMessage := fmt.Sprintf("Failed to connect to MySql Server! %s", err.Error())
		log.Fatal(errMessage)
	}

	db = database
}

func OpenConnection() *gorm.DB {
	StartDb(DbConfig{
		User:     os.Getenv("APP_DB_USER"),
		Password: os.Getenv("APP_DB_PASSWORD"),
		Host:     os.Getenv("APP_DB_HOST"),
		Port:     os.Getenv("APP_DB_PORT"),
		Name:     os.Getenv("APP_DB_NAME"),
	})

	return db
}

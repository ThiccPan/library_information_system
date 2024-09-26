package config

import (
	"fmt"
	"log"
	"os"

	"github.com/thiccpan/library_information_system/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	log.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to db")
	}
	dbAutoMigrate(db)
	return db
}

func dbAutoMigrate(db *gorm.DB) *gorm.DB {
	err := db.AutoMigrate(
		&entity.User{},
		&entity.Role{},
	)
	if err != nil {
		log.Fatal("migration error")
	}
	return db
}

package config

import (
	"fmt"
	"log"
	"os"

	"github.com/thiccpan/library_information_system/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	runSeeding(db)
	return db
}

func dbAutoMigrate(db *gorm.DB) *gorm.DB {
	err := db.AutoMigrate(&entity.Role{}, &entity.User{})
	if err != nil {
		log.Fatal("migration error", err)
	}

	err = db.AutoMigrate(&entity.Author{}, &entity.Topic{})
	if err != nil {
		log.Fatal("migration error", err)
	}

	err = db.AutoMigrate(&entity.Book{}, &entity.LoanStatus{})
	if err != nil {
		log.Fatal("migration error", err)
	}

	err = db.AutoMigrate(&entity.Loan{})
	if err != nil {
		log.Fatal("migration error", err)
	}

	return db
}

func runSeeding(db *gorm.DB) *gorm.DB {
	statusMigrationData := []*entity.LoanStatus{
		{Id: 1, Status: entity.BORROWED},
		{Id: 2, Status: entity.RETURNED},
	}
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"status"}),
	}).Create(statusMigrationData)
	return db
}

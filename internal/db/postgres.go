package db

import (
	"fmt"
	"log"
	"os"
	"mess-manager/internal/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PgDb *gorm.DB

func ConnectPostgresDB() {
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	dsn := os.Getenv("DATABASE_URL")
	fmt.Println("Database URL:", dsn)
	if dsn == "" {
		log.Fatal("DATABASE_URL not set in environment")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	PgDb = db
}

func AutoMigrateModels() {

	if err := PgDb.AutoMigrate(&model.User{}); err != nil {
		log.Fatal("failed to auto-migrate User Model: ", err)
	} else if err := PgDb.AutoMigrate(&model.Meal{}); err != nil {
		log.Fatal("failed to auto-migrate Meal Model: ", err)
	} else if err := PgDb.AutoMigrate(&model.Expense{}); err != nil {
		log.Fatal("failed to auto-migrate Expense Model: ", err)
	}else if err := PgDb.AutoMigrate(&model.ExpenseType{}); err != nil {
		log.Fatal("failed to auto-migrate Expense Model: ", err)
	}else if err := PgDb.AutoMigrate(&model.MealType{}); err != nil {
		log.Fatal("failed to auto-migrate Expense Model: ", err)
	}else if err := PgDb.AutoMigrate(&model.Mess{}); err != nil {
		log.Fatal("failed to auto-migrate Expense Model: ", err)
	}else if err := PgDb.AutoMigrate(&model.MessMember{}); err != nil {
		log.Fatal("failed to auto-migrate Expense Model: ", err)
	}else {
		log.Println("Database migrated successfully")
	}
}

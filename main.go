package main

import (
	"fmt"
	"go-article/src/business/domain"
	"go-article/src/business/entity"
	"go-article/src/business/usecase"
	"go-article/src/handler/rest"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := initDb()

	d := domain.Init(db)

	uc := usecase.Init(d)

	r := rest.Init(uc)

	r.Run()
}

func initDb() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&entity.Article{}); err != nil {
		panic(err)
	}

	return db
}

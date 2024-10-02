package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/muyi2905/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDb() {
	dsn := os.Getenv("DSN")
	if dsn == "" {
		panic("dsn envireonment variable is not set")
	}
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database", err)
	}
	fmt.Println("Database connection successful")
}

func main() {
	initDb()
	db.AutoMigrate(models.Todo{}, models.User{})
	r := gin.Default()
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("error starting server")
	}

}

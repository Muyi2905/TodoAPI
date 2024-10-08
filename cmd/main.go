package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/muyi2905/models"
	"github.com/muyi2905/routes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDb() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

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
	routes.RegisterRoutes(r, db)
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("error starting server")
	}

}

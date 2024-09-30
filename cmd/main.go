package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	r := gin.Default()
	err := r.Run()
	if err != nil {
		fmt.Println("error starting server")
	}
	db.AutoMigrate(models.Todo{}, models.User{})

	validate := validator.New()
	user := models.User{
		Email:    "obaremimuyiwa@gmail.com",
		Password: "74ydbhduudu",
	}
	if err := validate.Struct(user); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.StructField(), err.Tag(), err.Param())
		}
	} else {
		fmt.Println("validation passed")
	}

}

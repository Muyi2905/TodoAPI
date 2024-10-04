package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muyi2905/models"
	"gorm.io/gorm"
)

func GetTodo(c *gin.Context, db *gorm.DB) {
	var todos []models.Todo
	if err := db.Find(&todos); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"err": err.Error})
	}

}

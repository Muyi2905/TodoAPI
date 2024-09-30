package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(db *gorm.DB) {
	r := gin.Default()
	r.GET("/users", controllers.GetUsers)
}

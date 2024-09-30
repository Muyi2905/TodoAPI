package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.Engine, db *gorm.DB) {
	r.Group("/api/v1/")
	{
		r.GET("/users", controllers.GetUsers)
		r.GET("/user", controllers.GetUserById)
		r.POST("/user", controllers.CreateUser)
		r.PUT("/user", controllers.UpdateUser)
		r.DELETE("/user", controllers.DeleteUser)
	}
}

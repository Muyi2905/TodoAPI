package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.Engine, db *gorm.DB) {
	r.Group("/api/v1/users")
	{
		r.GET("/", controllers.GetUsers)
		r.GET("/:id", controllers.GetUserById)
		r.POST("/", controllers.CreateUser)
		r.PUT("/:id", controllers.UpdateUser)
		r.DELETE("/:id", controllers.DeleteUser)
		r.POST("/signup", controllers.Signup)
		r.POST("/login", controllers.Login)
	}
}

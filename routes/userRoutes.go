package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/muyi2905/controllers"
	"github.com/muyi2905/middleware"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.Engine, db *gorm.DB) {
	userGroup := r.Group("/api/v1/users")
	userGroup.Use(middleware.AuthMiddleware())
	{
		userGroup.GET("", func(c *gin.Context) { controllers.GetUsers(c, db) })
		userGroup.GET("/:id", func(c *gin.Context) { controllers.GetUserById(c, db) })
		userGroup.POST("", func(c *gin.Context) { controllers.CreateUser(c, db) })
		userGroup.PUT("/:id", func(c *gin.Context) { controllers.UpdateUser(c, db) })
		userGroup.DELETE("/:id", func(c *gin.Context) { controllers.DeleteUser(c, db) })
	}

	authGroup := r.Group("/api/v1/auth")
	{
		authGroup.POST("/signup", func(c *gin.Context) { controllers.Signup(c, db) })
		authGroup.POST("/login", func(c *gin.Context) { controllers.Login(c, db) })
	}
}

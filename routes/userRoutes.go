package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/muyi2905/controllers"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.Engine, db *gorm.DB) {
	r.Group("/api/v1/users")
	{
		r.GET("/", func(c *gin.Context) { controllers.GetUsers(c, db) })
		r.GET("/:id", func(c *gin.Context) { controllers.GetUserById(c, db) })
		r.POST("/", func(c *gin.Context) { controllers.CreateUser(c, db) })
		r.PUT("/:id", func(c *gin.Context) { controllers.UpdateUser(c, db) })
		r.DELETE("/:id", func(c *gin.Context) { controllers.DeleteUser(c, db) })
		r.POST("/signup", func(c *gin.Context) { controllers.Signup(c, db) })
		r.POST("/login", func(c *gin.Context) { controllers.Login }(c, db))
	}
}

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/muyi2905/controllers"
	"gorm.io/gorm"
)

func TodoRoutes(r *gin.Engine, db *gorm.DB) {
	r.Group("/api/v1/todo")
	{
		r.GET("/", func(c *gin.Context) { controllers.GetTodo(c, db) })
		r.GET("/:id", func(c *gin.Context) { controllers.GetTodoById(c, db) })
		r.POST("/", func(c *gin.Context) { controllers.CreateTodo(c, db) })
		r.PUT("/:id", func(c *gin.Context) { controllers.UpdateTodo(c, db) })
		r.DELETE("/:id", func(c *gin.Context) { controllers.DeleteTodo(c, db) })
	}
}

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/muyi2905/controllers"
	"gorm.io/gorm"
)

func TodoRoutes(r *gin.Engine, db *gorm.DB) {
	todoGroup := r.Group("/api/v1/todo")
	{
		todoGroup.GET("", func(c *gin.Context) { controllers.GetTodo(c, db) })
		todoGroup.GET("/:id", func(c *gin.Context) { controllers.GetTodoById(c, db) })
		todoGroup.POST("", func(c *gin.Context) { controllers.CreateTodo(c, db) })
		todoGroup.PUT("/:id", func(c *gin.Context) { controllers.UpdateTodo(c, db) })
		todoGroup.DELETE("/:id", func(c *gin.Context) { controllers.DeleteTodo(c, db) })
	}
}

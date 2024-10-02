package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/muyi2905/controllers"
	"gorm.io/gorm"
)

func TodoRoutes(r *gin.Engine, db *gorm.DB) {
	r.Group("/api/v1/todo")
	{
		r.GET("/", func(ctx *gin.Context) { controllers.GetTodo(ctx, db) })
		r.GET("/:id", controllers.GetTodoById)
		r.POST("/", controllers.CreateTask)
		r.PUT("/:id", controllers.UpdateTask)
		r.DELETE("/:id", controllers.DeleteTask)
	}
}

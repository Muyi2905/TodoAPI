package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TodoRoutes(r *gin.Engine, db *gorm.DB) {
	r.Group("/api/v1/todo")
	{
		r.GET("/", controllers.GetTodo)
		r.GET("/:id", controllers.GetTodoById)
		r.POST("/", controllers.CreateTask)
		r.PUT("/:id", controllers.UpdateTask)
		r.DELETE("/:id", controllers.DeleteTask)
	}
}

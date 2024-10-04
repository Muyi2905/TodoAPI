package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muyi2905/models"
	"gorm.io/gorm"
)

func GetTodo(c *gin.Context, db *gorm.DB) {
	userId, _ := c.Get("user_id")
	var todos []models.Todo
	if err := db.Where("user_id = ?", userId).Find(&todos); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "error getting todos"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"todos": todos})

}

func CreateTodo(c *gin.Context, db *gorm.DB) {

	var todos models.Todo
	if err := c.ShouldBindJSON(&todos); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	userId, _ := c.Get("user_id")
	todos.ID = userId.(uint)

	if err := db.Create(&todos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "error creating todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"todos": todos})

}

func UpdateTodo(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	userId, _ := c.Get("user_id")
	var updatedTodo models.Todo
	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var todo models.Todo
	if err := db.Where("id = ?", id).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	if todo.UserId != userId {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to update this todo"})
		return
	}

	if err := db.Model(&todo).Updates(updatedTodo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo updated successfully", "todo": todo})
}

func GetTodoById(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	userId, _ := c.Get("userId")

	var todo models.Todo
	if err := db.Where("id = ? AND user_id = ?", id, userId).First(&todo).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todo"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"todo": todo})
}

func DeleteTodo(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	userId, _ := c.Get("user_id")

	var todo models.Todo

	if err := db.Where("id = ? AND user_id = ?", id, userId).First(&todo).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todo"})
		}
		return
	}

	if err := db.Delete(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}

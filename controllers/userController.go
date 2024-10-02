package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/muyi2905/models"
	"gorm.io/gorm"
)

var validate = validator.New()

func GetUsers(c *gin.Context, db *gorm.DB) {
	var users []models.User

	if err := db.Find(&users).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func CreateUser(c *gin.Context, db *gorm.DB) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	if err := validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	var existingUser models.User
	err := db.Where("email = ?", user.Email).First(&existingUser).Error
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "user alredy exist "})
	}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create user",
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created sucessful",
	})
}

func GetUserById(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")

	var user models.User

	if err := db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "user not found"})
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch user"})
	}
	c.JSON(http.StatusOK, gin.H{
		"user": "user",
	})
}

func UpdateUser(c *gin.Context, db *gorm.DB) {

	id := c.Param("id")

	var updatedUser models.User
	err := c.ShouldBindJSON(&updatedUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error()})
		return
	}

	if err := validate.Struct(updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := db.Find(&updatedUser, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"err": "user not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		}
	}

	if err := db.Model(&user).Updates(&updatedUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "user updated successfully"})
}

func DeleteUser(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")

	var user models.User

	if err := db.First(user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"err": "user not found"})
		}
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}

	if err := db.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}

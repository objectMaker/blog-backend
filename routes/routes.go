package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/objectMaker/blog-backend/db"
	"github.com/objectMaker/blog-backend/models"
)

func CreateUser(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "missing name",
		})
		return
	}
	user := models.User{
		Name: name,
	}
	result := db.DB.Create(&user)

	if result.Error != nil {
		log.Fatalf("failed to create user: %v", result.Error)
	}
	c.JSON(200, gin.H{
		"message": "success",
		"body":    user,
	})
}

func GetUserList(c *gin.Context) {
	users := models.User{}
	result := db.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error,
		})
		log.Fatal("failed to get user list: %w", result.Error)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"body":    users,
	})
}

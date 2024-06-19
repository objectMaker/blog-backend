package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/objectMaker/blog-backend/db"
	"github.com/objectMaker/blog-backend/models"
)

func Ping(c *gin.Context) {

	result := db.DB.Create(&models.User{
		Name: "new user",
	})
	if result.Error != nil {
		log.Fatalf("failed to create user: %v", result.Error)
	}
	c.JSON(200, gin.H{
		"message": "success",
	})
}

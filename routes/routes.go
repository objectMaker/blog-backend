package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/objectMaker/blog-backend/db"
	"github.com/objectMaker/blog-backend/models"
)

func CreateUser(c *gin.Context) {
	userInfo := models.User{}
	c.BindJSON(&userInfo)

	if userInfo.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "missing userName",
		})
		return
	}

	if len(userInfo.Username) <= 5 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "userName must longer than five characters",
		})
		return
	}
	user := models.User{
		Username: userInfo.Username,
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
	var users []models.User
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

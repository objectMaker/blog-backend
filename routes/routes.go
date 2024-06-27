package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/objectMaker/blog-backend/db"
	"github.com/objectMaker/blog-backend/jwt"
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
	token, err := jwt.New(user.Username)
	if err != nil {
		log.Fatal("failed to create token: %w", err)
	}

	c.SetCookie("token", token, 3600, "/", "127.0.0.1", false, true)
	type ResResult struct {
		Token string `json:"token"`
		models.User
	}
	payload, err := jwt.ParseToken(token)
	if err != nil {
		log.Fatal("failed to parse token: %w", err)
	}
	fmt.Println(payload.Username, "username")
	fmt.Println(payload.Exp, "exp")

	c.JSON(200, gin.H{
		"message": "success",
		"body": ResResult{
			Token: token,
			User:  user,
		},
	})
}

func GetUserList(c *gin.Context) {
	fmt.Println(c.GetHeader("token"), "tmd header token")
	var users []models.User
	result := db.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error,
		})
		log.Fatal("failed to get user list: %w", result.Error)
	}
	c.SetCookie("token", "token+++++++++++++", 36000, "/", "127.0.0.1", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"body":    users,
	})
}

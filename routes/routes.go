package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/objectMaker/blog-backend/db"
	"github.com/objectMaker/blog-backend/jwt"
	"github.com/objectMaker/blog-backend/models"
	"github.com/objectMaker/blog-backend/tools"
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
	password, err := tools.Crypto(userInfo.Password)
	if err != nil {
		fmt.Println(err.Error())
		tools.Res(c, gin.H{
			"message": "failed to crypto password",
		}, http.StatusInternalServerError)
		return
	}
	user := models.User{
		Username: userInfo.Username,
		Password: password,
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

	tools.Res(c, ResResult{
		Token: token,
		User:  user,
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

	tools.Res(c, users)
}

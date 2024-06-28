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
		tools.Res(c, "missing username", http.StatusBadRequest)
		return
	}
	//judge current user is exist or not
	var dbCurrentUser models.User
	db.DB.First(&dbCurrentUser, "username = ?", userInfo.Username)
	if dbCurrentUser.ID != 0 {
		tools.Res(c, "username already exists", http.StatusBadRequest)
		return
	}

	if len(userInfo.Username) <= 5 {
		tools.Res(c, "username too short", http.StatusBadRequest)
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
		tools.Res(c, fmt.Errorf("failed to create user: %v", result.Error), http.StatusInternalServerError)
		return
	}

	tools.Res(c, "success")
}
func SignIn(c *gin.Context) {
	type LogInfo struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var logInfo = LogInfo{}
	c.BindJSON(&logInfo)
	//logInfo
	var userInfo models.User
	db.DB.First(&userInfo, "username = ?", logInfo.Username)
	if userInfo.ID == 0 {
		tools.Res(c, "don't have the account", http.StatusUnauthorized)
		return
	}
	//have the account then validate
	err := tools.ComparePassword(logInfo.Password, userInfo.Password)

	if err != nil {
		fmt.Printf("error: %v\n", err)
		tools.Res(c, "password is not correct", http.StatusUnauthorized)
		return
	}
	token, err := jwt.New(userInfo.Username)
	if err != nil {
		log.Fatal("failed to create token: %w", err)
	}
	c.SetCookie("token", token, 3600, "/", "127.0.0.1", false, true)
	tools.Res(c, nil)
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

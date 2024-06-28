package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/objectMaker/blog-backend/routes"
)

var Router *gin.Engine

func New() {
	Router = gin.Default()
	Router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:3000", "http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	Router.POST("/createUser", routes.CreateUser)
	Router.GET("/getUserList", routes.GetUserList)
	Router.POST("/signIn", routes.SignIn)

	Router.Run("127.0.0.1:8080")
}

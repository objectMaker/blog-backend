package router

import (
	"github.com/gin-gonic/gin"
	"github.com/objectMaker/blog-backend/routes"
)

var Router *gin.Engine

func New() {
	Router = gin.Default()

	Router.GET("/ping", routes.Ping)

	Router.Run() // listen and serve on 0.0.0.0:8080
}

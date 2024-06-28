package tools

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Res(c *gin.Context, data interface{}, businessCode ...int) {
	var code int
	if len(businessCode) == 0 {
		code = http.StatusOK
	} else {
		code = businessCode[0]
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
	})
}

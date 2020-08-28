package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.GET("", Home)

	apiRG := r.Group("api")
	apiRG.Use(CheckToken)
	{
		apiRG.GET("", APIHome)
	}

	return r
}

// home
func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

// middleware

// 检查是否登陆
func CheckToken(c *gin.Context) {
	token := c.GetHeader("token")
	if token != "" {
		return
	}

	c.AbortWithStatusJSON(http.StatusExpectationFailed, gin.H{
		"message": "token not found",
	})
}

func APIHome(c *gin.Context) {
	Home(c)
}

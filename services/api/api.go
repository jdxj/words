package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jdxj/words/models"
	"github.com/jdxj/words/services"
)

func GetUserClaims(c *gin.Context) *models.UserClaims {
	bearerToken := c.GetHeader("Authorization")
	tokenStr := services.ExtractToken(bearerToken)

	uc := &models.UserClaims{}
	token, _ := jwt.ParseWithClaims(tokenStr, uc, services.KeyFunc)
	return token.Claims.(*models.UserClaims)
}

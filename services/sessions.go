package services

import (
	"net/http"
	"strings"

	"github.com/jdxj/words/config"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
	"github.com/jdxj/words/logger"
	"github.com/jdxj/words/models"
)

// PostSessions 表示创建/添加一个会话,
// 实际上是为客户端生成一个 token.
func PostSessions(c *gin.Context) {
	user := new(models.User)
	err := c.ShouldBind(user)
	if err != nil {
		logger.Error("gin.Context.ShouldBind: %s", err)
		resp := models.NewResponse(123, "invalid param", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	if err := user.Verify(); err != nil {
		logger.Error("User.Verify: %s", err)
		resp := models.NewResponse(123, "name or pass error", nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	uc := models.NewUserClaims(user.ID, user.Name)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	ss, err := token.SignedString(config.GetSecret())
	if err != nil {
		logger.Error("Token.SignedString: %s", err)
		resp := models.NewResponse(123, "invalid param", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := models.NewResponse(0, "ok", ss)
	c.JSON(http.StatusOK, resp)
}

func CheckToken(c *gin.Context) {
	bearerToken := c.GetHeader("Authorization")
	if bearerToken == "" {
		resp := models.NewResponse(123, "token not found", nil)
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}
	tokenStr := extractToken(bearerToken)

	uc := &models.UserClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, uc, keyFunc)
	if err != nil {
		logger.Error("jwt.ParseWithClaims: %s", err)
		resp := models.NewResponse(123, "invalid token", nil)
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	uc = token.Claims.(*models.UserClaims)
	logger.Info("sign in, id: %d, name: %s", uc.ID, uc.Name)
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	return config.GetSecret(), nil
}

func extractToken(tok string) string {
	if len(tok) > 6 && strings.ToUpper(tok[0:7]) == "BEARER " {
		return tok[7:]
	}
	return tok
}

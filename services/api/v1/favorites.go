package v1

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jdxj/words/logger"
	"github.com/jdxj/words/models"
)

func GetFavorites(c *gin.Context) {
	userIDStr := c.Param("userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		logger.Error("strconv.Atoi: %s", err)

		resp := &models.Response{
			Code:      123,
			Timestamp: time.Now().Unix(),
			Message:   "invalid param",
			Data:      nil,
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	fs := &models.Favorites{
		UserID: userID,
		Words:  nil,
	}
	words, err := fs.GetFavorites()
	if err != nil {
		logger.Error("favorites.GetFavorites: %s", err)

		resp := &models.Response{
			Code:      123,
			Timestamp: time.Now().Unix(),
			Message:   "无法查询",
			Data:      nil,
		}
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := &models.Response{
		Code:      0,
		Timestamp: time.Now().Unix(),
		Message:   "ok",
		Data:      words,
	}
	c.JSON(http.StatusOK, resp)
}

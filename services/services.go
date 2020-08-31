package services

import (
	"fmt"
	"net/http"

	"github.com/jdxj/words/config"
	v1 "github.com/jdxj/words/services/api/v1"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	mode := config.GetMode()
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.GET("", Home)
	r.POST("sessions", PostSessions)

	apiRG := r.Group("api")
	apiRG.Use(CheckToken)

	v1RG := apiRG.Group("v1")
	{
		// favorites
		v1RG.GET("favorites/:userID", v1.GetFavorites)
		v1RG.POST("favorites/:word")
		v1RG.DELETE("favorites/:word")

		// words
		wordsRG := v1RG.Group("words")
		wordsRG.Use(v1.CheckWordMW)
		{
			wordsRG.GET(":word", v1.Search)
			wordsRG.GET(":word/voice", v1.Voice)
		}
	}

	return r
}

// home
func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func NewServer() *http.Server {
	bind := fmt.Sprintf(":%s", config.GetPort())
	router := NewRouter()
	srv := &http.Server{
		Addr:    bind,
		Handler: router,
	}
	return srv
}

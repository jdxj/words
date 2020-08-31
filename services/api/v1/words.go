package v1

import (
	"database/sql"
	"net/http"

	"github.com/jdxj/words/logger"
	"github.com/jdxj/words/models"
	"github.com/jdxj/words/models/words"

	"github.com/gin-gonic/gin"
)

const (
	voiceType = "audio/mpeg"
)

//
func Search(c *gin.Context) {
	wordField := c.Param("word")
	word := &words.Word{
		Word: wordField,
	}
	// 1. 先向数据库中查询, 如果存在就返回
	err := word.Query()
	if err == nil {
		resp := models.NewResponse(0, "ok", word)
		c.JSON(http.StatusOK, resp)
		return
	}

	// 2. 出现其它未知错误, 通知客户端
	if err != sql.ErrNoRows {
		logger.Error("word.Query: %s", err)
		resp := models.NewResponse(123, "invalid param", nil)
		c.JSON(http.StatusInternalServerError, resp)
	}
}

func Voice(c *gin.Context) {
	wordField := c.Param("word")
	word := &words.Word{
		Word: wordField,
	}

	// 1. 查询 voice 是否在数据库中
	voice, err := word.QueryVoice()
	if err == nil {
		logger.Debug("hit voice on db: %s", word.Word)
		c.Data(http.StatusOK, voiceType, voice)
		return
	}

	// 2. 不在的话使用翻译器查询
	pool := models.TranslatorPool
	translator := pool.Get().(models.Translator)
	defer pool.Put(translator)

	voice, err = translator.Pronounce(wordField)
	if err != nil {
		logger.Error("translator.Pronounce: %s", err)
		resp := models.NewResponse(123, "voice not found", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	// 保存到数据库中
	if _, err := word.SaveVoice(voice); err != nil {
		logger.Error("word.SaveVoice: %s", err)
	}
	// 返回给客户端
	c.Data(http.StatusOK, voiceType, voice)
}

// CheckWord 是个中间件, 提前检查 word 有效性.
// 目的是降低之后路由处理逻辑的复杂性.
func CheckWordMW(c *gin.Context) {
	wordField := c.Param("word")
	word := &words.Word{Word: wordField}
	err := word.Query()
	// 1. 数据库中存在该 word
	if err == nil {
		return
	}

	// 2. 出现其它未知错误
	if err != sql.ErrNoRows {
		logger.Error("word.Query: %s", err)
		resp := models.NewResponse(123, "invalid param", nil)
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}

	// 3. 数据库中没有该 word
	pool := models.TranslatorPool
	translator := pool.Get().(models.Translator)
	defer pool.Put(translator)
	// 查询当前 word
	word, err = translator.Translate(wordField)
	if err != nil {
		logger.Error("translator.Translate: %s", err)
		resp := models.NewResponse(123, "invalid param", nil)
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}

	if !word.IsValid() {
		resp := models.NewResponse(123, "there may be a spelling error", word)
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	if _, err := word.Insert(); err != nil {
		logger.Error("word.Insert: %s", err)
	}
}

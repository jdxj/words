package v1

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jdxj/words/logger"
	"github.com/jdxj/words/models"
	"github.com/jdxj/words/models/words"
)

const (
	voiceType = "audio/mpeg"
)

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
		return
	}

	// 3. 数据库中没有存储该单词, 使用翻译器
	pool := models.TranslatorPool
	translator := pool.Get().(models.Translator)
	defer pool.Put(translator)

	word, err = translator.Translate(wordField)
	if err != nil {
		logger.Error("translator.Translate: %s", err)
		resp := models.NewResponse(123, "invalid param", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	if _, err := word.Insert(); err != nil {
		logger.Error("word.Insert: %s", err)
	}

	resp := models.NewResponse(0, "ok", word)
	c.JSON(http.StatusOK, resp)
}

func Voice(c *gin.Context) {
	wordField := c.Param("word")
	word := &words.Word{
		Word: wordField,
	}

	pool := models.TranslatorPool
	translator := pool.Get().(models.Translator)
	defer pool.Put(translator)

	// 1. 查询数据库中是否存在该 word,
	// 如果存在就查询 voice 是否存在
	// err 要么是 nil, 要么是 sql.ErrNoRows
	err := word.Query()
	if err == nil {
		voice, err := word.QueryVoice()
		// 获取 voice 成功, 所以不能返回 json
		if err == nil {
			c.Data(http.StatusOK, voiceType, voice)
			return
		}
		// 从翻译器获取 voice
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
		return
	}

	// 2. 出现其它未知错误, 通知客户端
	if err != sql.ErrNoRows {
		logger.Error("word.Query: %s", err)
		resp := models.NewResponse(123, "invalid param", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	// 3. 数据库中不存在该 word,
	// 那么先查询 word, 然后查询 voice
	word, err = translator.Translate(wordField)
	if err != nil {
		logger.Error("translator.Translate: %s", err)
		resp := models.NewResponse(123, "invalid param", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	_, err = word.Insert()
	if err != nil {
		logger.Error("word.Insert: %s", err)
		resp := models.NewResponse(123, "invalid param", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	// 从翻译器获取 voice
	voice, err := translator.Pronounce(wordField)
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

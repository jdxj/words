package models

import (
	"sync"

	"github.com/jdxj/words/config"
	"github.com/jdxj/words/models/google"
	"github.com/jdxj/words/models/words"
)

const (
	GoogleTranslator = "google"
)

var TranslatorPool = &sync.Pool{
	New: func() interface{} {
		t := config.GetTranslator()
		switch t {
		case GoogleTranslator:
			return google.NewGoogleTranslate()
		}

		// 如果输入了无法识别的翻译器,
		// 则使用谷歌翻译.
		return google.NewGoogleTranslate()
	},
}

type Translator interface {
	Translate(string) (*words.Word, error)
	Pronounce(string) ([]byte, error)
}

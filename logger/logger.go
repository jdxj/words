package logger

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/jdxj/words/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	fileName = "words.log"
)

var (
	sugar *zap.SugaredLogger

	ErrInvalidMode = errors.New("invalid mode")
)

func init() {
	ec := newEncoderConfig()
	encoder := zapcore.NewJSONEncoder(ec)

	writer := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    1,
		MaxAge:     3,
		MaxBackups: 10,
	}
	ws := zapcore.AddSync(writer)

	core := zapcore.NewCore(encoder, ws, level())
	sugar = zap.New(core).Sugar()

}

func newEncoderConfig() zapcore.EncoderConfig {
	mode := config.GetMode()
	switch mode {
	case gin.DebugMode:
		return zap.NewDevelopmentEncoderConfig()
	case gin.ReleaseMode:
		return zap.NewProductionEncoderConfig()
	}

	panic(ErrInvalidMode)
}

func level() zapcore.Level {
	mode := config.GetMode()
	switch mode {
	case gin.ReleaseMode:
		return zap.InfoLevel
	}
	return zap.DebugLevel
}

func Debug(template string, args ...interface{}) {
	sugar.Debugf(template, args...)
}

func Info(template string, args ...interface{}) {
	sugar.Infof(template, args...)
}

func Warn(template string, args ...interface{}) {
	sugar.Warnf(template, args...)
}

func Error(template string, args ...interface{}) {
	sugar.Errorf(template, args...)
}

func Sync() error {
	return sugar.Sync()
}

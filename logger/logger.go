package logger

import (
	"errors"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jdxj/words/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	sugar *zap.SugaredLogger

	ErrInvalidMode = errors.New("invalid mode")
)

func init() {
	ec := newEncoderConfig()
	encoder := zapcore.NewJSONEncoder(ec)
	ws := writeSyncer()

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

func writeSyncer() zapcore.WriteSyncer {
	mode := config.GetMode()
	switch mode {
	case gin.DebugMode:
		return zapcore.AddSync(os.Stdout)

	case gin.ReleaseMode:
		writer := &lumberjack.Logger{
			Filename:   config.GetLogPath(),
			MaxSize:    1,
			MaxAge:     3,
			MaxBackups: 10,
		}
		return zapcore.AddSync(writer)
	}

	panic(ErrInvalidMode)
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

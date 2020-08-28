package logger

import (
	"testing"
)

func TestInfo(t *testing.T) {
	defer Sync()

	Debug("test debug: %s", "34")
	Info("test info: %s", "ab")
	Warn("test warn: %s", "12")
	Error("test error: %s", "cd")
}

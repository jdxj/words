package config

import (
	"fmt"
	"testing"
)

func TestGetPort(t *testing.T) {
	port := GetPort()
	fmt.Printf("%s\n", port)
}

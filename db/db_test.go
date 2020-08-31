package db

import (
	"fmt"
	"testing"

	"github.com/jdxj/words/config"
)

func TestClose(t *testing.T) {
	if err := database.Ping(); err != nil {
		t.Fatalf("%s\n", err)
	}
	database.Close()
}

func TestOpenMySQL(t *testing.T) {
	kind := config.GetDatabase()
	fmt.Printf("%s\n", kind)

	if err := database.Ping(); err != nil {
		t.Fatalf("%s\n", err)
	}
	database.Close()
}

package db

import (
	"testing"
)

func TestClose(t *testing.T) {
	if err := sqlite3.Ping(); err != nil {
		t.Fatalf("%s\n", err)
	}
	sqlite3.Close()
}

func TestCreateWordsTable(t *testing.T) {
	defer sqlite3.Close()

	if err := CreateWordsTable(); err != nil {
		t.Fatalf("%s\n", err)
	}
}

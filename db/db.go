package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

const (
	// todo: 调试用, 路径应该是配置的
	dsn = "/home/jdxj/workspace/words/db/words.db"

	// Table Name
	WordsTN     = "words"
	UsersTN     = "users"
	FavoritesTN = "favorites"
)

var (
	sqlite3 *sql.DB
)

func init() {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		panic(err)
	}
	sqlite3 = db

	_, err = EnableForeignKey()
	if err != nil {
		db.Close()
		panic(err)
	}
}

func Close() error {
	return sqlite3.Close()
}

// Query
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	return sqlite3.Query(query, args...)
}

func QueryRow(query string, args ...interface{}) *sql.Row {
	return sqlite3.QueryRow(query, args...)
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	return sqlite3.Exec(query, args...)
}

func EnableForeignKey() (sql.Result, error) {
	return setForeignKeyStat(true)
}

func DisableForeignKey() (sql.Result, error) {
	return setForeignKeyStat(false)
}

func setForeignKeyStat(v bool) (sql.Result, error) {
	query := fmt.Sprintf("PRAGMA foreign_keys = %t", v)
	return Exec(query)
}

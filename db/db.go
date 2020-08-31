package db

import (
	"database/sql"
	"errors"

	"github.com/jdxj/words/config"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

const (
	// db kind
	mysqlDB  = "mysql"
	sqliteDB = "sqlite"

	// Table Name
	WordsTN     = "words"
	UsersTN     = "users"
	FavoritesTN = "favorites"
)

var (
	ErrInvalidDatabase = errors.New("invalid database")

	database *sql.DB
)

func init() {
	var err error

	kind := config.GetDatabase()
	switch kind {
	case mysqlDB:
		database, err = openMySQL()
	case sqliteDB:
		database, err = openSQLite()
	default:
		panic(ErrInvalidDatabase)
	}

	if err != nil {
		panic(err)
	}
}

func Close() error {
	return database.Close()
}

// Query
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	return database.Query(query, args...)
}

func QueryRow(query string, args ...interface{}) *sql.Row {
	return database.QueryRow(query, args...)
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	return database.Exec(query, args...)
}

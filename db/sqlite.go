package db

import (
	"database/sql"
	"fmt"

	"github.com/jdxj/words/config"
)

func openSQLite() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", config.GetSQLite())
	if err != nil {
		return nil, err
	}

	err = enableForeignKey()
	if err != nil {
		_ = db.Close()
		return nil, err
	}
	return db, nil
}

func enableForeignKey() error {
	_, err := setForeignKeyStat(true)
	return err
}

func disableForeignKey() error {
	_, err := setForeignKeyStat(false)
	return err
}

func setForeignKeyStat(v bool) (sql.Result, error) {
	query := fmt.Sprintf("PRAGMA foreign_keys = %t", v)
	return Exec(query)
}

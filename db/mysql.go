package db

import (
	"database/sql"
	"fmt"

	"github.com/jdxj/words/config"
)

const (
	dsnFormat = "%s:%s@tcp(%s)/%s?loc=Local&parseTime=true"
)

func openMySQL() (*sql.DB, error) {
	mysql := config.GetMySQL()
	dsn := fmt.Sprintf(dsnFormat, mysql.User, mysql.Pass, mysql.Addr, mysql.Base)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, db.Ping()
}

package models

import (
	"database/sql"
	"fmt"

	"github.com/jdxj/words/db"
)

type User struct {
	ID       int
	Name     string `form:"name"`
	Password string `form:"pass"`
}

func (u *User) Insert() (sql.Result, error) {
	query := fmt.Sprintf(`insert into %s (name,password) values (?,?)`, db.UsersTN)
	return db.Exec(query, u.Name, u.Password)
}

// Verify 用于登陆场景
func (u *User) Verify() error {
	query := fmt.Sprintf(`select id from %s where name=? and password=?`, db.UsersTN)
	row := db.QueryRow(query, u.Name, u.Password)
	return row.Scan(&u.ID)
}

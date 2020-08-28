package models

import (
	"database/sql"
	"fmt"

	"github.com/jdxj/words/db"
)

type User struct {
	ID       int
	Name     string
	Password string
}

func (u *User) Insert() (sql.Result, error) {
	query := fmt.Sprintf(`insert into %s (name,password) values (?,?)`, db.UsersTN)
	return db.Exec(query, u.Name, u.Password)
}

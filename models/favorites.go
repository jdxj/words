package models

import (
	"database/sql"
	"fmt"

	"github.com/jdxj/words/db"
)

type Favorite struct {
	ID     int
	UserID int
	WordID int
}

func (f *Favorite) Insert() (sql.Result, error) {
	query := fmt.Sprintf(`insert into %s (user_id,word_id) values (?,?)`, db.FavoritesTN)
	return db.Exec(query, f.UserID, f.WordID)
}

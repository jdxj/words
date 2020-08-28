package models

import (
	"database/sql"
	"fmt"

	"github.com/jdxj/words/db"
)

type Word struct {
	ID       int    `json:"id"`
	Word     string `json:"word"`
	Phonetic string `json:"phonetic"`
	Meaning  string `json:"meaning"`
	Sound    []byte `json:"sound"`
}

func (w *Word) Insert() (sql.Result, error) {
	query := fmt.Sprintf(`insert into %s (word,phonetic,meaning,sound) values (?,?,?,?)`, db.WordsTN)
	return db.Exec(query, w.Word, w.Phonetic, w.Meaning, w.Sound)
}

func (w *Word) Query() error {
	query := fmt.Sprintf(`select id,phonetic,meaning from %s where word=?`, db.WordsTN)
	row := db.QueryRow(query, w.Word)
	return row.Scan(&w.ID, &w.Phonetic, &w.Meaning)
}

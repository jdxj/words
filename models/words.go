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
	Voice    []byte `json:"voice"`
}

func (w *Word) Insert() (sql.Result, error) {
	query := fmt.Sprintf(`insert into %s (word,phonetic,meaning) values (?,?,?)`, db.WordsTN)
	return db.Exec(query, w.Word, w.Phonetic, w.Meaning, w.Voice)
}

func (w *Word) Query() error {
	query := fmt.Sprintf(`select id,phonetic,meaning from %s where word=?`, db.WordsTN)
	row := db.QueryRow(query, w.Word)
	return row.Scan(&w.ID, &w.Phonetic, &w.Meaning)
}

func (w *Word) QueryVoice() ([]byte, error) {
	query := fmt.Sprintf(`select voice from %s where id=?`, db.WordsTN)
	var data []byte
	row := db.QueryRow(query, w.ID)
	return data, row.Scan(&data)
}

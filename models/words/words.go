package words

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
	Voice    []byte `json:"voice,omitempty"`

	Wrong string `json:"wrong,omitempty"`
}

// IsValid 检查当前 word 对象所持有的数据是否是有效的.
// 一个有效的 word 数据的条件是指定的字段不为空.
func (w *Word) IsValid() bool {
	if w.Word == "" || w.Phonetic == "" || w.Meaning == "" {
		return false
	}
	return true
}

func (w *Word) Insert() (sql.Result, error) {
	query := fmt.Sprintf(`insert into %s (word,phonetic,meaning) values (?,?,?)`, db.WordsTN)
	return db.Exec(query, w.Word, w.Phonetic, w.Meaning)
}

// Query 查询指定 word, 并将其它数据绑定到自身.
func (w *Word) Query() error {
	query := fmt.Sprintf(`select id,phonetic,meaning from %s where word=?`, db.WordsTN)
	row := db.QueryRow(query, w.Word)
	return row.Scan(&w.ID, &w.Phonetic, &w.Meaning)
}

func (w *Word) QueryVoice() ([]byte, error) {
	query := fmt.Sprintf(`select voice from %s where word=? and voice is not null`, db.WordsTN)
	var data []byte
	row := db.QueryRow(query, w.Word)
	return data, row.Scan(&data)
}

func (w *Word) SaveVoice(voice []byte) (sql.Result, error) {
	query := fmt.Sprintf(`update %s set voice=? where word=?`, db.WordsTN)
	return db.Exec(query, voice, w.Word)
}

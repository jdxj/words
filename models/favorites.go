package models

import (
	"database/sql"
	"fmt"

	"github.com/jdxj/words/models/words"

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

type Favorites struct {
	UserID int
	Words  []*words.Word
}

func (fs *Favorites) GetFavorites() ([]*words.Word, error) {
	if len(fs.Words) != 0 {
		return fs.Words, nil
	}

	query := fmt.Sprintf(`select w.id, w.word, w.phonetic, w.meaning
from %s as f left outer join %s as w on f.word_id=w.id
where f.user_id=?`, db.FavoritesTN, db.WordsTN)

	rows, err := db.Query(query, fs.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var myWords []*words.Word
	for rows.Next() {
		w := &words.Word{}
		err := rows.Scan(&w.ID, &w.Word, &w.Phonetic, &w.Meaning)
		if err != nil {
			return nil, err
		}
		myWords = append(myWords, w)
	}

	fs.Words = myWords
	return fs.Words, nil
}

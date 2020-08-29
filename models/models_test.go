package models

import (
	"fmt"
	"testing"

	"github.com/jdxj/words/db"
)

func TestWord_Insert(t *testing.T) {
	defer db.Close()

	w := &Word{
		ID:       0,
		Word:     "abc",
		Phonetic: "abc",
		Meaning:  "ABC",
		Voice:    nil,
	}
	if _, err := w.Insert(); err != nil {
		t.Fatalf("%s\n", err)
	}
}

func TestWord_Query(t *testing.T) {
	defer db.Close()

	w := &Word{Word: "abc"}
	if err := w.Query(); err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%#v\n", *w)
}

func TestWord_QueryVoice(t *testing.T) {
	defer db.Close()

	w := &Word{ID: 1}
	data, err := w.QueryVoice()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%t\n", data == nil)
	fmt.Printf("%x\n", data)
	fmt.Printf("%s\n", data)
}

func TestUser_Insert(t *testing.T) {
	defer db.Close()

	u := &User{
		ID:       0,
		Name:     "jdxj",
		Password: "jdxj",
	}
	if _, err := u.Insert(); err != nil {
		t.Fatalf("%s\n", err)
	}
}

func TestFavorite_Insert(t *testing.T) {
	defer db.Close()

	f := &Favorite{
		ID:     0,
		UserID: 1,
		WordID: 1,
	}
	if _, err := f.Insert(); err != nil {
		t.Fatalf("%s\n", err)
	}
}

func TestFavorites_GetFavorites(t *testing.T) {
	defer db.Close()

	fs := &Favorites{
		UserID: 1,
		Words:  nil,
	}

	words, err := fs.GetFavorites()
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	for _, w := range words {
		fmt.Printf("%#v\n", *w)
	}
}

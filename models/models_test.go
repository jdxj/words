package models

import (
	"fmt"
	"testing"

	"github.com/jdxj/words/db"
)

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

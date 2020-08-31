package words

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

	// 查询不存在的单词
	w = &Word{Word: "kjh"}
	if err := w.Query(); err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%#v\n", *w)
}

func TestWord_QueryVoice(t *testing.T) {
	defer db.Close()

	w := &Word{Word: "apple"}
	data, err := w.QueryVoice()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%t\n", data == nil)
	fmt.Printf("%x\n", data)
	fmt.Printf("%s\n", data)
}

func TestWord_SaveVoice(t *testing.T) {
	defer db.Close()

	w := &Word{
		Word:     "56l",
		Phonetic: "abc",
		Meaning:  "abc",
	}
	if _, err := w.Insert(); err != nil {
		t.Fatalf("%s\n", err)
	}

	if _, err := w.SaveVoice([]byte("mock voice")); err != nil {
		t.Fatalf("%s\n", err)
	}
}

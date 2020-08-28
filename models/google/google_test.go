package google

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGoogleTranslate_Translate(t *testing.T) {
	gt := NewGoogleTranslate()
	tr, err := gt.Translate("apple")
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	fmt.Printf("sentences:\n")
	for _, t := range tr.Sentences {
		fmt.Printf("%#v\n", *t)
	}

	fmt.Printf("dict:\n")
	for _, a := range tr.Dict {
		fmt.Printf("%#v\n", *a)
	}

	fmt.Printf("src: %s\n", tr.Src)
	fmt.Printf("confidence: %f\n", tr.Confidence)
	fmt.Printf("spell: %s\n", tr.Spell)
	fmt.Printf("ldresult: %v\n", *tr.LdResult)

	fmt.Printf("definitions:\n")
	for _, d := range tr.Definitions {
		fmt.Printf("%#v\n", *d)
	}
}

func TestGoogleTranslate_Pronounce(t *testing.T) {
	gt := NewGoogleTranslate()
	voice, err := gt.Pronounce("apple")
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	//file, err := os.Create("voice")
	//if err != nil {
	//	t.Fatalf("%s\n", err)
	//}
	//defer file.Close()
	//defer file.Sync()
	//
	//_, err = file.Write(voice)
	//if err != nil {
	//	t.Fatalf("%s\n", err)
	//}

	typ := http.DetectContentType(voice)
	fmt.Printf("type: %s\n", typ)
	fmt.Printf("%x\n", voice)
}

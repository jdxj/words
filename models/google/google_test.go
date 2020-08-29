package google

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGoogleTranslate_Translate(t *testing.T) {
	gt := NewGoogleTranslate()
	// 查询空字符串会怎样. 经测试返回空数据
	word, err := gt.Translate("")
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%#v\n", *word)
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

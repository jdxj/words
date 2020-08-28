package google

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"

	"github.com/jdxj/words/models"
)

func newTranslateResponse(response *http.Response) (*TranslateResponse, error) {
	var isGzip bool
	ce := response.Header.Get("Content-Encoding")
	if ce == "gzip" {
		isGzip = true
	}

	body := response.Body
	defer body.Close()

	reader := body.(io.Reader)
	if isGzip {
		r, err := gzip.NewReader(reader)
		if err != nil {
			return nil, err
		}
		reader = r
	}

	tr := &TranslateResponse{}
	decoder := json.NewDecoder(reader)
	return tr, decoder.Decode(tr)
}

type TranslateResponse struct {
	Sentences   []*Translation  `json:"sentences"`
	Dict        []*Annotation   `json:"dict"`
	Src         string          `json:"src"`
	Confidence  float64         `json:"confidence"`
	Spell       json.RawMessage `json:"spell"`
	LdResult    *LdResult       `json:"ld_result"`
	Definitions []*Definition   `json:"definitions"`
}

func (tr *TranslateResponse) ToWord() *models.Word {
	w := &models.Word{}
	if len(tr.Sentences) < 2 {
		return w
	}
	t1 := tr.Sentences[0]
	t2 := tr.Sentences[2]

	w.Word = t1.Orig
	w.Phonetic = t2.TranslIt
	w.Meaning = t1.Trans
	return w
}

type Translation struct {
	// 1
	Trans   string `json:"trans"`   // 中文
	Orig    string `json:"orig"`    // 英文
	Backend int    `json:"backend"` // 未知

	// 2
	TranslIt    string `json:"translit"`     // 拼音
	SrcTranslIt string `json:"src_translit"` // 音标
}

type Annotation struct {
	Pos      string     `json:"pos"`
	Terms    []string   `json:"terms"`
	Entry    []*Reverse `json:"entry"`
	BaseForm string     `json:"base_form"`
	PosEnum  int        `json:"pos_enum"`
}

// Reverse 反查
type Reverse struct {
	Word               string   `json:"word"`
	ReverseTranslation []string `json:"reverse_translation"`
	Score              float64  `json:"score"`
}

type LdResult struct {
	SrcLangs            []string  `json:"srclangs"`
	SrcLangsConfidences []float64 `json:"srclangs_confidences"`
	ExtendedSrcLangs    []string  `json:"extended_srclangs"`
}

type Definition struct {
	Pos      string     `json:"pos"`
	Entry    []*Example `json:"entry"`
	BaseForm string     `json:"base_form"`
}

type Example struct {
	Gloss        string `json:"gloss"`
	DefinitionID string `json:"definition_id"`
}

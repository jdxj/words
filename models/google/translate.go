package google

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"time"
)

const (
	// todo: 这个值需要测试
	timeout = 10 * time.Second

	host           = "translate.google.cn"
	acceptEncoding = "gzip"
	userAgent      = "GoogleTranslate/6.11.0.06.325960053 (Linux; U; Android 10; Redmi K20 Pro)"

	// 重要的是 q 参数
	translateAPI = "https://translate.google.cn/translate_a/single?dj=1&q=%s&sl=en&tl=zh-CN&hl=zh-CN&ie=UTF-8&oe=UTF-8&client=at&dt=t&dt=ld&dt=qca&dt=rm&dt=bd&dt=md&dt=ss&dt=ex&dt=sos&otf=3"
	// 重要的参数有: q, textlen
	ttsAPI = "https://translate.google.cn/translate_tts?ie=utf-8&client=at&q=%s&tl=en&total=1&idx=0&textlen=%d&prev=input"
)

func NewGoogleTranslate() *GoogleTranslate {
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar:     jar,
		Timeout: timeout,
	}

	gt := &GoogleTranslate{
		c: client,
	}
	return gt
}

type GoogleTranslate struct {
	c *http.Client
}

// Translate 翻译指定单词
func (gt *GoogleTranslate) Translate(q string) (*TranslateResponse, error) {
	req := newTranslateRequest(q)
	resp, err := gt.c.Do(req)
	if err != nil {
		return nil, err
	}
	return newTranslateResponse(resp)
}

// Pronounce 获取指定单词发音
func (gt *GoogleTranslate) Pronounce(q string) ([]byte, error) {
	req := newTTSRequest(q)
	resp, err := gt.c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func newTTSRequest(q string) *http.Request {
	URL := fmt.Sprintf(ttsAPI, q, len(q))
	return newRequest(URL)
}

func newTranslateRequest(q string) *http.Request {
	URL := fmt.Sprintf(translateAPI, q)
	return newRequest(URL)
}

func newRequest(URL string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, URL, nil)
	h := req.Header

	h.Set("Accept-Encoding", acceptEncoding)
	h.Set("User-Agent", userAgent)
	h.Set("Host", host)
	return req
}

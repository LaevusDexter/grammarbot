package grammarbot

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"sync"
	"unsafe"
)

const MaxLength = 50000

func New(apiKey string) *GrammarBot {
	if apiKey == "" {
		apiKey = "xyz"
	}

	return &GrammarBot{
		Language: "en-US",
		ApiKey:   apiKey,
		BaseURI:  "http://api.grammarbot.io",
		Version:  "v2",
		ApiName:  "check",
		Client: http.DefaultClient,
	}
}

func Validate(text string) bool {
	if ValidateErr(text) != nil {
		return true
	}

	return false
}

var ErrExceedLimits = errors.New("Text longer than 50,000 characters is not allowed.")
var ErrEmpty = errors.New("Text is empty.")

func ValidateErr(text string) error {
	if len(text) > MaxLength {
		return ErrExceedLimits
	}

	if len(text) == 0 {
		return ErrEmpty
	}

	return nil
}

var bufferPool = &sync.Pool{}

func acquireBuffer(size int) []byte {
	result, ok := bufferPool.Get().([]byte)

	if !ok || cap(result) < size {
		return make([]byte, size)
	}

	result = result[:size]

	return result
}

// Check - Check a given piece of text for grammatical errors.
func (api GrammarBot) Check(text string) (*Response, error) {

	/*buf := make([]byte,
		len(text) +
		len(api.BaseURI) +
		len(api.ApiName) +
		len(api.Version) +
		len(api.ApiKey) +
		len(api.Language) +
		len("//api_key=&text=&language=")) */

	if err := ValidateErr(text); err != nil {
		return nil, err
	}

	text = url.QueryEscape(text)

	buf := acquireBuffer(
		len(text) +
		len(api.BaseURI) +
		len(api.ApiName) +
		len(api.Version) +
		len(api.ApiKey) +
		len(api.Language) +
		len("//api_key=&text=&language="))

	p := copy(buf[:], api.BaseURI)
	p += copy(buf[p:], "/")
	p += copy(buf[p:], api.Version)
	p += copy(buf[p:], "/")
	p += copy(buf[p:], api.ApiName)

	endpoint := p

	p += copy(buf[p:], "api_key=")
	p += copy(buf[p:], api.ApiKey)
	p += copy(buf[p:], "&text=")
	p += copy(buf[p:], text)
	p += copy(buf[p:], "&language=")
	copy(buf[p:], api.Language)

	req, err := http.NewRequest("POST", b2s(buf[:endpoint]), nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = b2s(buf[endpoint:])

	resp, err := api.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("GrammarBot.Check->response: " + resp.Status)
	}

	var result Response

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	bufferPool.Put(buf)

	return &result, nil
}

// CheckBytes - Check a given piece of text for grammatical errors.
func (api GrammarBot) CheckBytes(text []byte) (*Response, error) {
	return api.Check(b2s(text))
}

var defaultGrammarBot = *New("")

// Check - Check a given piece of text for grammatical errors.
func Check(text string) (*Response, error) {
	return defaultGrammarBot.Check(text)
}

// CheckBytes - Check a given piece of text for grammatical errors.
func CheckBytes(text []byte) (*Response, error) {
	return defaultGrammarBot.CheckBytes(text)
}

func b2s(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

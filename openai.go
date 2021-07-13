package openai

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

var OPENAI_API_KEY string

// Auth sets OPENAI authentication key
func Auth(key string) {
	OPENAI_API_KEY = key
}

// OPENAI API documentation:
// curl https://api.openai.com/v1/engines/davinci/completions \
//   -H "Content-Type: application/json" \
//   -H "Authorization: Bearer $OPENAI_API_KEY" \
//   -d '{
//   "prompt": "",
//   "temperature": 0,
//   "max_tokens": 60,
//   "top_p": 1.0,
//   "frequency_penalty": 0.0,
//   "presence_penalty": 0.0,
//   "stop": ["\"\"\""]
// }'

// Prompt makes an HTTP request to the OpenAI api and returns the response JSON as a string
// it takes the following arguments: prompt, temperature
func Prompt(prompt string, temperature float64) (string, error) {

	if OPENAI_API_KEY == "" {
		return "", errors.New("OPENAI_API_KEY not set")
	}
	if prompt == "" {
		return "", errors.New("Empty prompt... nothing to do")
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/engines/davinci/completions", nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+OPENAI_API_KEY)
	req.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(`{"prompt": ` + prompt + `, "temperature": ` + strconv.FormatFloat(temperature, 'f', -1, 64) + `, "max_tokens": 60, "top_p": 1.0, "frequency_penalty": 0.0, "presence_penalty": 0.0, "stop": ["\"\"\""]}`)))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), err
}

package engine

import (
	"bufio"
	"io"
	"net/http"
	"strings"
)

type Engine struct {
	url      string
	keywords []string
}

func NewEngine(url string, keywords []string) *Engine {
	return &Engine{url: url, keywords: keywords}
}

func (e Engine) CollectResults() ([]string, error) {
	body, err := e.getBody()
	if err != nil {
		return nil, err
	}

	body = e.colorKeywords(body)

	return e.getMatchingLines(body), nil
}

func (e Engine) getBody() (string, error) {
	response, err := http.Get(e.url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (e Engine) colorKeywords(str string) string {
	for _, keyword := range e.keywords {
		str = strings.ReplaceAll(str, keyword, green+keyword+reset)
	}
	return str
}

func (e Engine) getMatchingLines(str string) []string {
	var result []string

	// Scanner para percorrer cada linha da string
	scanner := bufio.NewScanner(strings.NewReader(str))

	for scanner.Scan() {
		line := scanner.Text()
		if e.containsKeyword(line) {
			result = append(result, strings.TrimSpace(line))
		}
	}

	return result
}

func (e Engine) containsKeyword(str string) bool {
	for _, keyword := range e.keywords {
		if strings.Contains(str, keyword) {
			return true
		}
	}
	return false
}

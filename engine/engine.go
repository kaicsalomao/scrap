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

func trimSpaces(str string) string {
	// Remover espacos em branco do início da string
	for len(str) > 0 && str[0] == ' ' {
		str = str[1:]
	}
	// Remover espacos em branco do final da string
	for len(str) > 0 && str[len(str)-1] == ' ' {
		str = str[:len(str)-1]
	}
	return str
}

func (e Engine) CollectResults() ([]string, error) {

	response, err := http.Get(e.url)
	if err != nil {
		return []string{}, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return []string{}, err
	}

	// Colorizar todas as keywords que estao no corpo
	for _, keyword := range e.keywords {
		if strings.Contains(string(body), keyword) {
			body = []byte(strings.ReplaceAll(string(body), keyword, green+keyword+reset))
		}
	}

	var result []string

	// Converter o body para um scanner para facilitar a leitura
	scanner := bufio.NewScanner(strings.NewReader(string(body)))

	// Percorrer cada linha do body
	for scanner.Scan() {

		line := scanner.Text()

		for _, keyword := range e.keywords {
			if strings.Contains(line, keyword) {
				result = append(result, trimSpaces(line))
				break
			}
		}
	}
	return result, nil
}

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

/* Cria um ponteiro para a Engine */
func NewEngine(url string, keywords []string) *Engine {
	return &Engine{url: url, keywords: keywords}
}

// Remove os espacos do inicio e do final de uma frase "  oi usuario! "-> "oi usuario"
func trimSpaces(str string) string {
	// Remove espacos em branco do início da string
	for len(str) > 0 && str[0] == ' ' {
		str = str[1:]
	}
	// Remove espacos em branco do final da string
	for len(str) > 0 && str[len(str)-1] == ' ' {
		str = str[:len(str)-1]
	}
	return str
}

/* Coleta os resultados */
func (e Engine) Collect() ([]string, error) {
	// Faz a requisicao http
	response, err := http.Get(e.url)
	if err != nil {
		return []string{}, err
	}
	defer response.Body.Close()

	// Le o corpo da resposta
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return []string{}, err
	}

	// Coloriza todas as keywords que estao no corpo
	for _, keyword := range e.keywords {
		if strings.Contains(string(body), keyword) {
			body = []byte(strings.ReplaceAll(string(body), keyword, Green+keyword+Reset))
		}
	}

	// Slice para armazenar os resultados
	var result []string

	// Converte o body para um scanner para facilitar a leitura
	scanner := bufio.NewScanner(strings.NewReader(string(body)))

	// Percorrer cada linha do body
	for scanner.Scan() {
		// Linha atual
		line := scanner.Text()

		// Verifica se um dos keywords esta na linha
		for _, keyword := range e.keywords {
			if strings.Contains(line, keyword) {
				// A respectiva linha vai para o resultado
				result = append(result, trimSpaces(line))

				// Para de verificar a linha atual
				break
			}
		}
	}
	return result, nil
}

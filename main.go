package main

import (
	"fmt"
	"os"
	"scrap/engine"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: scrap website palavra1 palavra2 palavra3 ...")
		os.Exit(0)
	}
	var url string
	var keywords []string

	// Coletar o nome do website
	url = os.Args[1]

	// Coletar as keywords
	for i := 2; i < len(os.Args); i++ {
		keywords = append(keywords, os.Args[i])
	}

	// Cria a engine
	eng := engine.NewEngine(url, keywords)
	result, err := eng.Collect()
	if err != nil {
		fmt.Println("Um erro aconteceu:", err)
		os.Exit(1)
	}

	// Mostrar os resultados
	for index, item := range result {
		fmt.Printf("%d: %s\n", index, item)
	}
}

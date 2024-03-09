package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"scrap/engine"
)

const HELP_MSG string = "Uso: scrap <url> <keywords>"

func collectArgs() ([]string, error) {
	help := flag.Bool("help", false, "exibe a mensagem de ajuda")

	flag.Usage = func() {
		fmt.Println(HELP_MSG)
		flag.PrintDefaults()
	}

	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	args := flag.Args()

	if len(args) < 1 {
		return nil, errors.New("nao ha argumentos")
	}

	return args, nil
}

func main() {
	args, err := collectArgs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	url := args[0]
	eng := engine.NewEngine(url, args[1:])

	result, err := eng.CollectResults()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for index, item := range result {
		fmt.Printf("%d: %s\n", index, item)
	}
}

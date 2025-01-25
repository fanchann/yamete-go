package main

import (
	"fmt"

	yamete "github.com/fanchann/yamete-go"
)

func main() {
	yameteCfg := yamete.YameteConfig{
		URL: "https://raw.githubusercontent.com/fanchann/toxic-word-list/refs/heads/master/id_toxic_371.txt",
		// File: "files/id_words_toxic.txt",
	}
	yameteInit, err := yamete.NewYamete(&yameteCfg)
	if err != nil {
		panic(err)
	}

	response := yameteInit.AnalyzeText("dasar lu bot!")

	fmt.Printf("response.OriginalText: %v\n", response.OriginalText)
	fmt.Printf("response.CensoredText: %v\n", response.CensoredText)
	fmt.Printf("response.CensoredWords: %v\n", response.CensoredWords)
	fmt.Printf("response.CensoredCount: %v\n", response.CensoredCount)
}

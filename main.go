package main

import (
	"dictionaryProject/dict"
	"fmt"
	"log"
)

func main() {
	dictionary := dict.Dictionary{}

	word := "hello"
	definition := "Greeting"

	err := dictionary.Add(word, definition)

	if err != nil {
		log.Fatalln(err)
	}

	def, sErr := dictionary.Search(word)

	if sErr != nil {
		log.Fatalln(sErr)
	}

	fmt.Println(def)

	addErr := dictionary.Add(word, definition)

	if addErr != nil {
		log.Fatalln(addErr)
	}

}

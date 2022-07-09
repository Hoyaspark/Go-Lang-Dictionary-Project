package main

import (
	"dictionaryProject/dict"
	"fmt"
	"log"
)

func main() {
	dictionary := dict.Dictionary{}

	word := "hello"

	dictionary.Add(word, "First")
	err := dictionary.Update(word, "Second")
	dictionary.Delete(word)

	if err != nil {
		log.Fatalln(err)
	}
	def, sErr := dictionary.Search(word)

	if sErr != nil {
		log.Fatalln(sErr)
	}

	fmt.Println(def)
}

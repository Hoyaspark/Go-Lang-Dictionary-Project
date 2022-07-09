package main

import (
	"dictionaryProject/dict"
	"fmt"
	"log"
)

func main() {
	dictionary := dict.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("second")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(definition)

}

package main

import (
	"fmt"
	"log"

	"github.com/kr/pretty"
	"github.com/piprate/json-gold/ld"
)

func main() {
	url := "https://kgsearch.googleapis.com/v1/entities:search?query=juan+rulfo&key=<insert-api-ip>&limit=10&indent=True"

	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")

	// expanding remote document
	expanded, err := proc.Expand(url, options)
	if err != nil {
		log.Println("Error when expanding JSON-LD document:", err)
		return
	}

	//ld.PrintDocument("JSON-LD expansion succeeded", expanded)

	data := expanded[0].(map[string]interface{})["http://schema.org/itemListElement"]
	itemListElement := data.([]interface{})
	element := itemListElement[len(itemListElement)-1]
	results := element.(map[string]interface{})["http://schema.org/result"]
	value := results.([]interface{})[0].(map[string]interface{})["http://schema.org/description"]

	fmt.Printf("%# v ", pretty.Formatter(value))
}

package main

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/knakk/rdf"
	"github.com/piprate/json-gold/ld"
)

func main() {
	url := "https://kgsearch.googleapis.com/v1/entities:search?query=juan+rulfo&key=<insert-api-ip>&limit=1&indent=True"

	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	options.ProcessingMode = ld.JsonLd_1_1
	options.Format = "application/n-quads"
	options.Algorithm = "URDNA2015"

	// expanding remote document
	normalized, err := proc.Normalize(url, options)
	if err != nil {
		log.Println("Error when expanding JSON-LD document:", err)
		return
	}

	//print(normalized.(string))
	dec := rdf.NewTripleDecoder(strings.NewReader(normalized.(string)), rdf.Turtle)
	for triple, err := dec.Decode(); err != io.EOF; triple, err = dec.Decode() {
		fmt.Printf("Object\n: %s\n", triple.Obj)
		fmt.Printf("Predicate\n %s\n: ", triple.Pred)
		fmt.Printf("Subject\n: %s\n", triple.Subj)
	}
}

package main

import (
	"log"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/lercher/rdf/sparql"
)

func main() {
	input, err := antlr.NewFileStream(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	sparql.Parse(input)
	if err != nil {
		log.Fatalln(err)
	}
}

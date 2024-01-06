package main

import (
	"log"
	"os"

	"github.com/antlr4-go/antlr/v4"
	"github.com/lercher/rdf/sparql"
)

func main() {
	input, err := antlr.NewFileStream(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	ast, err := sparql.Parse(input)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("base", ast.Base)
}

package sparql

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/lercher/rdf/sparql/parser"
)

type walker struct {
	*parser.BaseSparqlListener
}

func (w *walker) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

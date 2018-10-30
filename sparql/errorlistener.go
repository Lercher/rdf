package sparql

import (
	"log"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// ErrorListener records and logs syntax errors produced by the antlr4 parser
type ErrorListener struct {
	*antlr.DiagnosticErrorListener
	*errors
}


// SyntaxError implements an interface method of antlr.DiagnosticErrorListener
func (el ErrorListener) SyntaxError(
	recognizer antlr.Recognizer,
	offendingSymbol interface{},
	line, column int,
	msg string,
	e antlr.RecognitionException,
) {
	el.ErrorCount++
	log.Printf("%d:%d Syntax: %s\n", line, column, msg)
}

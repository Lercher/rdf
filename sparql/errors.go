package sparql

import (
	"log"

	"github.com/antlr4-go/antlr/v4"
)

// Baseerror logs and counts semantic errors in a file
type errors struct {
	ErrorCount   int
	WarningCount int
}

// Semantic Errors

// SemErr logs a semantic error
func (e *errors) SemErr(
	offendingToken antlr.Token,
	id int,
	msg string,
) {
	line := offendingToken.GetLine()
	column := offendingToken.GetColumn()
	e.SemanticError(line, column, id, msg)
}

// SemanticErrorAt logs a semantic error at a scanned Token
func (e *errors) SemanticErrorAt(t antlr.Token, id int, msg string) {
	if t == nil {
		e.SemanticErrorGlobal(id, msg)
	}
	e.SemanticError(t.GetLine(), t.GetColumn(), id, msg+" at "+t.GetText())
}

// SemanticError logs a semantic error
func (e *errors) SemanticError(
	line, column int,
	id int,
	msg string,
) {
	e.ErrorCount++
	log.Printf("%d:%d Error E%d: %s\n", line, column, id, msg)
}

// SemanticError logs a semantic error not bound to a location
func (e *errors) SemanticErrorGlobal(
	id int,
	msg string,
) {
	e.SemanticError(0, 0, id, msg)
}

// Semantic Warnings

// SemWarn logs a semantic warning
func (e *errors) SemWarn(
	offendingToken antlr.Token,
	id int,
	msg string,
) {
	line := offendingToken.GetLine()
	column := offendingToken.GetColumn()
	e.SemanticWarning(line, column, id, msg)
}

// SemanticWarning logs a semantic warning
func (e *errors) SemanticWarning(
	line, column int,
	id int,
	msg string,
) {
	e.WarningCount++
	log.Printf("%d:%d Warning W%d: %s\n", line, column, id, msg)
}

// SemanticWarningGlobal logs a semantic warning not bound to a location
func (e *errors) SemanticWarningGlobal(
	id int,
	msg string,
) {
	e.SemanticWarning(0, 0, id, msg)
}

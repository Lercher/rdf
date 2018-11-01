package sparql

import (
	"github.com/lercher/rdf/graph"
	"github.com/lercher/rdf/algebra"
	"github.com/lercher/rdf/sparql/parser"
)

// the suffix C of a variable denotes a parser context
// the suffix T of a variable denotes an antlr Token

func convertVerbContext(ctx *parser.VerbContext, symbols *symbols) (*algebra.Verb, error) {
	aC := ctx.GetA()
	if aC != nil {
		return algebra.VerbLiteral(A), nil
	}	
	viC := ctx.GetVi()
	vC := viC.GetVariable()
	if vC != nil {
		v := symbols.Variables.Register(vC.GetText())
		return algebra.VerbVariable(v), nil
	}
	iriC := viC.GetIri()
	literalT := iriC.GetLiteral()
	if literalT != nil {
		literal := graph.IRIParse(literalT.GetText())
		return algebra.VerbLiteral(literal), nil
	}
	prefixedC := iriC.GetPrefixed()
	literal, err := graph.IRIPrefixed(prefixedC.GetText(), symbols.PrefixedIRIs)
	if err != nil {
		return nil, err
	}
	return algebra.VerbLiteral(literal), err
}

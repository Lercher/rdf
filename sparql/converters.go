package sparql

import (
	"fmt"

	"github.com/lercher/rdf/algebra"
	"github.com/lercher/rdf/graph"
	"github.com/lercher/rdf/sparql/parser"
)

// the suffix C of a variable denotes a parser context
// the suffix T of a variable denotes an antlr Token

func convertIriRef(iriC parser.IIriRefContext, symbols *symbols) (graph.IRI, error) {
	literalT := iriC.GetLiteral()
	if literalT != nil {
		literal := graph.IRIParse(literalT.GetText())
		return literal, nil
	}
	prefixedC := iriC.GetPrefixed()
	literal, err := graph.IRIPrefixed(prefixedC.GetText(), symbols.PrefixedIRIs)
	if err != nil {
		return graph.NotAnIRI, err
	}
	return literal, nil
}

func convertLiteral(lC parser.IRdfLiteralContext, symbols *symbols) (*graph.Literal, error) {
	tx := lC.GetStr().GetText()
	lang := lC.GetLang()
	if lang != nil {
		lit := graph.LiteralFrom(tx, lang.GetText(), graph.NotAnIRI)
		return lit, nil
	}
	if lC.GetIri() != nil {
		dt, err := convertIriRef(lC.GetIri(), symbols)
		if err != nil {
			return nil, err
		}
		return graph.LiteralFrom(tx, "", dt), nil
	}
	return graph.LiteralFrom(tx, "", graph.NotAnIRI), nil
}

func convertVerbContext(ctx *parser.VerbContext, symbols *symbols) (algebra.Verb, error) {
	aC := ctx.GetA()
	if aC != nil {
		return A, nil
	}
	viC := ctx.GetVi()
	vC := viC.GetVariable()
	if vC != nil {
		v := symbols.Variables.Register(vC.GetText())
		return v, nil
	}
	iriC := viC.GetIri()
	iri, err := convertIriRef(iriC, symbols)
	if err != nil {
		return nil, err
	}
	return iri, nil
}

func convertVarOrTermContext(ctx *parser.VarOrTermContext, symbols *symbols) (algebra.Verb, error) {
	vC := ctx.GetVariable()
	if vC != nil {
		v := symbols.Variables.Register(vC.GetText())
		return v, nil
	}

	// graphTerm : iriRef|rdfLiteral|numericLiteral|booleanLiteral|blankNode|NIL;
	gtC := ctx.GetGt()

	// iriRef
	if gtC.GetIri() != nil {
		iri, err := convertIriRef(gtC.GetIri(), symbols)
		if err != nil {
			return nil, err
		}
		return iri, nil
	}

	// rdfLiteral
	if gtC.GetLit() != nil {
		lit, err := convertLiteral(gtC.GetLit(), symbols)
		if err != nil {
			return nil, err
		}
		return lit, nil
	}

	return nil, fmt.Errorf("not implemnted: %v", gtC.GetText())
}

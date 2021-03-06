package sparql

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lercher/rdf/algebra"
	"github.com/lercher/rdf/sparql/parser"
	"github.com/lercher/rdf/values"
)

// the suffix C of a variable denotes a parser context
// the suffix T of a variable denotes an antlr Token

func convertIriRef(iriC parser.IIriRefContext, symbols *symbols) (values.IRI, error) {
	literalT := iriC.GetLiteral()
	if literalT != nil {
		literal := values.IRIParse(literalT.GetText())
		return literal, nil
	}
	prefixedC := iriC.GetPrefixed()
	literal, err := values.IRIPrefixed(prefixedC.GetText(), symbols.PrefixedIRIs)
	if err != nil {
		return values.NotAnIRI, err
	}
	return literal, nil
}

func convertLiteral(lC parser.IRdfLiteralContext, symbols *symbols) (*values.Literal, error) {
	tx := lC.GetStr().GetText()
	lang := lC.GetLang()
	if lang != nil {
		lit := values.LiteralFrom(tx, lang.GetText(), values.NotAnIRI)
		return lit, nil
	}
	if lC.GetIri() != nil {
		dt, err := convertIriRef(lC.GetIri(), symbols)
		if err != nil {
			return nil, err
		}
		return values.LiteralFrom(tx, "", dt), nil
	}
	return values.LiteralFrom(tx, "", values.NotAnIRI), nil
}

func convertVerbContext(ctx parser.IVerbContext, symbols *symbols) (algebra.Term, error) {
	aC := ctx.GetA()
	if aC != nil {
		return values.A, nil
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

func convertVarOrIRIRefContext(ctx parser.IVarOrIRIrefContext, symbols *symbols) (algebra.Term, error) {
	vC := ctx.GetVariable()
	if vC != nil {
		v := symbols.Variables.Register(vC.GetText())
		return v, nil
	}

	return convertIriRef(ctx.GetIri(), symbols)
}

func convertVarOrTermContext(ctx parser.IVarOrTermContext, symbols *symbols) (algebra.Term, error) {
	vC := ctx.GetVariable()
	if vC != nil {
		v := symbols.Variables.Register(vC.GetText())
		return v, nil
	}

	// graphTerm : iriRef|rdfLiteral|numericLiteral|booleanLiteral|blankNode|NIL;
	gtC := ctx.GetGt()

	// iriRef
	if gtC.GetIri() != nil {
		return convertIriRef(gtC.GetIri(), symbols)
	}

	// rdfLiteral
	if gtC.GetLit() != nil {
		return convertLiteral(gtC.GetLit(), symbols)
	}

	// numericLiteral
	if gtC.GetNum() != nil {
		num := gtC.GetNum().GetText()
		if strings.Contains(num, ".") || strings.Contains(num, "e") || strings.Contains(num, "E") {
			f, err := strconv.ParseFloat(num, 64)
			if err != nil {
				return nil, err
			}
			return values.Float(f), nil
		}
		i, err := strconv.Atoi(num)
		if err != nil {
			return nil, err
		}
		return values.Int(i), nil
	}

	// booleanLiteral
	if gtC.GetBol() != nil {
		bol := gtC.GetBol().GetText()
		bol = strings.ToLower(bol)
		if bol == "true" {
			return values.Bool(true), nil
		}
		return values.Bool(false), nil
	}

	// blankNode
	if gtC.GetBlk() != nil {
		return algebra.BlankVariable, nil
	}

	// NIL
	if gtC.GetEmp() != nil {
		return values.NIL, nil
	}

	return nil, fmt.Errorf("implementation: unknown graphTerm variant %v", gtC.GetText())
}

func convertGroupGraphPatternContext(ctx parser.IGroupGraphPatternContext, symbols *symbols, mode string) (*algebra.Tree, error) {
	tree := &algebra.Tree{Mode: mode}
	for _, tbC := range ctx.GetTb() {
		// triplesBlock
		blocks, err := convertTriplesBlockContext(tbC, symbols)
		if err != nil {
			return nil, err
		}
		for _, bl := range blocks {
			child := &algebra.Tree{
				Mode: "BGP",
				Term: bl,
			}
			tree.Children = append(tree.Children, child)
		}
	}
	for _, gpntC := range ctx.GetGpnt() {
		// graphPatternNotTriples
		child, err := convertGraphPatternNotTriplesContext(gpntC, symbols)
		if err != nil {
			return nil, err
		}
		tree.Children = append(tree.Children, child)
	}
	for _, fltC := range ctx.GetFlt() {
		// Filter
		return nil, fmt.Errorf("filter not yet supported: %v", fltC.GetStart()) // #TODO
	}
	return tree, nil
}

func convertTriplesBlockContext(ctx parser.ITriplesBlockContext, symbols *symbols) ([]*algebra.Block, error) {
	var list []*algebra.Block
	for ctx != nil {
		tssC := ctx.GetTss()
		blocks, err := convertTriplesSameSubjectContext(tssC, symbols)
		if err != nil {
			return nil, err
		}
		list = append(list, blocks...)

		ctx = ctx.GetTb() // tail recursion of triplesBlock
	}
	return list, nil
}

func convertTriplesSameSubjectContext(ctx parser.ITriplesSameSubjectContext, symbols *symbols) ([]*algebra.Block, error) {
	var list []*algebra.Block
	subjectC := ctx.GetSubject()
	if subjectC != nil {
		subject, err := convertVarOrTermContext(subjectC, symbols)
		if err != nil {
			return nil, err
		}
		propertiesC := ctx.GetProperties()
		verbs := propertiesC.GetVerbs()
		for i, v := range verbs {
			predicate, err := convertVerbContext(v, symbols)
			if err != nil {
				return nil, err
			}
			// https://www.w3.org/TR/rdf-sparql-query/#objLists
			objectlistC := propertiesC.GetOl()[i]
			for _, objC := range objectlistC.GetOb() {
				graphNodeC := objC.GetGn()
				vtC := graphNodeC.GetVt()
				if vtC != nil {
					object, err := convertVarOrTermContext(vtC, symbols)
					if err != nil {
						return nil, err
					}
					block := &algebra.Block{
						Subject:   subject,
						Predicate: predicate,
						Object:    object,
					}
					list = append(list, block)
				} else {
					tnC := graphNodeC.GetTn()                                                             // triplesNode
					return nil, fmt.Errorf("triplesNode as object not yet supported: %v", tnC.GetStart()) // #TODO
				}
			}
		}
	}
	return list, nil
}

func convertGraphPatternNotTriplesContext(ctx parser.IGraphPatternNotTriplesContext, symbols *symbols) (*algebra.Tree, error) {
	// optionalGraphPattern
	if ctx.GetOgp() != nil {
		return convertGroupGraphPatternContext(ctx.GetOgp().GetGgp(), symbols, "OPTIONAL")
	}

	// groupOrUnionGraphPattern
	if ctx.GetGup() != nil {
		tree := &algebra.Tree{Mode: "UNION"}
		for _, u := range ctx.GetGup().GetUnion() {
			child, err := convertGroupGraphPatternContext(u, symbols, "JOIN")
			if err != nil {
				return nil, err
			}
			tree.Children = append(tree.Children, child)
		}
		return tree, nil
	}

	// graphGraphPattern
	if ctx.GetGgp() != nil {
		tree, err := convertGroupGraphPatternContext(ctx.GetGgp().GetGgp(), symbols, "GRAPH")
		if err != nil {
			return nil, err
		}
		tree.Term, err = convertVarOrIRIRefContext(ctx.GetGgp().GetVi(), symbols)
		return tree, err
	}

	return nil, fmt.Errorf("implementation: unknown graphPatternNotTriples %v", ctx.GetStart())
}

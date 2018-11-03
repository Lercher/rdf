package processor

import (
	"fmt"

	"github.com/lercher/rdf/algebra"
	"github.com/lercher/rdf/graph"
)

func join(ctx *context, tree *algebra.PatternTree, current algebra.Binding) error {
	recursiveJoin(ctx, tree.Children, current) // #TODO: children need to be part of join
	return nil
}

func recursiveJoin(ctx *context, children []*algebra.PatternTree, current algebra.Binding) (bool, error) {
	if len(children) == 0 {
		return ctx.receiver(ctx.g, ctx.variables, current)
	}

	child0 := children[0]
	switch child0.Mode { // #TODO: need to replace this switch by an interface
	default:
		return false, fmt.Errorf("%q not allowed", child0.Mode)
	case "BGP":
		block := child0.Term.(*algebra.Block)
		sp := patternFromTerm(ctx.g, current, block.Subject)
		pp := patternFromTerm(ctx.g, current, block.Predicate)
		op := patternFromTerm(ctx.g, current, block.Object)
		tp := &algebra.TriplePattern{S: sp, P: pp, O: op}
		ms := Match(ctx.g, tp)
		for _, t := range ms {
			currentBound := current.BindTriple(tp, t)
			cont, err := recursiveJoin(ctx, children[1:], currentBound)
			if !cont || err != nil {
				return cont, err
			}
		}
	case "OPTIONAL":
		return false, fmt.Errorf("%q not yet supported", child0.Mode)
	case "GRAPH":
		return false, fmt.Errorf("%q not yet supported", child0.Mode)
	case "FILTER":
		return false, fmt.Errorf("%q not yet supported", child0.Mode)
	case "UNION":
		return false, fmt.Errorf("%q not yet supported", child0.Mode)
	}
	return true, nil // #TODO: is continue=true OK here?
}

func patternFromTerm(g *graph.Graph, current algebra.Binding, t algebra.Term) *algebra.Pattern {
	bt, virtual := current.BindTerm(t)
	if bt != nil {
		return algebra.PatternFromTerm(g, bt)
	}
	return algebra.PatternLiteral(virtual)
}

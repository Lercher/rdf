package processor

import (
	"fmt"

	"github.com/lercher/rdf/algebra"
	"github.com/lercher/rdf/graph"
)

func join(ctx *context, tree *algebra.Tree, current algebra.Binding) error {
	// "JOIN"

	// only non empty joins have results, so we need to check, if we have children
	if len(tree.Children) != 0 {		
		recursiveJoin(ctx, tree.Children, current)
	} 
	return nil
}

func recursiveJoin(ctx *context, children []*algebra.Tree, current algebra.Binding) (bool, error) {
	if len(children) == 0 {
		return ctx.receiver(current, ctx.variables)
	}

	my := children[0]
	switch my.Mode { // #TODO: need to replace this switch by an interface
	default:
		return false, fmt.Errorf("%q not allowed", my.Mode)
	case "BGP":
		// A BGP (basic graph pattern) has to pass each result of a Match to my siblings.
		block := my.Term.(*algebra.Block)
		sp := patternFromTerm(ctx.g, current, block.Subject)
		pp := patternFromTerm(ctx.g, current, block.Predicate)
		op := patternFromTerm(ctx.g, current, block.Object)
		tp := &algebra.TriplePattern{S: sp, P: pp, O: op}
		ms := Match(ctx.g, tp)
		for _, t := range ms {
			nextBinding := current.BindTriple(tp, t)
			cont, err := recursiveJoin(ctx, children[1:], nextBinding) // do next sibling with my new bindings
			if !cont || err != nil {
				return cont, err
			}
		}
	case "OPTIONAL":
		// An OPTIONAL has to give current to its children, that are BGP/OPTIONAL/UNION/FILTER/GRAPH
		// to join them all and pass any result bindings (bs) of the join to my children.
		// If no match succeeded at all, hand over the current to my siblings,
		// as if the OPTIONAL wasn't there at all
		myChildrenHadResults := false
		newctx := ctx.WithReceiver(func(bs algebra.Binding, vs *algebra.Variables) (bool, error) {
			// intercept results of the OPTIONAL's tree, to hand them over to my next BGP/OPTIONAL/whatever
			myChildrenHadResults = true
			return recursiveJoin(ctx, children[1:], bs) // present the completed optional bindings (bs) to my next sibling
		})
		err := join(newctx, my, current) // go deeper with the new receiver and my current bindings
		if err != nil {
			return false, err
		}
		if !myChildrenHadResults {
			cont, err := recursiveJoin(ctx, children[1:], current) // just ignore this optional node and do the next sibling with my original bindings
			if !cont || err != nil {
				return cont, err
			}
		} // else is nothing more to do, because the newctx.receiver propagated all results to the next sibling
	case "UNION":
		// A UNION has to give current to all of its children, that are JOINs
		// and pass each of the resulting bindings (bs) to my siblings
		newctx := ctx.WithReceiver(func(bs algebra.Binding, vs *algebra.Variables) (bool, error) {
			// intercept results of the inner JOIN tree, to hand them over to my next BGP/OPTIONAL/whatever
			return recursiveJoin(ctx, children[1:], bs) // present the completed optional bindings (bs) to my next sibling
		})
		for _, innerJoinTree := range my.Children {
			err := join(newctx, innerJoinTree, current)
			if err != nil {
				return false, err
			}
		}
	case "FILTER":
		return false, fmt.Errorf("%q not yet supported", my.Mode)
	case "GRAPH":
		return false, fmt.Errorf("%q not yet supported", my.Mode)
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

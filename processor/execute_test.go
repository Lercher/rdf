package processor_test

import (
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/lercher/rdf/algebra"
	"github.com/lercher/rdf/graph"
	"github.com/lercher/rdf/processor"
	"github.com/lercher/rdf/sparql"
)

func parse(t *testing.T, q string) *sparql.AST {
	t.Helper()

	input := antlr.NewInputStream(q)
	ast, err := sparql.Parse(input)
	if err != nil {
		t.Fatal(err)
	}
	return ast
}

func TestExecute2Proj(t *testing.T) {
	exec(t, `select $sub $obj where {?sub ?pred ?obj}`)
}

func TestExecute1Proj(t *testing.T) {
	exec(t, `select $obj where {?sub ?pred ?obj}`)
}

func TestExecuteAll(t *testing.T) {
	exec(t, `select * where {?sub ?pred ?obj}`)
}

func TestExecute1Join(t *testing.T) {
	exec(t, `select ?sub ?sub2 ?pred 
	where {
		?sub ?pred ?obj .
		?sub2 ?pred ?obj .
	}`)
}

func TestExecuteSome(t *testing.T) {
	exec(t, `select * where {?sub ?pred "justus"}`)
}

func exec(t *testing.T, sparql string) {
	t.Helper()
	g, _, _ , _ := gr2()
	ast := parse(t, sparql)
	a := ast.Algebra()
	a = a.Optimize()
	err := processor.Execute(a, g, func(g *graph.Graph, vs *algebra.Variables, bs algebra.Binding) (bool, error) {
		m := bs.Materialize(g, vs)
		t.Log(m)
		return true, nil
	})
	if err != nil {
		t.Error(err)
	}
	t.Error("#TODO")
}

func gr2() (*graph.Graph, *graph.Triple, *graph.Triple, *graph.Triple) {
	g := graph.New()
	t0 := g.Assert("martin", "telefon", "+49897482400")
	t1 := g.Assert("andreas", "telefon", "+49897482400")
	t2 := g.Assert("martin", "boss", "justus")
	return g, t0, t1, t2
}

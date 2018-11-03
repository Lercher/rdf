package processor_test

import (
	"fmt"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/lercher/rdf/algebra"
	"github.com/lercher/rdf/graph"
	"github.com/lercher/rdf/processor"
	"github.com/lercher/rdf/sparql"
)

func TestExecuteDblProj(t *testing.T) {
	exec(t, `select $sub $sub where {?sub ?pred ?obj}`)
	t.Error("#TODO")
}

func TestExecute2Proj(t *testing.T) {
	exec(t, `select $sub $obj where {?sub ?pred ?obj}`)
	t.Error("#TODO")
}

func TestExecute1Proj(t *testing.T) {
	ms := exec(t, `select $obj where {?sub ?pred ?obj} order by $obj`)
	want := `[obj=string:+49897482400] [obj=string:+49897482400] [obj=string:justus] [obj=string:martin]]`
	got := fmt.Sprint(ms)
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestExecuteAll(t *testing.T) {
	ms := exec(t, `select * where {?sub ?pred ?obj}`)
	if len(ms) != 4 {
		t.Errorf("want %d results, got %d", 4, len(ms))
	}
}

func TestExecuteSubEqObj(t *testing.T) {
	ms := exec(t, `select * where {$sub ?pred ?sub}`)
	if len(ms) != 1 {
		t.Errorf("want %d result, got %d", 1, len(ms))
	}
	want := `[[sub=string:martin pred=string:name]]`
	got := fmt.Sprint(ms)
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestExecute1Join(t *testing.T) {
	exec(t, `select ?sub ?sub2 ?pred
	where {
		?sub ?pred ?obj .
		?sub2 ?pred ?obj .
	}`)
	t.Error("#TODO")
}

func TestExecuteSome(t *testing.T) {
	exec(t, `select * where {?sub ?pred "justus"}`)
	t.Error("#TODO")
}

func exec(t *testing.T, sparql string) (list [][]*algebra.Materialized) {
	t.Helper()
	g, _, _, _ := gr2()
	ast := parse(t, sparql)
	a := ast.Algebra()
	a = a.Optimize()
	err := processor.Execute(a, g, func(bs algebra.Binding, vs *algebra.Variables) (bool, error) {
		m := bs.Materialize(g, vs)
		list = append(list, m)
		t.Log(m)
		return true, nil
	})
	if err != nil {
		t.Error(err)
	}
	return list
}

func gr2() (*graph.Graph, *graph.Triple, *graph.Triple, *graph.Triple) {
	g := graph.New()
	t0 := g.Assert("martin", "telefon", "+49897482400")
	t1 := g.Assert("andreas", "telefon", "+49897482400")
	t2 := g.Assert("martin", "boss", "justus")
	g.Assert("martin", "name", "martin")
	return g, t0, t1, t2
}

func parse(t *testing.T, q string) *sparql.AST {
	t.Helper()

	input := antlr.NewInputStream(q)
	ast, err := sparql.Parse(input)
	if err != nil {
		t.Fatal(err)
	}
	return ast
}

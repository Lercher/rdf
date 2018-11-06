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

func TestExecuteEmpty(t *testing.T) {
	ms := exec(t, `select * {}`)
	want(t, ms, 0, 0)
}

func TestExecuteEmpty1(t *testing.T) {
	ms := exec(t, `select * {{}}`)
	want(t, ms, 0, 0)
}

func TestExecuteEmpty11(t *testing.T) {
	ms := exec(t, `select * {{}{}}`)
	want(t, ms, 0, 0)
}

func TestExecuteEmpty2(t *testing.T) {
	ms := exec(t, `select * {{{}}}`)
	want(t, ms, 0, 0)
}

func TestExecuteAll(t *testing.T) {
	ms := exec(t, `select * where {?sub ?pred ?obj}`)
	want(t, ms, 4, 3)
}

func TestExecuteDblProj(t *testing.T) {
	exec1(t, `select $sub $sub where {?sub <boss> ?obj}`, `[sub="martin"]`)
}

func TestExecute2Proj(t *testing.T) {
	ms := exec(t, `select $sub $obj where {?sub ?pred ?obj}`)
	want(t, ms, 4, 2)
}

func TestExecuteOrderBy(t *testing.T) {
	ms := exec(t, `select $obj where {?sub ?pred ?obj} order by $obj`)
	want := `[[obj="+49897482400"] [obj="+49897482400"] [obj="justus"] [obj="martin"]]`
	got := fmt.Sprint(ms)
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestExecuteSubEqObj(t *testing.T) {
	exec1(t, `select * where {$sub ?pred ?sub}`, `[sub="martin" pred=<name>]`)
}

func TestExecute1Join(t *testing.T) {
	ms := exec(t, `select ?sub ?sub2 ?pred
	where {
		?sub ?pred ?obj .
		?sub2 ?pred ?obj .
	}`)
	want(t, ms, 6, 3)
}

func TestExecuteSome(t *testing.T) {
	exec1(t, `select * where {?sub ?pred "justus"}`, `[sub="martin" pred=<boss>]`)
}

func TestExecuteOptionalGood(t *testing.T) {
	exec1(t, `select * where {"andreas" ?pred ?andreas.           "martin" ?pred $martin} `, `[pred=<telefon> andreas="+49897482400" martin="+49897482400"]`)
	exec1(t, `select * where {"andreas" ?pred ?andreas. optional {"martin" ?pred $martin}}`, `[pred=<telefon> andreas="+49897482400" martin="+49897482400"]`)
}

func TestExecuteOptionalEmpty(t *testing.T) {
	exec1(t, `     select * where {"andreas" ?pred ?andreas.}                                `, `[pred=<telefon> andreas="+49897482400"]`)
	ms := exec(t, `select * where {"andreas" ?pred ?andreas.           $sub <boss> ?andreas}`)
	want(t, ms, 0, 0)
	exec1(t, `     select * where {"andreas" ?pred ?andreas. optional {$sub <boss> ?andreas}}`, `[pred=<telefon> andreas="+49897482400" sub=nil]`)
}

func TestExecuteOptional2Empty(t *testing.T) {
	exec1(t, `select *
	{"andreas" ?pred ?andreas. 
	optional {$sub1 <boss> ?andreas. 
         	  $sub2 <telefon> "martin"}
	}`, `[pred=<telefon> andreas="+49897482400" sub1=nil sub2=nil]`)
}

func TestExecute2OptionalEmpty(t *testing.T) {
	exec1(t, `select *
	{"andreas" ?pred ?andreas. 
	optional {$sub1 <boss> ?andreas.} 
	optional {$sub2 <telefon> "martin"}
	}`, `[pred=<telefon> andreas="+49897482400" sub1=nil sub2=nil]`)
}

func TestExecute2GoodOptional(t *testing.T) {
	exec1(t, `select *
	{        "andreas" <telefon> ?andreas_tel. 
	optional {"martin" <telefon> ?martin_tel}
	optional {"martin" <boss>    ?martin_boss} 
	}`, `[andreas_tel="+49897482400" martin_tel="+49897482400" martin_boss="justus"]`)
}

func TestExecuteMixed2Optional(t *testing.T) {
	exec1(t, `select *
	{        "andreas" <telefon> ?andreas_tel. 
	optional {"martin" <telefon> ?martin_tel}
	optional {?who     <boss>    ?andreas_tel} 
	}`, `[andreas_tel="+49897482400" martin_tel="+49897482400" who=nil]`) // who=nil b/c nobody has andreas' telefon number as her boss
}

func TestExecuteMixed1Optional(t *testing.T) {
	exec1(t, `select *
	{        "andreas" <telefon> ?andreas_tel. 
	optional {"martin" <telefon> ?martin_tel.
	          "martin" <boss>    ?andreas_tel} 
	}`, `[andreas_tel="+49897482400" martin_tel=nil]`) // martin_tel=nil b/c martin's boss is not andreas' telefon number
}

func TestExecuteBasicUnion(t *testing.T) {
	ms := exec(t, `select ?name ?tel ?boss {
		{"andreas" <telefon> ?tel}
		union
		{"martin"  <telefon> ?tel}
		union
		{"martin" <name> ?name ; <telefon> ?tel ; <boss> ?boss}
		union
		{?s <boss> ?boss ; <telefon> ?tel; <name> ?name}
	}`)
	want(t, ms, 4, 3)
}

func TestDistinct(t *testing.T) {
	ms := exec(t, `select distinct ?s {?s ?p ?o}`)
	want(t, ms, 2, 1) // reduced is 3
	ms = exec(t, `select distinct ?p {?s ?p ?o}`)
	want(t, ms, 3, 1)
	ms = exec(t, `select distinct ?o {?s ?p ?o}`)
	want(t, ms, 3, 1)
	ms = exec(t, `select distinct ?s ?p {?s ?p ?o}`)
	want(t, ms, 4, 2)
}

func TestReduced(t *testing.T) {
	ms := exec(t, `select Reduced ?s {?s ?p ?o}`)
	want(t, ms, 3, 1) // distinct is 2
	ms = exec(t, `select Reduced ?p {?s ?p ?o}`)
	want(t, ms, 3, 1)
	ms = exec(t, `select Reduced ?o {?s ?p ?o}`)
	want(t, ms, 3, 1)
	ms = exec(t, `select Reduced ?s ?p {?s ?p ?o}`)
	want(t, ms, 4, 2)
}
//------------------------ Helpers -------------------------------

func want(t *testing.T, ms [][]*algebra.Materialized, lines, cols int) {
	t.Helper()

	if len(ms) != lines {
		t.Errorf("want %d results, got %d", lines, len(ms))
	}
	if len(ms) > 0 && len(ms[0]) != cols {
		t.Errorf("want %d columns, got %d", cols, len(ms[0]))
	}
}

func exec1(t *testing.T, sparql, want string) {
	t.Helper()

	ms := exec(t, sparql)
	if len(ms) != 1 {
		t.Errorf("want 1 result line, got %d", len(ms))
	}
	if len(ms) > 0 {
		got := fmt.Sprint(ms[0])
		if got != want {
			t.Errorf("want %s, got %s", want, got)
		}
	}
}

func exec(t *testing.T, sparql string) (list [][]*algebra.Materialized) {
	t.Helper()
	g, _, _, _ := gr2()
	ast := parse(t, sparql)
	a := ast.Algebra()
	a = a.Optimize()
	t.Log(a.Tree)
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
	t0 := g.AssertLiterally("martin", "telefon", "+49897482400")
	t1 := g.AssertLiterally("andreas", "telefon", "+49897482400")
	t2 := g.AssertLiterally("martin", "boss", "justus")
	g.AssertLiterally("martin", "name", "martin")
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

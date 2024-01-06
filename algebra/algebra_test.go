package algebra_test // we need the sparql parser for testing

import (
	"strings"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/lercher/rdf/algebra"
	"github.com/lercher/rdf/sparql"
)

func TestSimpleAlgebra1(t *testing.T) {
	at(t,
		`select * where {$s ?p ?o}`, `
(PROJECTION (Projection:*)
  (JOIN
    (BGP (Block:{(Variable:$0) (Variable:$1) (Variable:$2)})
    )
  )
)
`)
}

func TestSimpleAlgebra2(t *testing.T) {
	at(t,
		`select $s {$s ?p ?o . $o $p $s.}`, `
(PROJECTION (Projection:[$0])
  (JOIN
    (BGP (Block:{(Variable:$0) (Variable:$1) (Variable:$2)})
    )
    (BGP (Block:{(Variable:$2) (Variable:$1) (Variable:$0)})
    )
  )
)
`)
}

func at(t *testing.T, s, want string) {
	t.Helper()
	a := parse(t, s)
	want = strings.TrimSpace(want)
	got := strings.TrimSpace(a.Tree.String())
	if got != want {
		t.Errorf("%s\nwant:\n%s\ngot:\n%s\n", s, want, got)
	}
}

func parse(t *testing.T, q string) *algebra.Algebra {
	t.Helper()

	input := antlr.NewInputStream(q)
	ast, err := sparql.Parse(input)
	if err != nil {
		t.Fatal(err)
	}
	a := ast.Algebra()
	a = a.Optimize()
	return a
}

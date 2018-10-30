package sparql

import (
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

const (
	emptyQuery = `   `

	simpleQuery = `
BASE <http://xyz>
PREFIX sch-ont:   <http://education.data.gov.uk/def/school/>
PREFIX ont:   <http://ontology/>
SELECT ?name WHERE {
  ?school a sch-ont:School.
  ?school sch-ont:establishmentName ?name.
  ?school sch-ont:districtAdministrative <http://statistics.data.gov.uk/id/local-authority-district/00AA>.
}
ORDER BY ?name
`

	badQuery = `
BASE <http://xyz>
PREFIX sch-ont:   <http://education.data.gov.uk/def/school/>
SELECT ?name WHERE { ?s ?p
`
)

func parse(t *testing.T, q string) (*AST, error) {
	t.Helper()

	input := antlr.NewInputStream(q)
	return Parse(input)
}

func TestSparqlParserEmpty(t *testing.T) {
	_, err := parse(t, emptyQuery)
	if err == nil {
		t.Error("empty should err")
	}
}

func TestSparqlParserBad(t *testing.T) {
	_, err := parse(t, badQuery)
	if err == nil {
		t.Error("bad should err, but it doesn't")
	}
}

func TestSparqlParserSimple(t *testing.T) {
	ast, err := parse(t, simpleQuery)
	if err != nil {
		t.Error("simple should not err, but", err)
	}

	w(t, "base", ast.Base, `<http://xyz>`)
	if len(ast.PrefixedIRIs) != 2 {
		t.Errorf("length prefix: want %d, got %d", 2, len(ast.PrefixedIRIs))
	} else {
		w(t, "prefix0", ast.PrefixedIRIs[0], `sch-ont:<http://education.data.gov.uk/def/school/>`)
		w(t, "prefix1", ast.PrefixedIRIs[1], `ont:<http://ontology/>`)
	}
	w(t, "type", str(ast.QueryType), "select")
}

func w(t *testing.T, what string, ser stringer, want string) {
	t.Helper()
	got := ser.String()
	if got != want {
		t.Errorf("%s: want %q, got %q", what, want, got)
	}
}

type stringer interface{ String() string }

type str string

func (s str) String() string { return string(s) }

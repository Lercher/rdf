package sparql

import (
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

const (
	emptyQuery = `   `

	simpleQuery = `
base <http://xyz>
PREFIX sch-ont:   <http://education.data.gov.uk/def/school/>
prefix ont:   <http://ontology/ÄÖÜßäöü/>
SELECT distinct ?name Where {
  ?school a sch-ont:School.
  ?school sch-ont:establishmentName ?name.
  ?school sch-ont:districtAdministrative <http://statistics.data.gov.uk/id/local-authority-district/00AA>.
}
ORDER by ?name
`

	badQuery = `
BASE <http://xyz>
PREFIX sch-ont:   <http://education.data.gov.uk/def/school/>
SELECT ?name FROM <http://unsupported> WHERE { ?s ?p
`

	unsupportedQuery = `SELECT ?name FROM <http://unsupported> WHERE { ?s ?s ?s.}`
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
		t.Fatal("bad should err, but it doesn't")
	}
	got := err.Error()
	ws(t, "bad error message", got, "1 syntax error(s)")
}

func TestSparqlParserUnsupported(t *testing.T) {
	_, err := parse(t, unsupportedQuery)
	if err == nil {
		t.Fatal("unsupported should err, but it doesn't")
	}
	got := err.Error()
	ws(t, "unsupported error message", got, "1 semantic error(s) and 0 warning(s)")
}

func TestSparqlParserSimple(t *testing.T) {
	t.Log(simpleQuery)
	ast, err := parse(t, simpleQuery)
	if err != nil {
		t.Fatal("simple should not err, but", err)
	}

	w(t, "base", ast.Base, `<http://xyz>`)
	if len(ast.PrefixedIRIs) != 2 {
		t.Errorf("length prefix: want %d, got %d", 2, len(ast.PrefixedIRIs))
	} else {
		w(t, "prefix0", ast.PrefixedIRIs[0], `sch-ont:<http://education.data.gov.uk/def/school/>`)
		w(t, "prefix1", ast.PrefixedIRIs[1], `ont:<http://ontology/ÄÖÜßäöü/>`)
	}
	ws(t, "type", ast.QueryType, "select")
	ws(t, "modifier", ast.Modifier, "distinct")
	ws(t, "projection", ast.Projection.String(ast.Variables), "[name]")
	w(t, "variables", ast.Variables, "map[name:0 school:1]")
}

func ws(t *testing.T, what string, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("%s: want %q, got %q", what, want, got)
	}
}

func w(t *testing.T, what string, ser stringer, want string) {
	t.Helper()
	ws(t, what, ser.String(), want)
}

type stringer interface{ String() string }

type str string

func (s str) String() string { return string(s) }

package sparql

import (
	"strings"
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
  ?school ?what 5.5, true, false, 12345 .
  ?school sch-ont:establishmentName ?name.
  ?school sch-ont:districtAdministrative <http://statistics.data.gov.uk/id/local-authority-district/00AA>.
  ?name 
  	a 'school'^^xs:string;
	a "double school"@EN-U5V5-Z7A7
   .
   ?name a ont:xyz.
   _:x a [].
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
	for i, q := range strings.Split(simpleQuery, "\n") {
		t.Log(i+1, q)
	}
	ast, err := parse(t, simpleQuery)
	if err != nil {
		t.Fatal("simple should not err, but", err)
	}

	w(t, "base", ast.Base, `<http://xyz>`)
	if len(ast.PrefixedIRIs) != 2+3 {
		t.Errorf("length prefix: want %d, got %d", 2+3, len(ast.PrefixedIRIs))
	} else {
		w(t, "prefix3", ast.PrefixedIRIs[3], `sch-ont:<http://education.data.gov.uk/def/school/>`)
		w(t, "prefix4", ast.PrefixedIRIs[4], `ont:<http://ontology/ÄÖÜßäöü/>`)
	}
	ws(t, "type", ast.QueryType, "select")
	ws(t, "modifier", ast.Modifier, "distinct")
	ws(t, "projection", ast.Projection.String(ast.Variables), "[name]")
	w(t, "variables", ast.Variables, "[name:$0 school:$1 what:$2]")
	if len(ast.Where) != 11 {
		for i, b := range ast.Where {
			t.Log("where", i, b.String())
		}
		t.Fatalf("want %d where blocks, got %d", 11, len(ast.Where))
	}
	w(t, "where0", ast.Where[0], `{BGP: (algebra.Variable $1) (graph.IRI <http://www.w3.org/1999/02/22-rdf-syntax-ns#type>) (graph.IRI <http://education.data.gov.uk/def/school/School>)}`)
	w(t, "where1", ast.Where[1], `{BGP: (algebra.Variable $1) (algebra.Variable $2) (graph.Float 5.5)}`)
	w(t, "where2", ast.Where[2], `{BGP: (algebra.Variable $1) (algebra.Variable $2) (graph.Bool true)}`)

	t.Error("TODO")
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

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
		t.Error("bad should err, but it dosn't")
	}
}

func TestSparqlParserSimple(t *testing.T) {
	ast, err := parse(t, simpleQuery)
	if err != nil {
		t.Error("simple should not err, but", err)
	}

	want := `http://xyz`
	if ast.Base.URI() != want {
		t.Errorf("base: want %s, got %s", want, ast.Base)
	}
}


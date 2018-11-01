package graph

import (
	"fmt"
	"strings"
)

// Literal is a string with opional language|datatype
type Literal struct {
	Text string
	LanguageTag
	DatatypeTag
}

// LiteralFrom constructs a Literal
func LiteralFrom(quoted, lang string, datatype IRI) *Literal {
	quoted = strings.TrimPrefix(quoted, `"`)
	quoted = strings.TrimPrefix(quoted, `'`)
	quoted = strings.TrimSuffix(quoted, `"`)
	quoted = strings.TrimSuffix(quoted, `'`)
	lang = strings.TrimPrefix(lang, "@")
	return &Literal{
		quoted,
		LanguageTag(lang),
		DatatypeTag(datatype),
	}
}

func (lit *Literal) String() string {
	if lit.LanguageTag != "" {
		return fmt.Sprintf("%q@%s", lit.Text, lit.LanguageTag)
	}
	if IRI(lit.DatatypeTag) != NotAnIRI {
		return fmt.Sprintf("%q^^%s", lit.Text, IRI(lit.DatatypeTag))
	}
	return fmt.Sprintf("%q", lit.Text)
}

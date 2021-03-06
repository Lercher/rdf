package values

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// Literal is a string with opional language|datatype
type Literal struct {
	Text string
	LanguageTag
	DatatypeTag
}

// LiteralString constructs a Literal
func LiteralString(s string) *Literal {
	return &Literal{
		s,
		LanguageTag(""),
		DatatypeTag(NotAnIRI),
	}
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

// Serialize to a typebyte and byte slice
func (lit *Literal) Serialize() (byte, []byte) {
	buf := new(bytes.Buffer)
	buf.Write(serializeString(lit.Text))
	buf.Write(serializeString(string(lit.LanguageTag)))
	buf.Write(serializeString(string(lit.DatatypeTag)))
	return 'l', buf.Bytes()
}

func constructLiteral(r io.Reader) (Value, error) {
	tx, err := readString(r)
	if err != nil {
		return nil, err
	}

	lang, err := readString(r)
	if err != nil {
		return nil, err
	}

	datatype, err := readString(r)
	if err != nil {
		return nil, err
	}

	return &Literal{
		tx,
		LanguageTag(lang),
		DatatypeTag(datatype),
	}, nil
}

// Inner returns the inner primitive of this Value
func (lit *Literal) Inner() interface{} {
	return lit.Text
}

// IsSameTypeAndLessThan compares this with another Value
// TODO: Literals don't care about their datatype, so they are compared by their string representation
func (lit *Literal) IsSameTypeAndLessThan(other Value) (bool, bool) {
	l2, ok := other.(*Literal)
	if ok {
		return true, lit.Text < l2.Text || string(lit.LanguageTag) < string(l2.LanguageTag) || string(lit.DatatypeTag) < string(l2.DatatypeTag)
	}
	s2, ok := other.(String)
	if ok {
		return true, lit.Text < string(s2)
	}
	return false, false
}

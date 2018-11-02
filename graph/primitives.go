package graph

import "fmt"

// make Verb'able primitive datatypes

// Int is a Verb'able int
type Int int

// Float is a Verb'able float64/double
type Float float64

// String is a Verb'able string
type String string

// Bool is a Verb'able boolean
type Bool bool

// Nil is a Verb'able nil
type Nil struct{}

// NIL is the only plausible instance of Nil
var NIL = Nil{}

func (i Int) String() string {
	return fmt.Sprint(int(i))
}

func (f Float) String() string {
	return fmt.Sprint(float64(f))
}

func (s String) String() string {
	return string(s)
}

func (b Bool) String() string {
	return fmt.Sprint(bool(b))
}

func (n Nil) String() string {
	return "()"
}

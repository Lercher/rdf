package values

import "io"

// Nil is a Verb'able nil
type Nil struct{}

// NIL is the only plausible instance of Nil
var NIL = Nil{}

func (n Nil) String() string {
	return "()"
}

// Serialize to typebyte '\0' and a nil byte slice
func (n Nil) Serialize() (byte, []byte) {
	return 0, nil
}

func constructNil(r io.Reader) (Value, error) {
	return NIL, nil
}

// Inner returns the inner primitive of this Value
func (n Nil) Inner() interface{} {
	return nil
}

// IsSameTypeAndLessThan compares this with another Value
func (n Nil) IsSameTypeAndLessThan(other Value) (bool, bool) {
	_, ok := other.(Nil)
	return ok, false
}

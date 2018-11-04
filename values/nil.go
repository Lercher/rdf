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

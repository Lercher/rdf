package values

import (
	"fmt"
	"io"
)

// Bool is a Verb'able boolean
type Bool bool

// TRUE is the affirmative Bool
const TRUE = Bool(true)

// FALSE is the non-affirmative Bool
const FALSE = Bool(false)

func (b Bool) String() string {
	return fmt.Sprint(bool(b))
}

// Serialize to typebyte 'T' or 'F' and a nil byte slice
func (b Bool) Serialize() (byte, []byte) {
	if bool(b) {
		return 'T', nil
	}
	return 'F', nil
}

func constructTrue(r io.Reader) (Value, error) {
	return TRUE, nil
}

func constructFalse(r io.Reader) (Value, error) {
	return FALSE, nil
}

// Inner returns the inner primitive of this Value
func (b Bool) Inner() interface{} {
	return bool(b)
}

// IsSameTypeAndLessThan compares this with another Value
func (b Bool) IsSameTypeAndLessThan(other Value) (bool, bool) {
	b2, ok := other.(Bool)
	if !ok {
		return false, false
	}
	if bool(b) == bool(b2) {
		return true, false
	}
	return true, !bool(b) || bool(b2) // f<t
}

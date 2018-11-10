package values

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Int is a Verb'able int
type Int int64

func (i Int) String() string {
	return fmt.Sprint(int(i))
}

// Serialize to a typebyte and byte slice
func (i Int) Serialize() (byte,[]byte) {
	b := make([]byte, 10)
	n := binary.PutVarint(b, int64(i))
	return 'i', b[:n]
}

func constructInt(r io.Reader) (Value, error) {
	br := toByteReader(r)
	i64, err := binary.ReadVarint(br)
	return Int(int(i64)), err
}

// Inner returns the inner primitive of this Value
func (i Int) Inner() interface{} {
	return int(i)
}

// IsSameTypeAndLessThan compares this with another Value
func (i Int) IsSameTypeAndLessThan(other Value) (bool, bool) {
	i2, ok := other.(Int)
	if !ok {
		f2, ok := other.(Float)
		if !ok {
			return false, false
		}
		return ok, float64(int(i)) < float64(f2)	
	}
	return ok, int(i) < int(i2)
}

package values

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

// Float is a Verb'able float64/double
type Float float64

func (f Float) String() string {
	return fmt.Sprint(float64(f))
}

// Serialize to a typebyte and byte slice
func (f Float) Serialize() (byte, []byte) {
	b := make([]byte, 10)
	n := binary.PutUvarint(b, math.Float64bits(float64(f)))
	return 'f', b[:n]
}

func constructFloat(r io.Reader) (Value, error) {
	br := toByteReader(r)
	ui64, err := binary.ReadUvarint(br)
	return Float(math.Float64frombits(ui64)), err
}

// Inner returns the inner primitive of this Value
func (f Float) Inner() interface{} {
	return float64(f)
}

// IsSameTypeAndLessThan compares this with another Value
func (f Float) IsSameTypeAndLessThan(other Value) (bool, bool) {
	f2, ok := other.(Float)
	if !ok {
		i2, ok := other.(Int)
		if !ok {
			return false, false
		}
		return ok, float64(f) < float64(i2)	
	}
	return ok, float64(f) < float64(f2)
}

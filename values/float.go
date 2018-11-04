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

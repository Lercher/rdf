package graph

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Virtual is the datatype for virtualizing values inside a Graph.
type Virtual int32

// Virtual2 is a pair of Virtuals
type Virtual2 struct {
	V1, V2 Virtual
}

// IsNil holds true, if the underlying datatype is zero, i.e. 4 zero bytes
func (v Virtual) IsNil() bool {
	return v == 0
}

func (v Virtual) String(g *Graph) string {
	val := v.Value(g)
	return fmt.Sprintf("(%T:%[1]v)", val)
}

// Value returns the real Value of a virtualized Triple part in subject/predicate/object position
// according to the records of the creating Graph. It is undefined behaviour to
// retrieve a Virtual's Value from a different Graph
func (v Virtual) Value(g *Graph) Value {
	val := g.valuelist[v-1]
	return val
}

// Encode writes the binary representation
func (v Virtual) Encode(w io.Writer) error {
	_, err := w.Write(v.Bytes())
	return err
}

// Bytes gets a copy of the bytes for this hash, see also VirtualFrom
func (v Virtual) Bytes() []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, uint32(v))
	return b
}

// VirtualDecode reads a virtualized Value
func VirtualDecode(r io.Reader) (Virtual, error) {
	b := make([]byte, 4)
	n, err := io.ReadFull(r, b)
	if err != nil {
		return 0, err
	}
	if n != len(b) {
		return 0, fmt.Errorf("wanted to read a %d byte hash but got only %d bytes", len(b), n)
	}
	return VirtualFrom(b), nil
}

// VirtualFrom creates a Virtual from 16 bytes, see also Bytes
func VirtualFrom(b []byte) Virtual {
	lower := binary.LittleEndian.Uint32(b)
	return Virtual(lower)
}

// Pattern creates a literal Pattern from this Virtual
func (v Virtual) Pattern() *Pattern {
	return PatternLiteral(v)
}

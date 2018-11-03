package graph

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Virtual is the datatype for virtualizing values inside a Graph. It's the index+1 into a slice ov Values, so 0 is Nil
type Virtual int32

// WontMatchVirtual represents a Virtual that won't match, because it is no member of the graph
const WontMatchVirtual = Virtual(-1)

// NotAVirtual is a Virtual Value that doesn't represent a Virtual
const NotAVirtual = Virtual(0)

// Virtual2 is a pair of Virtuals
type Virtual2 struct {
	V1, V2 Virtual
}

// IsNil holds true, if the underlying datatype is zero, i.e. 4 zero bytes
func (v Virtual) IsNil() bool {
	return v == 0
}

// WontMatch holds true, if the Virtual Represents something that is not a member of a graphs Values
func (v Virtual) WontMatch() bool {
	return v == -1
}

func (v Virtual) String(g *Graph) string {
	val := v.Value(g)
	return fmt.Sprintf("(%T:%[1]v)", val)
}

// Value returns the real Value of a virtualized Triple part in subject/predicate/object position
// according to the records of the creating Graph. It is undefined behaviour to
// retrieve a Virtual's Value from a different Graph
func (v Virtual) Value(g *Graph) Value {
	index := int(v)-1
	if 0 <= index && index < len(g.valuelist) {
		val := g.valuelist[index]
		return val
	}
	return nil
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

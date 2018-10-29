package graph

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/zhenjl/cityhash"
)

// Virtual is the datatype for virtualizing values inside a Graph.
// It consists of 16 bytes and is formed from the Cityhash of a binary representation
// of a Value
type Virtual cityhash.Uint128

// Virtual2 is a pair of Virtuals
type Virtual2 struct {
	V1, V2 Virtual
}

// IsNil holds true, if the underlying datatype is zero, i.e. 16 zero bytes
func (v Virtual) IsNil() bool {
	return cityhash.Uint128(v).Higher64() == 0 && cityhash.Uint128(v).Lower64() == 0
}

func (v Virtual) String(g *Graph) string {
	val := v.Value(g)
	return fmt.Sprintf("(%T:%[1]v)", val)
}

// Value returns the real Value of a virtualized Triple part in subject/predicate/object position
// according to the records of the creating Graph. It is undefined behaviour to
// retrieve a Virtual's Value from a different Graph
func (v Virtual) Value(g *Graph) Value {
	val := g.valuemap[v]
	return val
}

func hash(seedByte byte, block []byte) Virtual {
	seed := cityhash.Uint128{uint64(seedByte), uint64(seedByte)}
	s := cityhash.CityHash128WithSeed(block, uint32(len(block)), seed)
	return Virtual(s)
}

// Encode writes the binary representation
func (v Virtual) Encode(w io.Writer) error {
	_, err := w.Write(v.Bytes())
	return err
}

// Bytes gets a copy of the bytes for this hash, see also VirtualFrom
func (v Virtual) Bytes() []byte {
	return cityhash.Uint128(v).Bytes()
}

// VirtualDecode reads a virtualized Value
func VirtualDecode(r io.Reader) (Virtual, error) {
	b := make([]byte, 16)
	n, err := io.ReadFull(r, b)
	if err != nil {
		return Virtual{}, err
	}
	if n != len(b) {
		return Virtual{}, fmt.Errorf("wanted to read a %d byte hash but got only %d bytes", len(b), n)
	}
	return VirtualFrom(b), nil
}

// VirtualFrom creates a Virtual from 16 bytes, see also Bytes
func VirtualFrom(b []byte) Virtual {
	lower := binary.LittleEndian.Uint64(b[:8])
	upper := binary.LittleEndian.Uint64(b[8:])
	ui := cityhash.Uint128{lower, upper}
	return Virtual(ui)
}

// Pattern creates a literal Pattern from this Virtual
func (v Virtual) Pattern() *Pattern {
	return PatternLiteral(v)
}
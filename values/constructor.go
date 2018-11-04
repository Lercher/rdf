package values

import (
	"fmt"
	"io"
)

// Constructor reads from a stream to construct a hashable Value from it
type Constructor func(r io.Reader) (Value, error)

var readers = make(map[byte]Constructor)

// RegisterConstructor registers a constructor function under a typebyte.
// The constructed Hashable needs to Serialize under this typebyte again.
func RegisterConstructor(typebyte byte, cons Constructor) {
	if _, ok := readers[typebyte]; ok {
		panic(fmt.Sprintf("typebyte %v is used twice", typebyte))
	}
	readers[typebyte] = cons
}

// ConstructNext reads the type byte, looks for its constructor function and calls it
func ConstructNext(r io.Reader) (Value, error) {
	tb := []byte{0}
	n, err := r.Read(tb)
	if err != nil {
		return nil, fmt.Errorf("can't read type byte: %v", err)
	}
	if n != len(tb) {
		return nil, fmt.Errorf("can't read type byte, n=%d", n)
	}
	typebyte := tb[0]
	cons, ok := readers[typebyte]
	if !ok {
		return nil, fmt.Errorf("no registered constructor for type byte %v", typebyte)
	}
	return cons(r)
}

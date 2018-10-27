package graph

import (
	"fmt"
	"io"
	"encoding/binary"
)

// Value is virtualized and can be stored in a Graphs assertions
type Value interface{}

// ValueEncode serializes a value: 1 byte type, 16 bytes for value of int/float64 or string byte length of string
func ValueEncode(w io.Writer, v Value) error {
	switch val := v.(type) {
	case int:
		return encodeVal(w, valueTypeInt, uint64(val))
	case float64:
		return encodeVal(w, valueTypeFloat, uint64(val))
	case string:
		b := []byte(val)
		err := encodeVal(w, valueTypeString, uint64(len(b)))
		if err != nil {
			return err
		}
		_, err = w.Write(b)
		return err
	default:
		return fmt.Errorf("encode: illegal datatype %T:%[1]v", v)
	}
}

func encodeVal(w io.Writer, typ byte, payload uint64) error {
	b := make([]byte, 8)
	b[0] = typ
	_, err := w.Write(b[:1])
	if err != nil {
		return err
	}
	binary.LittleEndian.PutUint64(b, payload)
	_, err = w.Write(b)
	return err
}
package graph

import (
	"encoding/binary"
	"fmt"
	"io"
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

func ValueDecode(r io.Reader) (interface{}, error) {
	b8 := make([]byte, 8)
	b1 := b8[:1]
	n, err := io.ReadFull(r, b1)
	if err != nil {
		return nil, err
	}
	if n != len(b1) {
		return 0, fmt.Errorf("wanted to read %d byte type, got only %d", len(b1), n)
	}
	t := b1[0]
	n, err = io.ReadFull(r, b8)
	if err != nil {
		return nil, err
	}
	if n != len(b8) {
		return nil, fmt.Errorf("wanted to read %d bytes size/value, got only %d", len(b8), n)
	}
	ui := binary.LittleEndian.Uint64(b8)
	switch t {
	default:
		return nil, fmt.Errorf("invalid type byte %d", t)
	case valueTypeFloat:
		return float64(ui), nil
	case valueTypeInt:
		return int(ui), nil
	case valueTypeString:
		buf := make([]byte, int(ui))
		n, err := io.ReadFull(r, buf)
		if err != nil {
			return nil, err
		}
		if n != len(buf) {
			return nil, fmt.Errorf("wanted to read %d bytes of utf8 string, got only %d", len(buf), n)
		}
		return string(buf), nil
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

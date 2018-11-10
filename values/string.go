package values

import (
	"encoding/binary"
	"fmt"
	"io"
)

// String is a Verb'able string
type String string

func (s String) String() string {
	return string(s)
}

func serializeString(s string) []byte {
	bs := []byte(s)
	b := make([]byte, 10 + len(bs))
	n := binary.PutVarint(b, int64(len(bs)))
	copy(b[n:], bs)
	return b[:n+len(bs)]
}

func readString(r io.Reader) (string, error) {
	br := toByteReader(r)
	length64, err := binary.ReadVarint(br)
	if err != nil {
		return "", fmt.Errorf("can't read string length: %v", err)
	}

	buf := make([]byte, int(length64))
	n, err := io.ReadFull(r, buf)
	if err != nil {
		return "", err
	}
	if n != len(buf) {
		return "", fmt.Errorf("wanted to read %d bytes of utf8 string, got only %d", len(buf), n)
	}

	return string(buf), nil
}

// Serialize to a typebyte and byte slice
func (s String) Serialize() (byte, []byte) {
	return 's', serializeString(string(s))
}

func constructString(r io.Reader) (Value, error) {
	s, err := readString(r)
	return String(s), err
}

// Inner returns the inner primitive of this Value
func (s String) Inner() interface{} {
	return string(s)
}

// IsSameTypeAndLessThan compares this with another Value
func (s String) IsSameTypeAndLessThan(other Value) (bool, bool) {
	s2, ok := other.(String)
	if !ok {
		l2, ok := other.(*Literal)
		if ! ok {
			return false, false
		}
		return true, string(s) < l2.Text
	}
	return true, string(s) < string(s2)
}

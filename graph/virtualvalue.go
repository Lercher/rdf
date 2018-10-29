package graph

import (
	"encoding/binary"
	"fmt"
)

// VirtualValue pairs a Virtual (i.e. the hash value) with its real Value
type VirtualValue struct {
	Virtual
	Value
	Size  int
	Known bool
}

// newVirtualValue creates a VirtualValue from a string, int or float64 and associates it with a Graph.
// Any other type raises a panic.
func newVirtualValue(g *Graph, primitive interface{}) VirtualValue {
	switch v := primitive.(type) {
	case int:
		return virtualValueInt(g, v)
	case float64:
		return virtualValueFloat(g, v)
	case string:
		return virtualValueString(g, v)
	default:
		panic(fmt.Errorf("unsupported value type %T", primitive))
	}
}

// virtualValueString creates a VirtualValue from a string
func virtualValueString(g *Graph, s string) VirtualValue {
	return vvalue(g, valueTypeString, []byte(s), s)
}

// virtualValueInt creates a VirtualValue from an int
func virtualValueInt(g *Graph, i int) VirtualValue {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(i))
	b = append(b, valueTypeInt)
	return vvalue(g, valueTypeInt, b, i)
}

// virtualValueFloat creates a VirtualValue from a float64
func virtualValueFloat(g *Graph, f float64) VirtualValue {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(f))
	return vvalue(g, valueTypeFloat, b, f)
}

func vvalue(g *Graph, seed byte, bytes []byte, v interface{}) VirtualValue {
	h := hash(seed, bytes)
	val, ok := g.valuemap[h]
	if !ok {
		val = v
	}
	return VirtualValue{
		Virtual: h,
		Value:   val,
		Known:   ok,
		Size:    len(bytes),
	}
}

// Pattern returns a constant Pattern from the virtual value
func (vv VirtualValue) Pattern() *Pattern {
	return PatternLiteral(vv.Virtual)
}

const (
	valueTypeString = byte('S')
	valueTypeInt    = byte('I')
	valueTypeFloat  = byte('F')
)

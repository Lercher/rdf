package graph

import (
	"encoding/binary"
	"fmt"
)

// VirtualValue pairs a Virtual (i.e. the hash value) with its real Value
type VirtualValue struct {
	Virtual
	Value
}

// NewVirtualValue creates a VirtualValue from a string, int or float64 and associates it with a Graph.
// Any other type raises a panic.
func NewVirtualValue(g *Graph, primitive interface{}) VirtualValue {
	switch v := primitive.(type) {
	case int:
		return VirtualValueInt(g, v)
	case float64:
		return VirtualValueFloat(g, v)
	case string:
		return VirtualValueString(g, v)
	default:
		panic(fmt.Errorf("unsupported value type %T", primitive))
	}
}

// VirtualValueString creates a VirtualValue from a string
func VirtualValueString(g *Graph, s string) VirtualValue {
	return vvalue(g, valueTypeString, []byte(s), s)
}

// VirtualValueInt creates a VirtualValue from an int
func VirtualValueInt(g *Graph, i int) VirtualValue {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(i))
	b = append(b, valueTypeInt)
	return vvalue(g, valueTypeInt, b, i)
}

// VirtualValueFloat creates a VirtualValue from a float64
func VirtualValueFloat(g *Graph, f float64) VirtualValue {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(f))
	return vvalue(g, valueTypeFloat, b, f)
}

func vvalue(g *Graph, seed byte, bytes []byte, v interface{}) VirtualValue {
	h := hash(seed, bytes)
	val, ok := g.valuemap[h]
	if !ok {
		val = Value(v)
		g.valuemap[h] = val
	}
	return VirtualValue{
		Virtual: h,
		Value:   val,
	}
}

const (
	valueTypeString = byte('S')
	valueTypeInt= byte('I')
	valueTypeFloat = byte('F')
)

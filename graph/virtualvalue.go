package graph

import (
	"encoding/binary"
	"fmt"
)

// VirtualValue pairs a Virtual (i.e. the hash value) with its real Value
type VirtualValue struct {
	vhash
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
		// #TODO: must support iriRef|rdfLiteral|numericLiteral|booleanLiteral|blankNode|NIL
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
	val := v
	h := hash(seed, bytes)
	idx1, ok := g.vhash2indexPlus1[h]
	if ok {
		val = g.valuelist[idx1-1]
	}
	return VirtualValue{
		vhash:   h,
		Virtual: idx1,
		Value:   val,
		Known:   ok,
		Size:    len(bytes),
	}
}

const (
	valueTypeString = byte('S')
	valueTypeInt    = byte('I')
	valueTypeFloat  = byte('F')
)

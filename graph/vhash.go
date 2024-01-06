package graph

import (
	"github.com/zhenjl/cityhash"
)

// It consists of 16 bytes and is formed from the Cityhash of a binary representation
// of a Value
type vhash cityhash.Uint128

func hash(seedByte byte, block []byte) vhash {
	seed := cityhash.Uint128{uint64(seedByte), uint64(seedByte)}
	s := cityhash.CityHash128WithSeed(block, uint32(len(block)), seed)
	return vhash(s)
}

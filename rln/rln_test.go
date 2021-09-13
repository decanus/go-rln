package rln

import "testing"

func TestHash(t *testing.T) {
	rln := &RLN{}
	rln.Hash([]byte{1, 2, 3})
}

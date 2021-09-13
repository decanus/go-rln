package rln

import (
	"testing"
)

func TestHash(t *testing.T) {
	rln := New(0, []byte{})
	rln.Hash([]byte{1, 2, 3})
}

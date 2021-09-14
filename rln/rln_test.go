package rln_test

import (
	"encoding/hex"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/decanus/go-rln/rln"
)

func TestNew(t *testing.T) {
	params, err := ioutil.ReadFile("./testdata/parameters.key")
	if err != nil {
		t.Fatal(err)
	}

	_, err = rln.New(32, params)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGenerateKey(t *testing.T) {
	params, err := ioutil.ReadFile("./testdata/parameters.key")
	if err != nil {
		t.Fatal(err)
	}

	r, err := rln.New(32, params)
	if err != nil {
		t.Fatal(err)
	}

	k, err := r.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}

	if reflect.DeepEqual(k.Key, [32]byte{}) {
		t.Fatal("k.Key was empty")
	}

	if reflect.DeepEqual(k.Commitment, [32]byte{}) {
		t.Fatal("k.Commitment was empty")
	}
}

func TestRLN_Hash(t *testing.T) {
	// This test is based on tests from:
	// https://github.com/status-im/nim-waku/blob/b7998de09d1ef04599a699938da69aecfa63cc6f/tests/v2/test_waku_rln_relay.nim#L527

	params, err := ioutil.ReadFile("./testdata/parameters.key")
	if err != nil {
		t.Fatal(err)
	}

	r, err := rln.New(32, params)
	if err != nil {
		t.Fatal(err)
	}

	input := byteArray(32, 1)

	output, err := r.Hash(input)
	if err != nil {
		t.Fatal(err)
	}

	expected := "53a6338cdbf02f0563cec1898e354d0d272c8f98b606c538945c6f41ef101828"
	if expected != hex.EncodeToString(output) {
		t.Fatalf("value %x did not match expected %s", output, expected)
	}
}

func TestRLN_GetRoot(t *testing.T) {
	// This test is based on tests from:
	// https://github.com/status-im/nim-waku/blob/b7998de09d1ef04599a699938da69aecfa63cc6f/tests/v2/test_waku_rln_relay.nim#L320

	params, err := ioutil.ReadFile("./testdata/parameters.key")
	if err != nil {
		t.Fatal(err)
	}

	r, err := rln.New(32, params)
	if err != nil {
		t.Fatal(err)
	}

	root1, err := r.GetRoot()
	if err != nil {
		t.Fatal(err)
	}

	root2, err := r.GetRoot()
	if err != nil {
		t.Fatal(err)
	}

	if hex.EncodeToString(root1) != hex.EncodeToString(root2) {
		t.Fatalf("value %x did not match expected %x", root1, root2)
	}
}

func byteArray(length int, value byte) []byte {
	arr := make([]byte, length)

	for i := 0; i < length; i++ {
		arr[i] = value
	}

	return arr
}

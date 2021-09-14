package rln_test

import (
	"encoding/hex"
	"fmt"
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

	fmt.Print(hex.EncodeToString(output))
}

func byteArray(length int, value byte) []byte {
	arr := make([]byte, 0)

	for i := 0; i < length; i++ {
		arr = append(arr, value)
	}

	return arr
}

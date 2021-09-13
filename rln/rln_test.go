package rln_test

import (
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

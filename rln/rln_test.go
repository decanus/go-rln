package rln

import (
	"io/ioutil"
	"testing"
)

func TestNew(t *testing.T) {
	params, err := ioutil.ReadFile("./testdata/parameters.key")
	if err != nil {
		t.Fatal(err)
	}

	rln := New(32, params)

	if rln.ptr == nil {
		t.Fatal("pointer not initialized.")
	}
}

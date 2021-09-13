package rln_test

import (
	"io/ioutil"
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

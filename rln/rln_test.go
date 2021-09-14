package rln_test

import (
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/decanus/go-rln/rln"
)

func TestRLN_New(t *testing.T) {
	params, err := ioutil.ReadFile("./testdata/parameters.key")
	if err != nil {
		t.Fatal(err)
	}

	_, err = rln.New(32, params)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRLN_GenerateKey(t *testing.T) {
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

func TestRLN_Verify(t *testing.T) {
	// This test is based on a similar test in nim-waku:
	// https://github.com/status-im/nim-waku/blob/b7998de09d1ef04599a699938da69aecfa63cc6f/tests/v2/test_waku_rln_relay.nim#L559
	// as well as tests in the rust repository:
	// https://github.com/kilic/rln/blob/master/src/ffi.rs#L158

	params, err := ioutil.ReadFile("./testdata/parameters.key")
	if err != nil {
		t.Fatal(err)
	}

	r, err := rln.New(32, params)
	if err != nil {
		t.Fatal(err)
	}

	index := 5

	auth, err := r.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		if i == index {
			err := r.UpdateNextMember(auth.Commitment[:])
			if err != nil {
				t.Fatalf("failed to add member: %s", err)
			}
		} else {
			key, err := r.GenerateKey()
			if err != nil {
				t.Fatal(err)
			}

			err = r.UpdateNextMember(key.Commitment[:])
			if err != nil {
				t.Fatalf("failed to add member: %s", err)
			}
		}
	}

	// @TODO create inputs

}

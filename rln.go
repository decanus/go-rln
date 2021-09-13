package main

import (
	"fmt"

	"github.com/decanus/go-rln/rln"
)

func main() {
	r := &rln.RLN{}
	test := []byte{}
	r.GenerateKey(test)
	fmt.Printf("Test: %v", test)
	//C
}

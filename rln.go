package main

/*
//#cgo CFLAGS: -I../lib
#cgo LDFLAGS: -L./lib -llibrnl
#include "./lib/librln.h"
*/
import "C"

func main() {
	C.verify()
}

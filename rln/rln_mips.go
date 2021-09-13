// +build linux,mips

package rln

/*
#cgo LDFLAGS: -L${SRCDIR}/../libs/mips-unknown-linux-gnu -lrln -ldl -lm
*/
import "C"

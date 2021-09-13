// +build linux,mips64

package rln

/*
#cgo LDFLAGS: -L${SRCDIR}/../libs/mips64-unknown-linux-gnuabi64 -lrln -ldl -lm
*/
import "C"

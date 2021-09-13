// +build linux,arm,!arm7

package rln

/*
#cgo LDFLAGS: -L${SRCDIR}/../libs/arm-unknown-linux-gnueabi -lrln -ldl -lm
*/
import "C"

// +build linux,arm64

package rln

/*
#cgo LDFLAGS: -L${SRCDIR}/../libs/aarch64-unknown-linux-gnu -lrln -ldl -lm
*/
import "C"

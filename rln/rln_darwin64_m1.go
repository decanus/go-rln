// +build darwin,arm64,!ios

package rln

/*
#cgo LDFLAGS: -L${SRCDIR}/../libs/aarch64-apple-darwin -lrln -ldl -lm
*/
import "C"

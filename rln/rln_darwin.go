// +build darwin,386,!ios

package rln

/*
#cgo LDFLAGS: -L${SRCDIR}/../libs/i686-apple-darwin -lrln -ldl -lm
*/
import "C"

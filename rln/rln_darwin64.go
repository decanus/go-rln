// +build darwin,amd64,!ios

package rln

/*
#cgo LDFLAGS: -L${SRCDIR}/../libs/x86_64-apple-darwin -lrln -ldl -lm
*/
import "C"

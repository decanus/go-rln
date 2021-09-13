// +build !android,musl

package rln

/*
#cgo LDFLAGS: -L${SRCDIR}/../libs/x86_64-unknown-linux-musl -lrln -ldl -lm
*/
import "C"

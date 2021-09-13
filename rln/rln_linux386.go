// +build !android,linux,386,!musl

package rln

/*
#cgo LDFLAGS: -L${SRCDIR}/../libs/i686-unknown-linux-gnu -lrln -ldl -lm -lpthread
*/
import "C"

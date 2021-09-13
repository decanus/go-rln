// +build windows,386

package rln

/*
#cgo LDFLAGS: -L${SRCDIR}/../libs/i686-pc-windows-gnu -lrln -lm -lws2_32 -luserenv
*/
import "C"

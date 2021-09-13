// +build windows,amd64

package rln

/*
#cgo LDFLAGS: -L${SRCDIR}/../libs/x86_64-pc-windows-gnu -lrln -lm -lws2_32 -luserenv
*/
import "C"

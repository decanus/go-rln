// +build android

package rln

/*
#cgo LDFLAGS: -L${SRCDIR}/../libs/aarch64-linux-android -L${SRCDIR}/../libs/armv7-linux-androideabi -L${SRCDIR}/../libs/i686-linux-android -L${SRCDIR}/../libs/x86_64-linux-android -lrln -ldl -lm
*/
import "C"

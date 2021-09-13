// +build linux,mips64le

package rln

/*
#cgo LDFLAGS: -L${SRCDIR}/../libs/mips64el-unknown-linux-gnuabi64 -lrln -ldl -lm
*/
import "C"

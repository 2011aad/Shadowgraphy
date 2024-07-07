//go:build linux

package fpe

// #cgo LDFLAGS: -L${SRCDIR}/lib -lshadow_fpe_export -Wl,-rpath=${SRCDIR}/lib
import "C"

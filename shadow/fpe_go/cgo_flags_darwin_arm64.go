//go:build darwin && arm64

package fpe

// #cgo LDFLAGS: -L${SRCDIR}/lib -lshadow_fpe_export_arm -Wl,-rpath,${SRCDIR}/lib
import "C"

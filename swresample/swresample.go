// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

// Package swresample provides a high-level interface to the libswresample library audio resampling utilities
// The process of changing the sampling rate of a discrete signal to obtain a new discrete representation of the underlying continuous signal.
package swresample

/*
	#cgo pkg-config: libswresample
	#include <libswresample/swresample.h>
*/
import "C"

type (
	Context      C.struct_SwrContext
	Frame        C.struct_AVFrame
	Class        C.struct_AVClass
	SampleFormat C.enum_AVSampleFormat
)

// GetClass gets the Class for Context.
//
// C-Function:
func GetClass() *Class {
	return (*Class)(C.swr_get_class())
}

// NewContext allocates Context.
//
// C-Function: swr_alloc
func NewContext() *Context {
	return (*Context)(C.swr_alloc())
}

// Version returns the swresample version.
//
// C-Function: swresample_version
func Version() uint {
	return uint(C.swresample_version())
}

// Configuration returns the build time configuration of swresample.
//
// C-Function: swresample_configuration
func Configuration() string {
	return C.GoString(C.swresample_configuration())
}

// License returns the swresample license.
//
// C-Function: swresample_license
func License() string {
	return C.GoString(C.swresample_license())
}

// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)
// Corbatto (luca@corbatto.de)

// Package avutil is a utility library to aid portable multimedia programming.
// It contains safe portable string functions, random number generators, data structures,
// additional mathematics functions, cryptography and multimedia related functionality.
// Some generic features and utilities provided by the libavutil library
package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/avutil.h>
//#include <stdlib.h>
import "C"
import (
	"unsafe"
)

type (
	Options     C.struct_AVOptions
	Dictionary  C.struct_AVDictionary
	Tree        C.struct_AVTree
	Rational    C.struct_AVRational
	MediaType   C.enum_AVMediaType
	PictureType C.enum_AVPictureType
	File        C.FILE
)

// Version returns the LIBAVUTIL_VERSION_INT constant.
//
// C-Function: avutil_version
func Version() uint {
	return uint(C.avutil_version())
}

// Configuration returns the libavutil build-time configuration.
//
// C-Function: avutil_configuration
func Configuration() string {
	return C.GoString(C.avutil_configuration())
}

// License returns the libavutil license.
//
// C-Function: avutil_license
func License() string {
	return C.GoString(C.avutil_license())
}

// MediaTypeString returns a string describing the media_type enum, NULL if media_type is unknown.
//
// C-Function: av_get_media_type_string
func MediaTypeString(mt MediaType) string {
	return C.GoString(C.av_get_media_type_string((C.enum_AVMediaType)(mt)))
}

// PictureTypeChar returns a single letter to describe the given picture type pict_type.
//
// C-Function: av_get_picture_type_char
func PictureTypeChar(pt PictureType) string {
	return string(C.av_get_picture_type_char((C.enum_AVPictureType)(pt)))
}

// XIfNull returns x default pointer in case p is NULL.
//
// C-Function: av_x_if_null
func XIfNull(p, x int) {
	C.av_x_if_null(unsafe.Pointer(&p), unsafe.Pointer(&x))
}

// IntListLengthForSize computes the length of an integer list.
//
// C-Function: av_int_list_length_for_size
func IntListLengthForSize(e uint, l int, t uint64) uint {
	return uint(C.av_int_list_length_for_size(C.uint(e), unsafe.Pointer(&l), (C.uint64_t)(t)))
}

// FopenUtf8 opens a file using a UTF-8 filename.
//
// C-Function: av_fopen_utf8
func FopenUtf8(p, m string) *File {
	f := C.av_fopen_utf8(C.CString(p), C.CString(m))
	return (*File)(f)
}

// TimeBaseQ returns the fractional representation of the internal time base.
//
// C-Function: av_get_time_base_q
func TimeBaseQ() Rational {
	return (Rational)(C.av_get_time_base_q())
}

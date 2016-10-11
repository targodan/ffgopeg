// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

//Package swscale performs highly optimized image scaling and colorspace and pixel format conversion operations.
//Rescaling: is the process of changing the video size. Several rescaling options and algorithms are available.
//Pixel format conversion: is the process of converting the image format and colorspace of the image.
package swscale

//#cgo pkg-config: libswscale libavutil
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <string.h>
//#include <stdint.h>
//#include <libswscale/swscale.h>
import "C"
import (
	"unsafe"
)

type (
	Context     C.struct_SwsContext
	Filter      C.struct_SwsFilter
	Vector      C.struct_SwsVector
	Class       C.struct_AVClass
	PixelFormat C.enum_AVPixelFormat
)

// Version returns the LIBSWSCALE_VERSION_INT constant.
//
// C-Function:
//
// C-Function: swscale_version
func Version() uint {
	return uint(C.swscale_version())
}

// Configuration returns the libswscale build-time configuration.
//
// C-Function: GoString
func Configuration() string {
	return C.GoString(C.swscale_configuration())
}

// License returns the libswscale license.
//
// C-Function: GoString
func License() string {
	return C.GoString(C.swscale_license())
}

// Coefficients returns a pointer to yuv<->rgb coefficients for the given colorspace suitable for sws_setColorspaceDetails().
//
// C-Function: sws_getCoefficients
func Coefficients(c int) *int {
	return (*int)(unsafe.Pointer(C.sws_getCoefficients(C.int(c))))
}

// IsSupportedInput returns true if pix_fmt is a supported input format,false otherwise.
//
// C-Function: sws_isSupportedInput
func IsSupportedInput(p PixelFormat) bool {
	return int(C.sws_isSupportedInput((C.enum_AVPixelFormat)(p))) != 0
}

// IsSupportedOutput returns true if pix_fmt is a supported output format, false otherwise.
//
// C-Function: sws_isSupportedOutput
func IsSupportedOutput(p PixelFormat) bool {
	return int(C.sws_isSupportedOutput((C.enum_AVPixelFormat)(p))) != 0
}

// IsSupportedEndiannessConversion returns true if endianess conversion is supported.
//
// C-Function: sws_isSupportedEndiannessConversion
func IsSupportedEndiannessConversion(p PixelFormat) bool {
	return int(C.sws_isSupportedEndiannessConversion((C.enum_AVPixelFormat)(p))) != 0
}

// Scale scales the image slice in srcSlice and put the resulting scaled slice in the image in dst.
//
// C-Function: sws_scale
func (ctxt *Context) Scale(src *uint8, str int, y, h int, d *uint8, ds int) int {
	cctxt := (*C.struct_SwsContext)(unsafe.Pointer(ctxt))
	csrc := (*C.uint8_t)(unsafe.Pointer(src))
	cstr := (*C.int)(unsafe.Pointer(&str))
	cd := (*C.uint8_t)(unsafe.Pointer(d))
	cds := (*C.int)(unsafe.Pointer(&ds))
	return int(C.sws_scale(cctxt, &csrc, cstr, C.int(y), C.int(h), &cd, cds))
}

// SetColorspaceDetails sets the colorspace details.
//
// C-Function: sws_setColorspaceDetails
func (ctxt *Context) SetColorspaceDetails(it *int, sr int, t *int, dr, b, c, s int) int {
	cit := (*C.int)(unsafe.Pointer(it))
	ct := (*C.int)(unsafe.Pointer(t))
	return int(C.sws_setColorspaceDetails((*C.struct_SwsContext)(ctxt), cit, C.int(sr), ct, C.int(dr), C.int(b), C.int(c), C.int(s)))
}

// ColorspaceDetails returns the colorspace details.
//
// C-Function: sws_getColorspaceDetails
func ColorspaceDetails(ctxt *Context, it, sr, t, dr, b, c, s *int) int {
	cit := (**C.int)(unsafe.Pointer(it))
	csr := (*C.int)(unsafe.Pointer(sr))
	ct := (**C.int)(unsafe.Pointer(t))
	cdr := (*C.int)(unsafe.Pointer(dr))
	cb := (*C.int)(unsafe.Pointer(b))
	cc := (*C.int)(unsafe.Pointer(c))
	cs := (*C.int)(unsafe.Pointer(s))
	return int(C.sws_getColorspaceDetails((*C.struct_SwsContext)(ctxt), cit, csr, ct, cdr, cb, cc, cs))
}

// DefaultFilter returns the default filter.
//
// C-Function: sws_getDefaultFilter
func DefaultFilter(lb, cb, ls, cs, chs, cvs float32, v int) *Filter {
	return (*Filter)(unsafe.Pointer(C.sws_getDefaultFilter(C.float(lb), C.float(cb), C.float(ls), C.float(cs), C.float(chs), C.float(cvs), C.int(v))))
}

// Free frees the filter.
//
// C-Function: sws_freeFilter
func (f *Filter) Free() {
	C.sws_freeFilter((*C.struct_SwsFilter)(f))
}

// ConvertPalette8ToPacked32 converts an 8-bit paletted frame into a frame with a color depth of 32 bits.
//
// C-Function: sws_convertPalette8ToPacked32
func ConvertPalette8ToPacked32(s, d *uint8, px int, p *uint8) {
	C.sws_convertPalette8ToPacked32((*C.uint8_t)(s), (*C.uint8_t)(d), C.int(px), (*C.uint8_t)(p))
}

// ConvertPalette8ToPacked24 converts an 8-bit paletted frame into a frame with a color depth of 24 bits.
//
// C-Function: sws_convertPalette8ToPacked24
func ConvertPalette8ToPacked24(s, d *uint8, px int, p *uint8) {
	C.sws_convertPalette8ToPacked24((*C.uint8_t)(s), (*C.uint8_t)(d), C.int(px), (*C.uint8_t)(p))
}

// GetClass gets the Class for swsContext.
//
// C-Function: sws_get_class
func GetClass() *Class {
	return (*Class)(C.sws_get_class())
}

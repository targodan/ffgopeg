// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package swscale

//#cgo pkg-config: libswscale libavutil
//#include <libswscale/swscale.h>
import "C"
import (
	"unsafe"
)

// NewVector allocates and returns an uninitialized vector with length coefficients.
//
// C-Function: sws_allocVec
func NewVector(length int) *Vector {
	return (*Vector)(C.sws_allocVec(C.int(length)))
}

// NewGaussianVector returns a normalized Gaussian curve used to filter stuff quality = 3 is high quality, lower is lower quality.
//
// C-Function: sws_getGaussianVec
func NewGaussianVector(variance, quality float64) *Vector {
	return (*Vector)(unsafe.Pointer(C.sws_getGaussianVec(C.double(variance), C.double(quality))))
}

// Scale scales all the coefficients of a by the scalar value.
//
// C-Function: sws_scaleVec
func (v *Vector) Scale(fac float64) {
	C.sws_scaleVec((*C.struct_SwsVector)(unsafe.Pointer(v)), C.double(fac))
}

// Normalize scales all the coefficients of a so that their sum equals height.
//
// C-Function: sws_normalizeVec
func (v *Vector) Normalize(height float64) {
	C.sws_normalizeVec((*C.struct_SwsVector)(v), C.double(height))
}

// Free frees the vector.
//
// C-Function: sws_freeVec
func (a *Vector) Free() {
	C.sws_freeVec((*C.struct_SwsVector)(a))
}

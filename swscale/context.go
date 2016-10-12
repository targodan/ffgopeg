// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package swscale

//#cgo pkg-config: libswscale libavutil
//#include <libswscale/swscale.h>
import "C"
import (
	"unsafe"
)

// NewContext allocates an empty Context.
//
// C-Function: sws_alloc_context
func NewContext() *Context {
	return (*Context)(C.sws_alloc_context())
}

// Init initializes the swscaler context sws_context.
//
// C-Function: sws_init_context
func (ctxt *Context) Init(sf, df *Filter) int {
	return int(C.sws_init_context((*C.struct_SwsContext)(ctxt), (*C.struct_SwsFilter)(sf), (*C.struct_SwsFilter)(df)))
}

// Free frees the swscaler context swsContext.
//
// C-Function: sws_freeContext
func (ctxt *Context) Free() {
	C.sws_freeContext((*C.struct_SwsContext)(ctxt))
}

// NewContextOpts allocates and returns an Context.
//
// C-Function: sws_getContext
func NewContextOpts(sw, sh int, sf PixelFormat, dw, dh int, df PixelFormat, f int, sfl, dfl *Filter, p *int) *Context {
	return (*Context)(C.sws_getContext(C.int(sw), C.int(sh), (C.enum_AVPixelFormat)(sf), C.int(dw), C.int(dh), (C.enum_AVPixelFormat)(df), C.int(f), (*C.struct_SwsFilter)(sfl), (*C.struct_SwsFilter)(dfl), (*C.double)(unsafe.Pointer(p))))
}

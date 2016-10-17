// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)
// Corbatto (luca@corbatto.de)

//Package avfilter contains methods that deal with ffmpeg filters
//filters in the same linear chain are separated by commas, and distinct linear chains of filters are separated by semicolons.
//FFmpeg is enabled through the "C" libavfilter library
package avfilter

/*
	#cgo pkg-config: libavfilter
	#include <libavfilter/avfilter.h>
*/
import "C"
import (
	"unsafe"

	"github.com/targodan/ffgopeg/avutil"
)

type (
	Filter        C.struct_AVFilter
	FilterContext C.struct_AVFilterContext
	FilterLink    C.struct_AVFilterLink
	FilterGraph   C.struct_AVFilterGraph
	FilterInOut   C.struct_AVFilterInOut
	FilterPad     C.struct_AVFilterPad
	Dictionary    C.struct_AVDictionary
	Class         C.struct_AVClass
	MediaType     C.enum_AVMediaType
)

// Version returns the LIBAVFILTER_VERSION_INT constant.
//
// C-Function: avfilter_version
func Version() uint {
	return uint(C.avfilter_version())
}

// Configuration returns the libavfilter build-time configuration.
//
// C-Function: avfilter_configuration
func Configuration() string {
	return C.GoString(C.avfilter_configuration())
}

// License returns the libavfilter license.
//
// C-Function: avfilter_license
func License() string {
	return C.GoString(C.avfilter_license())
}

// FilterPadCount gets the number of elements in a NULL-terminated array of FilterPads (e.g.
//
// C-Function: avfilter_pad_count
func FilterPadCount() ([]*FilterPad, avutil.ReturnCode) {
	var p *FilterPad
	err := avutil.NewReturnCode(int(C.avfilter_pad_count((*C.struct_AVFilterPad)(p))))
	arr := (*[1 << 30]*FilterPad)(unsafe.Pointer(p))
	var len int
	// Determine length
	for i, v := range arr {
		if v == nil {
			len = i
			break
		}
	}
	return arr[:len:len], err
}

// Name returns the name of a FilterPad.
//
// C-Function: avfilter_pad_get_name
func (p *FilterPad) Name() string {
	return C.GoString(C.avfilter_pad_get_name((*C.struct_AVFilterPad)(p), C.int(0)))
}

// Type returns the type of an FilterPad.
//
// C-Function: avfilter_pad_get_type
func (p *FilterPad) Type() MediaType {
	return (MediaType)(C.avfilter_pad_get_type((*C.struct_AVFilterPad)(p), C.int(0)))
}

// Link links two filters together.
//
// C-Function: avfilter_link
func Link(s *FilterContext, sp uint, d *FilterContext, dp uint) avutil.ReturnCode {
	return avutil.NewReturnCode(int(C.avfilter_link((*C.struct_AVFilterContext)(s), C.uint(sp), (*C.struct_AVFilterContext)(d), C.uint(dp))))
}

//Free the link in *link, and set its pointer to NULL.
//
// C-Function: avfilter_link_free
func (l *FilterLink) Free() {
	C.avfilter_link_free((**C.struct_AVFilterLink)(unsafe.Pointer(&l)))
}

// Channels returns the number of channels of a link.
//
// C-Function:
func (l *FilterLink) Channels() int {
	return int(C.avfilter_link_get_channels((*C.struct_AVFilterLink)(l)))
}

// ConfigLinks negotiates the media format, dimensions, etc of all inputs to a filter.
//
// C-Function: avfilter_config_links
func (f *FilterContext) ConfigLinks() avutil.ReturnCode {
	return avutil.NewReturnCode(int(C.avfilter_config_links((*C.struct_AVFilterContext)(f))))
}

// ProcessCommand makes the filter instance process a command.
//
// C-Function: avfilter_process_command
func (f *FilterContext) ProcessCommand(cmd, arg string, fl int) (string, avutil.ReturnCode) {
	const buffsize = 128
	var buf [buffsize]byte
	err := avutil.NewReturnCode(int(C.avfilter_process_command((*C.struct_AVFilterContext)(f), C.CString(cmd), C.CString(arg), (*C.char)(unsafe.Pointer(&buf[0])), C.int(buffsize), C.int(fl))))
	return string(buf[:]), err
}

// RegisterAll initializes the filter system.
//
// C-Function: avfilter_register_all
func RegisterAll() {
	C.avfilter_register_all()
}

// NewFilterContext initializes a new filter with the supplied parameters.
//
// C-Function: avfilter_init_str
func NewFilterContext(args string) (*FilterContext, avutil.ReturnCode) {
	var f *FilterContext
	err := avutil.NewReturnCode(int(C.avfilter_init_str((*C.struct_AVFilterContext)(f), C.CString(args))))
	return f, err
}

// NewFilterContextDict initializes a filter with the supplied dictionary of options.
//
// C-Function: avfilter_init_dict
func NewFilterContextDict(o *Dictionary) (*FilterContext, avutil.ReturnCode) {
	var f *FilterContext
	err := avutil.NewReturnCode(int(C.avfilter_init_dict((*C.struct_AVFilterContext)(f), (**C.struct_AVDictionary)(unsafe.Pointer(&o)))))
	return f, err
}

// Free frees a filter context.
//
// C-Function: avfilter_free
func (ctx *FilterContext) Free() {
	C.avfilter_free((*C.struct_AVFilterContext)(ctx))
}

// InsertFilter insterts a filter in the middle of an existing link.
//
// C-Function: avfilter_insert_filter
func (l *FilterLink) InsertFilter(f *FilterContext, fsi, fdi uint) avutil.ReturnCode {
	return avutil.NewReturnCode(int(C.avfilter_insert_filter((*C.struct_AVFilterLink)(l), (*C.struct_AVFilterContext)(f), C.uint(fsi), C.uint(fdi))))
}

// GetClass returns the Class.
//
// C-Function: avfilter_get_class
func GetClass() *Class {
	return (*Class)(C.avfilter_get_class())
}

// NewFilterInOut allocates a single FilterInOut entry.
//
// C-Function: avfilter_inout_alloc
func NewFilterInOut() *FilterInOut {
	return (*FilterInOut)(C.avfilter_inout_alloc())
}

// Free frees the supplied list of FilterInOut and set *inout to NULL.
//
// C-Function: avfilter_inout_free
func (i *FilterInOut) Free() {
	var tmp [2]*FilterInOut
	tmp[0] = i
	tmp[1] = nil
	C.avfilter_inout_free((**C.struct_AVFilterInOut)(unsafe.Pointer(&tmp[0])))
}

// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

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

//Return the LIBAvFILTER_VERSION_INT constant.
func AvfilterVersion() uint {
	return uint(C.avfilter_version())
}

//Return the libavfilter build-time configuration.
func AvfilterConfiguration() string {
	return C.GoString(C.avfilter_configuration())
}

//Return the libavfilter license.
func AvfilterLicense() string {
	return C.GoString(C.avfilter_license())
}

//Get the number of elements in a NULL-terminated array of FilterPads (e.g.
func AvfilterFilterPadCount(p *FilterPad) int {
	return int(C.avfilter_pad_count((*C.struct_AVFilterPad)(p)))
}

//Get the name of an FilterPad.
func AvfilterFilterPadGetName(p *FilterPad, pi int) string {
	return C.GoString(C.avfilter_pad_get_name((*C.struct_AVFilterPad)(p), C.int(pi)))
}

//Get the type of an FilterPad.
func AvfilterFilterPadGetType(p *FilterPad, pi int) MediaType {
	return (MediaType)(C.avfilter_pad_get_type((*C.struct_AVFilterPad)(p), C.int(pi)))
}

//FilterLink two filters together.
func AvfilterFilterLink(s *FilterContext, sp uint, d *FilterContext, dp uint) int {
	return int(C.avfilter_link((*C.struct_AVFilterContext)(s), C.uint(sp), (*C.struct_AVFilterContext)(d), C.uint(dp)))
}

//Free the link in *link, and set its pointer to NULL.
func AvfilterFilterLinkFree(l **FilterLink) {
	C.avfilter_link_free((**C.struct_AVFilterLink)(unsafe.Pointer(l)))
}

//Get the number of channels of a link.
func AvfilterFilterLinkGetChannels(l *FilterLink) int {
	return int(C.avfilter_link_get_channels((*C.struct_AVFilterLink)(l)))
}

//Set the closed field of a link.
func AvfilterFilterLinkSetClosed(l *FilterLink, c int) {
	C.avfilter_link_set_closed((*C.struct_AVFilterLink)(l), C.int(c))
}

//Negotiate the media format, dimensions, etc of all inputs to a filter.
func AvfilterConfigFilterLinks(f *FilterContext) int {
	return int(C.avfilter_config_links((*C.struct_AVFilterContext)(f)))
}

//Make the filter instance process a command.
func AvfilterProcessCommand(f *FilterContext, cmd, arg, res string, l, fl int) int {
	return int(C.avfilter_process_command((*C.struct_AVFilterContext)(f), C.CString(cmd), C.CString(arg), C.CString(res), C.int(l), C.int(fl)))
}

//Initialize the filter system.
func AvfilterRegisterAll() {
	C.avfilter_register_all()
}

//Initialize a filter with the supplied parameters.
func (ctx *FilterContext) AvfilterInitStr(args string) int {
	return int(C.avfilter_init_str((*C.struct_AVFilterContext)(ctx), C.CString(args)))
}

//Initialize a filter with the supplied dictionary of options.
func (ctx *FilterContext) AvfilterInitDict(o **Dictionary) int {
	return int(C.avfilter_init_dict((*C.struct_AVFilterContext)(ctx), (**C.struct_AVDictionary)(unsafe.Pointer(o))))
}

//Free a filter context.
func (ctx *FilterContext) AvfilterFree() {
	C.avfilter_free((*C.struct_AVFilterContext)(ctx))
}

//Insert a filter in the middle of an existing link.
func AvfilterInsertFilter(l *FilterLink, f *FilterContext, fsi, fdi uint) int {
	return int(C.avfilter_insert_filter((*C.struct_AVFilterLink)(l), (*C.struct_AVFilterContext)(f), C.uint(fsi), C.uint(fdi)))
}

//avfilter_get_class
func AvfilterGetClass() *Class {
	return (*Class)(C.avfilter_get_class())
}

//Allocate a single FilterInOut entry.
func AvfilterInoutAlloc() *FilterInOut {
	return (*FilterInOut)(C.avfilter_inout_alloc())
}

//Free the supplied list of FilterInOut and set *inout to NULL.
func AvfilterInoutFree(i *FilterInOut) {
	C.avfilter_inout_free((**C.struct_AVFilterInOut)(unsafe.Pointer(i)))
}

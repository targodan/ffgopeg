// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)
// Corbatto (luca@corbatto.de)

package avfilter

/*
	#cgo pkg-config: libavfilter
	#include <libavfilter/avfilter.h>
*/
import "C"
import (
	"unsafe"

	"github.com/colek42/ffgopeg/avutil"
)

// NewFilterGraph allocates a filter graph.
//
// C-Function: avfilter_graph_alloc
func NewFilterGraph() *FilterGraph {
	return (*FilterGraph)(C.avfilter_graph_alloc())
}

// NewFilterContext creates a new filter instance in a filter graph.
//
// C-Function: avfilter_graph_alloc_filter
func (g *FilterGraph) NewFilterContext(f *Filter, name string) *FilterContext {
	return (*FilterContext)(C.avfilter_graph_alloc_filter((*C.struct_AVFilterGraph)(g), (*C.struct_AVFilter)(f), C.CString(name)))
}

// Filter retursn a filter instance identified by instance name from graph.
//
// C-Function: avfilter_graph_get_filter
func (g *FilterGraph) Filter(n string) *FilterContext {
	return (*FilterContext)(C.avfilter_graph_get_filter((*C.struct_AVFilterGraph)(g), C.CString(n)))
}

// SetAutoConvert enables or disables automatic format conversion inside the graph.
// flags can be any of the AVFILTER_AUTO_CONVERT_* constants.
//
// C-Function: avfilter_graph_set_auto_convert
func (g *FilterGraph) SetAutoConvert(flags uint) {
	C.avfilter_graph_set_auto_convert((*C.struct_AVFilterGraph)(g), C.uint(flags))
}

// GraphConfig checks validity and configure all the links and formats in the graph.
//
// C-Function: avfilter_graph_config
func (g *FilterGraph) GraphConfig() (interface{}, avutil.ReturnCode) {
	var logCtx interface{}
	err := avutil.NewReturnCode(int(C.avfilter_graph_config((*C.struct_AVFilterGraph)(g), unsafe.Pointer(&logCtx))))
	return logCtx, err
}

// Free frees a graph, destroy its links, and set *graph to NULL.
//
// C-Function: avfilter_graph_free
func (g *FilterGraph) Free() {
	C.avfilter_graph_free((**C.struct_AVFilterGraph)(unsafe.Pointer(&g)))
}

// Parse adds a graph described by a string to a graph.
//
// C-Function: avfilter_graph_parse
func (g *FilterGraph) Parse(f string, i, o *FilterInOut, l int) avutil.ReturnCode {
	return avutil.NewReturnCode(int(C.avfilter_graph_parse((*C.struct_AVFilterGraph)(g), C.CString(f), (*C.struct_AVFilterInOut)(i), (*C.struct_AVFilterInOut)(o), unsafe.Pointer(&l))))
}

// ParsePtr adds a graph described by a string to a graph.
//
// C-Function: avfilter_graph_parse_ptr
func (g *FilterGraph) ParsePtr(f string, i, o **FilterInOut, l int) avutil.ReturnCode {
	return avutil.NewReturnCode(int(C.avfilter_graph_parse_ptr((*C.struct_AVFilterGraph)(g), C.CString(f), (**C.struct_AVFilterInOut)(unsafe.Pointer(i)), (**C.struct_AVFilterInOut)(unsafe.Pointer(o)), unsafe.Pointer(&l))))
}

// Parse2 adds a graph described by a string to a graph.
//
// C-Function: avfilter_graph_parse2
func (g *FilterGraph) Parse2(f string, i, o **FilterInOut) avutil.ReturnCode {
	return avutil.NewReturnCode(int(C.avfilter_graph_parse2((*C.struct_AVFilterGraph)(g), C.CString(f), (**C.struct_AVFilterInOut)(unsafe.Pointer(i)), (**C.struct_AVFilterInOut)(unsafe.Pointer(o)))))
}

// SendCommand sends a command to one or more filter instances.
//
// C-Function: avfilter_graph_send_command
func (g *FilterGraph) SendCommand(t, cmd, arg, res string, resl, f int) avutil.ReturnCode {
	return avutil.NewReturnCode(int(C.avfilter_graph_send_command((*C.struct_AVFilterGraph)(g), C.CString(t), C.CString(cmd), C.CString(arg), C.CString(res), C.int(resl), C.int(f))))
}

// QueueCommand queues a command for one or more filter instances.
//
// C-Function: avfilter_graph_queue_command
func (g *FilterGraph) QueueCommand(t, cmd, arg string, f int, ts C.double) avutil.ReturnCode {
	return avutil.NewReturnCode(int(C.avfilter_graph_queue_command((*C.struct_AVFilterGraph)(g), C.CString(t), C.CString(cmd), C.CString(arg), C.int(f), ts)))
}

// Dump dumps a graph into a human-readable string representation.
//
// C-Function: avfilter_graph_dump
func (g *FilterGraph) Dump(options string) string {
	return C.GoString(C.avfilter_graph_dump((*C.struct_AVFilterGraph)(g), C.CString(options)))
}

// RequestOldestlink requests a frame on the oldest sink
//
// C-Function: avfilter_graph_request_oldest
func (g *FilterGraph) RequestOldestlink() avutil.ReturnCode {
	return avutil.NewReturnCode(int(C.avfilter_graph_request_oldest((*C.struct_AVFilterGraph)(g))))
}

// CreateFilter creates and add a filter instance into an existing graph.
//
// C-Function: avfilter_graph_create_filter
func (g *FilterGraph) CreateFilter(cx **FilterContext, f *Filter, n, a string, o int) avutil.ReturnCode {
	return avutil.NewReturnCode(int(C.avfilter_graph_create_filter((**C.struct_AVFilterContext)(unsafe.Pointer(cx)), (*C.struct_AVFilter)(f), C.CString(n), C.CString(a), unsafe.Pointer(&o), (*C.struct_AVFilterGraph)(g))))
}

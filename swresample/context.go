// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package swresample

/*
	#cgo pkg-config: libswresample
	#include <libswresample/swresample.h>
*/
import "C"
import (
	"unsafe"
)

// Init initializes the context after user parameters have been set.
//
// C-Function: swr_init
func (s *Context) Init() int {
	return int(C.swr_init((*C.struct_SwrContext)(s)))
}

// IsInitialized checks whether an swr context has been initialized or not.
//
// C-Function: swr_is_initialized
func (s *Context) IsInitialized() bool {
	return int(C.swr_is_initialized((*C.struct_SwrContext)(s))) != 0
}

// SetOpts allocates Context if needed and set/reset common parameters.
//
// C-Function: swr_alloc_set_opts
func (s *Context) SetOpts(ocl int64, osf SampleFormat, osr int, icl int64, isf SampleFormat, isr, lo, lc int) *Context {
	return (*Context)(C.swr_alloc_set_opts((*C.struct_SwrContext)(s), C.int64_t(ocl), (C.enum_AVSampleFormat)(osf), C.int(osr), C.int64_t(icl), (C.enum_AVSampleFormat)(isf), C.int(isr), C.int(lo), unsafe.Pointer(&lc)))
}

// NewContextSetOpts allocates Context if needed and set/reset common parameters.
//
// C-Function: swr_alloc_set_opts
func NewContextSetOpts(ocl int64, osf SampleFormat, osr int, icl int64, isf SampleFormat, isr, lo, lc int) *Context {
	return (*Context)(C.swr_alloc_set_opts((*C.struct_SwrContext)(nil), C.int64_t(ocl), (C.enum_AVSampleFormat)(osf), C.int(osr), C.int64_t(icl), (C.enum_AVSampleFormat)(isf), C.int(isr), C.int(lo), unsafe.Pointer(&lc)))
}

// Free frees Context destructor functions. Free the given Context and set the pointer to NULL.
//
// C-Function: swr_free
func (s *Context) Free() {
	C.swr_free((**C.struct_SwrContext)(unsafe.Pointer(s)))
}

// Close closes the context so that swr_is_initialized() returns 0.
//
// C-Function: swr_close
func (s *Context) Close() {
	C.swr_close((*C.struct_SwrContext)(s))
}

// Convert core conversion functions. Convert audio
//
// C-Function: swr_convert
func (s *Context) Convert(out **uint8, oc int, in **uint8, ic int) int {
	return int(C.swr_convert((*C.struct_SwrContext)(s), (**C.uint8_t)(unsafe.Pointer(out)), C.int(oc), (**C.uint8_t)(unsafe.Pointer(in)), C.int(ic)))
}

// NextPts converts the next timestamp from input to output timestamps are in 1/(in_sample_rate * out_sample_rate) units.
//
// C-Function: swr_next_pts
func (s *Context) NextPts(pts int64) int64 {
	return int64(C.swr_next_pts((*C.struct_SwrContext)(s), C.int64_t(pts)))
}

// SetCompensation activates resampling compensation ("soft" compensation).
//
// C-Function: swr_set_compensation
func (s *Context) SetCompensation(sd, cd int) int {
	return int(C.swr_set_compensation((*C.struct_SwrContext)(s), C.int(sd), C.int(cd)))
}

// SetChannelMapping sets a customized input channel mapping.
//
// C-Function: swr_set_channel_mapping
func (s *Context) SetChannelMapping(cm *int) int {
	return int(C.swr_set_channel_mapping((*C.struct_SwrContext)(s), (*C.int)(unsafe.Pointer(cm))))
}

// SetMatrix sets a customized remix matrix.
//
// C-Function: swr_set_matrix
func (s *Context) SetMatrix(m *int, t int) int {
	return int(C.swr_set_matrix((*C.struct_SwrContext)(s), (*C.double)(unsafe.Pointer(m)), C.int(t)))
}

// DropOutput drops the specified number of output samples.
//
// C-Function: swr_drop_output
func (s *Context) DropOutput(c int) int {
	return int(C.swr_drop_output((*C.struct_SwrContext)(s), C.int(c)))
}

// InjectSilence injects the specified number of silence samples.
//
// C-Function: swr_inject_silence
func (s *Context) InjectSilence(c int) int {
	return int(C.swr_inject_silence((*C.struct_SwrContext)(s), C.int(c)))
}

// GetDelay gets the delay the next input sample will experience relative to the next output sample.
//
// C-Function: swr_get_delay
func (s *Context) GetDelay(b int64) int64 {
	return int64(C.swr_get_delay((*C.struct_SwrContext)(s), C.int64_t(b)))
}

// ConvertFrame converts the samples in the input Frame and write them to the output Frame.
//
// C-Function: swr_convert_frame
func (s *Context) ConvertFrame(o, i *Frame) int {
	return int(C.swr_convert_frame((*C.struct_SwrContext)(s), (*C.struct_AVFrame)(o), (*C.struct_AVFrame)(i)))
}

// ConfigFrame configures or reconfigures the Context using the information provided by the AvFrames.
//
// C-Function: swr_config_frame
func (s *Context) ConfigFrame(o, i *Frame) int {
	return int(C.swr_config_frame((*C.struct_SwrContext)(s), (*C.struct_AVFrame)(o), (*C.struct_AVFrame)(i)))
}

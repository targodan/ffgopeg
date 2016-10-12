// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)
// Corbatto (luca@corbatto.de)

package avutil

/*
	#cgo pkg-config: libavutil
	#include <libavutil/frame.h>
	#include <stdlib.h>

    int MACRO_NUM_DATA_POINTERS() {
        return AV_NUM_DATA_POINTERS;
    }

    unsigned ComparePointers(void* a, void* b) {
        return (unsigned)(a-b);
    }
*/
import "C"
import "unsafe"

type (
	Buffer            C.struct_AVBuffer
	BufferRef         C.struct_AVBufferRef
	BufferPool        C.struct_AVBufferPool
	Frame             C.struct_AVFrame
	FrameSideData     C.struct_AVFrameSideData
	FrameSideDataType C.enum_AVFrameSideDataType
)

func comparePointers(a, b unsafe.Pointer) uintptr {
	return uintptr(C.ComparePointers(a, b))
}

func MACRO_NUM_DATA_POINTERS() int {
	return int(C.MACRO_NUM_DATA_POINTERS())
}

// Metadatap returns metadatap.
//
// C-Function: avpriv_frame_get_metadatap
func (f *Frame) Metadatap() **Dictionary {
	return (**Dictionary)(unsafe.Pointer(C.avpriv_frame_get_metadatap((*C.struct_AVFrame)(unsafe.Pointer(f)))))
}

// SetQpTable sets the qp table.
//
// C-Function: av_frame_set_qp_table
func (f *Frame) SetQpTable(b *BufferRef, s, q int) int {
	return int(C.av_frame_set_qp_table((*C.struct_AVFrame)(unsafe.Pointer(f)), (*C.struct_AVBufferRef)(unsafe.Pointer(b)), C.int(s), C.int(q)))
}

// QpTable returns the qp table.
//
// C-Function: av_frame_get_qp_table
func (f *Frame) QpTable(stride, t *int) int8 {
	return int8(*C.av_frame_get_qp_table((*C.struct_AVFrame)(unsafe.Pointer(f)), (*C.int)(unsafe.Pointer(stride)), (*C.int)(unsafe.Pointer(t))))
}

// NewFrame allocates a Frame and set its fields to default values.
//
// C-Function: av_frame_alloc
func NewFrame() *Frame {
	return (*Frame)(unsafe.Pointer(C.av_frame_alloc()))
}

// Free frees the frame and any dynamically allocated objects in it, e.g.
//
// C-Function: av_frame_free
func (f *Frame) Free() {
	C.av_frame_free((**C.struct_AVFrame)(unsafe.Pointer(&f)))
}

// NewBuffer allocates new buffer(s) for audio or video data.
//
// C-Function: av_frame_get_buffer
func (f *Frame) NewBuffer(a int) int {
	return int(C.av_frame_get_buffer((*C.struct_AVFrame)(unsafe.Pointer(f)), C.int(a)))
}

// FrameRef sets up a new reference to the data described by a given frame.
//
// C-Function: av_frame_ref
func FrameRef(d, s *Frame) int {
	return int(C.av_frame_ref((*C.struct_AVFrame)(unsafe.Pointer(d)), (*C.struct_AVFrame)(unsafe.Pointer(s))))
}

// Clone creates a new frame that references the same data as src.
//
// C-Function: av_frame_clone
func Clone(f *Frame) *Frame {
	return (*Frame)(C.av_frame_clone((*C.struct_AVFrame)(unsafe.Pointer(f))))
}

// Unref unreferences all the buffers referenced by frame and reset the frame fields.
//
// C-Function: av_frame_unref
func (f *Frame) Unref() {
	cf := (*C.struct_AVFrame)(unsafe.Pointer(f))
	C.av_frame_unref(cf)
}

// FrameMoveRef moves everythnig contained in src to dst and reset src.
//
// C-Function: av_frame_move_ref
func FrameMoveRef(d, s *Frame) {
	C.av_frame_move_ref((*C.struct_AVFrame)(unsafe.Pointer(d)), (*C.struct_AVFrame)(unsafe.Pointer(s)))
}

// IsWritable checks if the frame data is writable.
//
// C-Function: av_frame_is_writable
func (f *Frame) IsWritable() bool {
	return int(C.av_frame_is_writable((*C.struct_AVFrame)(unsafe.Pointer(f)))) != 0
}

// MakeWritable ensures that the frame data is writable, avoiding data copy if possible.
//
// C-Function: av_frame_make_writable
func (f *Frame) MakeWritable() ReturnCode {
	return NewReturnCode(int(C.av_frame_make_writable((*C.struct_AVFrame)(unsafe.Pointer(f)))))
}

// CopyProps copies only "metadata" fields from src to dst.
//
// C-Function: av_frame_copy_props
func CopyProps(d, s *Frame) int {
	return int(C.av_frame_copy_props((*C.struct_AVFrame)(unsafe.Pointer(d)), (*C.struct_AVFrame)(unsafe.Pointer(s))))
}

// PlaneBuffer gets the buffer reference a given data plane is stored in.
//
// C-Function: av_frame_get_plane_buffer
func (f *Frame) PlaneBuffer(p int) *BufferRef {
	return (*BufferRef)(C.av_frame_get_plane_buffer((*C.struct_AVFrame)(unsafe.Pointer(f)), C.int(p)))
}

// NewSideData adds a new side data to a frame.
//
// C-Function: av_frame_new_side_data
func (f *Frame) NewSideData(d FrameSideDataType, s int) *FrameSideData {
	return (*FrameSideData)(C.av_frame_new_side_data((*C.struct_AVFrame)(unsafe.Pointer(f)), (C.enum_AVFrameSideDataType)(d), C.int(s)))
}

// SideData returns the frames sidedata.
//
// C-Function: av_frame_get_side_data
func (f *Frame) SideData(t FrameSideDataType) *FrameSideData {
	return (*FrameSideData)(C.av_frame_get_side_data((*C.struct_AVFrame)(unsafe.Pointer(f)), (C.enum_AVFrameSideDataType)(t)))
}

// Data returns the frames data.
//
// C-Variable: AVFrame::data
func (f *Frame) Data(i int, len int) []byte {
	return (*[1 << 30]byte)(unsafe.Pointer(f.data[i]))[:len:len]
}

// ExtendedData returns the frames extended_data.
//
// C-Variable: AVFrame::data
func (f *Frame) ExtendedData(i int, len int) []byte {
	// From the documentation
	// > For video, this should simply point to data[].
	arr := (*[1 << 30]*C.uint8_t)(unsafe.Pointer(f.extended_data))
	return (*[1 << 30]byte)(unsafe.Pointer(arr[i]))[:len:len]
}

// Linesize returns the frames linesize.
//
// C-Variable: AVFrame::linesize
func (f *Frame) Linesize(i int) int {
	return int(f.linesize[i])
}

// NbSamples returns the number of samples.
//
// C-Variable: AVFrame::nb_frames
func (f *Frame) NbSamples() int {
	return int(f.nb_samples)
}

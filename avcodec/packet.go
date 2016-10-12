// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)
// Corbatto (luca@corbatto.de)

package avcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"
import (
	"unsafe"
)

// Init initializes optional fields of a packet with default values.
//
// C-Function: av_init_packet
func (p Packet) Init() {
	C.av_init_packet((*C.struct_AVPacket)(&p))
}

// NewPacket allocates the payload of a packet and initialize its fields with default values.
//
// C-Function: av_new_packet
func (p *Packet) NewPacket(s int) int {
	return int(C.av_new_packet((*C.struct_AVPacket)(p), C.int(s)))
}

// ShrinkPacket reduces packet size, correctly zeroing padding.
//
// C-Function: av_shrink_packet
func (p *Packet) ShrinkPacket(s int) {
	C.av_shrink_packet((*C.struct_AVPacket)(p), C.int(s))
}

// GrowPacket increases packet size, correctly zeroing padding.
//
// C-Function: av_grow_packet
func (p *Packet) GrowPacket(s int) int {
	return int(C.av_grow_packet((*C.struct_AVPacket)(p), C.int(s)))
}

// PacketFromData initializes a reference-counted packet from av_malloc()ed data.
//
// C-Function: av_packet_from_data
func (p *Packet) PacketFromData(d []uint8, s int) int {
	return int(C.av_packet_from_data((*C.struct_AVPacket)(p), (*C.uint8_t)(&d[0]), C.int(s)))

}

// CopyPacket copies a packet, including contents.
//
// C-Function:
func (p *Packet) CopyPacket(src *Packet) int {
	return int(C.av_copy_packet((*C.struct_AVPacket)(p), (*C.struct_AVPacket)(src)))
}

// CopyPacketSideData copies packet side data.
//
// C-Function: av_copy_packet_side_data
func (p *Packet) CopyPacketSideData(src *Packet) int {
	return int(C.av_copy_packet_side_data((*C.struct_AVPacket)(p), (*C.struct_AVPacket)(src)))

}

// Unref wipes a packet.
//
// C-Function: av_packet_unref
func (p *Packet) Unref() {
	C.av_packet_unref((*C.struct_AVPacket)(p))
}

// NewSideData allocates new information of a packet.
//
// C-Function: av_packet_new_side_data
func (p *Packet) NewSideData(t PacketSideDataType, size int) []uint8 {
	return (*[1 << 30]uint8)(unsafe.Pointer(C.av_packet_new_side_data((*C.struct_AVPacket)(p), (C.enum_AVPacketSideDataType)(t), C.int(size))))[:size:size]
}

// ShrinkSideData shrinks the already allocated side data buffer.
//
// C-Function: av_packet_shrink_side_data
func (p *Packet) ShrinkSideData(t PacketSideDataType, size int) int {
	return int(C.av_packet_shrink_side_data((*C.struct_AVPacket)(p), (C.enum_AVPacketSideDataType)(t), C.int(size)))
}

// SideDataOfType get side information from packet.
//
// C-Function: av_packet_get_side_data
func (p *Packet) SideDataOfType(t PacketSideDataType) []uint8 {
	var s int
	buf := (*[1 << 30]uint8)(unsafe.Pointer(C.av_packet_get_side_data((*C.struct_AVPacket)(p), (C.enum_AVPacketSideDataType)(t), (*C.int)(unsafe.Pointer(&s)))))
	return buf[:s:s]
}

// MergeSideData does something undocumented...
//
// C-Function: av_packet_merge_side_data
func (p *Packet) MergeSideData() int {
	return int(C.av_packet_merge_side_data((*C.struct_AVPacket)(p)))
}

// SplitSideData does something undocumented...
//
// C-Function: av_packet_split_side_data
func (p *Packet) SplitSideData() int {
	return int(C.av_packet_split_side_data((*C.struct_AVPacket)(p)))
}

// FreeSideData is a convenience function to free all the side data stored.
//
// C-Function: av_packet_free_side_data
func (p *Packet) FreeSideData() {
	C.av_packet_free_side_data((*C.struct_AVPacket)(p))
}

// Ref sets up a new reference to the data described by a given packet.
//
// C-Function: av_packet_ref
func (p *Packet) Ref(src *Packet) int {
	return int(C.av_packet_ref((*C.struct_AVPacket)(p), (*C.struct_AVPacket)(src)))
}

// MoveRef moves every field in src to dst and reset src.
//
// C-Function: av_packet_move_ref
func (p *Packet) MoveRef(src *Packet) {
	C.av_packet_move_ref((*C.struct_AVPacket)(p), (*C.struct_AVPacket)(src))
}

// CopyProps copies only "properties" fields from src to dst.
//
// C-Function: av_packet_copy_props
func (p *Packet) CopyProps(src *Packet) int {
	return int(C.av_packet_copy_props((*C.struct_AVPacket)(p), (*C.struct_AVPacket)(src)))
}

// RescaleTs converts valid timing fields (timestamps / durations) in a packet from one timebase to another.
//
// C-Function: av_packet_rescale_ts
func (p *Packet) RescaleTs(r, r2 Rational) {
	C.av_packet_rescale_ts((*C.struct_AVPacket)(p), (C.struct_AVRational)(r), (C.struct_AVRational)(r2))
}

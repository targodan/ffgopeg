// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)
// Corbatto (luca@corbatto.de)

package avcodec

// Buf returns a reference to the reference-counted buffer where the packet data is stored.
// May be nil, then the packet data is not reference-counted.
//
// C-Field: AVPacket::buf
func (p *Packet) Buf() *BufferRef {
	return (*BufferRef)(p.buf)
}

// Duration returns the duration of this packet in AVStream->time_base units, 0 if unknown.
//
// C-Field: AVPacket::duration
func (p *Packet) Duration() int64 {
	return int64(p.duration)
}

// Flags returns a combination of AV_PKT_FLAG values.
//
// C-Field: AVPacket::flags
func (p *Packet) Flags() int {
	return int(p.flags)
}

// SideDataElems returns something undefined...
//
// C-Field: AVPacket::side_data_elems
func (p *Packet) SideDataElems() int {
	return int(p.side_data_elems)
}

// Size returns something undocumented...
//
// C-Field: AVPacket::size
func (p *Packet) Size() int {
	return int(p.size)
}

// StreamIndex returns the index of the stream this packet originated from.
//
// C-Field: AVPacket::stream_index
func (p *Packet) StreamIndex() int {
	return int(p.stream_index)
}

// Dts returns the decompression timestamp in AVStream->time_base units; the time at which the packet is decompressed.
// Can be AV_NOPTS_VALUE if it is not stored in the file.
//
// C-Field: AVPacket::dts
func (p *Packet) Dts() int64 {
	return int64(p.dts)
}

// Pos returns the byte position in stream, -1 if unknown.
//
// C-Field: AVPacket::pos
func (p *Packet) Pos() int64 {
	return int64(p.pos)
}

// Pts returns the presentation timestamp in AVStream->time_base units; the time at which the decompressed packet will be presented to the user.
//
// Can be AV_NOPTS_VALUE if it is not stored in the file. pts MUST be larger or equal to dts as presentation cannot happen before decompression, unless one wants to view hex dumps. Some formats misuse the terms dts and pts/cts to mean something different. Such timestamps must be converted to true pts/dts before they are stored in AVPacket.
//
// C-Field: AVPacket::pts
func (p *Packet) Pts() int64 {
	return int64(p.pts)
}

// Data returns the data.
//
// C-Field: AVPacket::data
func (p *Packet) Data() *uint8 {
	return (*uint8)(p.data)
}

// SideData returns the side data.
//
// C-Field: AVPacket::
func (p *Packet) SideData() *PacketSideData {
	return (*PacketSideData)(p.side_data)
}

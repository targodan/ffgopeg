// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)
// Corbatto (luca@corbatto.de)

package avcodec

import "unsafe"

// Data returns the data.
//
// C-Field: AVPacketSideData::data
func (sd *PacketSideData) Data() []uint8 {
	return (*[1 << 30]uint8)(unsafe.Pointer(sd.data))[:sd.Size():sd.Size()]
}

// Size returns size of the data.
//
// C-Field: AVPacketSideData::size
func (sd *PacketSideData) Size() int {
	return int(sd.size)
}

// Type returns the type of the side data.
func (sd *PacketSideData) Type() PacketSideDataType {
	return PacketSideDataType(sd._type)
}

// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avcodec

func (c CodecParameters) CodecType() MediaType {
	return MediaType(c.codec_type)
}

func (c CodecParameters) CodecId() CodecId {
	return CodecId(c.codec_id)
}

func (c CodecParameters) CodecTag() uint32 {
	return uint32(c.codec_tag)
}

func (c CodecParameters) Format() int {
	return int(c.format)
}

func (c CodecParameters) BitRate() int64 {
	return int64(c.bit_rate)
}

func (c CodecParameters) BitsPerCodedSample() int {
	return int(c.bits_per_coded_sample)
}

func (c CodecParameters) BitsPerRawSample() int {
	return int(c.bits_per_raw_sample)
}

func (c CodecParameters) Profile() int {
	return int(c.profile)
}

func (c CodecParameters) Level() int {
	return int(c.level)
}

func (c CodecParameters) Width() int {
	return int(c.width)
}

func (c CodecParameters) Height() int {
	return int(c.height)
}

func (c CodecParameters) SampleAspectRatio() Rational {
	return Rational(c.sample_aspect_ratio)
}

func (c CodecParameters) FieldOrder() FieldOrder {
	return FieldOrder(c.field_order)
}

func (c CodecParameters) ColorRange() ColorRange {
	return ColorRange(c.color_range)
}

func (c CodecParameters) ColorPrimaries() ColorPrimaries {
	return ColorPrimaries(c.color_primaries)
}

func (c CodecParameters) ColorTrc() ColorTransferCharacteristic {
	return ColorTransferCharacteristic(c.color_trc)
}

func (c CodecParameters) ColorSpace() ColorSpace {
	return ColorSpace(c.color_space)
}

func (c CodecParameters) ChromaLocation() ChromaLocation {
	return ChromaLocation(c.chroma_location)
}

func (c CodecParameters) VideoDelay() int {
	return int(c.video_delay)
}

func (c CodecParameters) ChannelLayout() uint64 {
	return uint64(c.channel_layout)
}

func (c CodecParameters) Channels() int {
	return int(c.channels)
}

func (c CodecParameters) SampleRate() int {
	return int(c.sample_rate)
}

func (c CodecParameters) BlockAlign() int {
	return int(c.block_align)
}

func (c CodecParameters) FrameSize() int {
	return int(c.frame_size)
}

func (c CodecParameters) InitialPadding() int {
	return int(c.initial_padding)
}

func (c CodecParameters) TrailingPadding() int {
	return int(c.trailing_padding)
}

func (c CodecParameters) SeekPreroll() int {
	return int(c.seek_preroll)
}

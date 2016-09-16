// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avcodec

func (c AvCodecParameters) CodecType() MediaType {
	return MediaType(c.codec_type)
}

func (c AvCodecParameters) CodecId() CodecId {
	return CodecId(c.codec_id)
}

func (c AvCodecParameters) CodecTag() uint32 {
	return uint32(c.codec_tag)
}

func (c AvCodecParameters) Format() int {
	return int(c.format)
}

func (c AvCodecParameters) BitRate() int64 {
	return int64(c.bit_rate)
}

func (c AvCodecParameters) BitsPerCodedSample() int {
	return int(c.bits_per_coded_sample)
}

func (c AvCodecParameters) BitsPerRawSample() int {
	return int(c.bits_per_raw_sample)
}

func (c AvCodecParameters) Profile() int {
	return int(c.profile)
}

func (c AvCodecParameters) Level() int {
	return int(c.level)
}

func (c AvCodecParameters) Width() int {
	return int(c.width)
}

func (c AvCodecParameters) Height() int {
	return int(c.height)
}

func (c AvCodecParameters) SampleAspectRatio() Rational {
	return Rational(c.sample_aspect_ratio)
}

func (c AvCodecParameters) FieldOrder() AvFieldOrder {
	return AvFieldOrder(c.field_order)
}

func (c AvCodecParameters) ColorRange() AvColorRange {
	return AvColorRange(c.color_range)
}

func (c AvCodecParameters) ColorPrimaries() AvColorPrimaries {
	return AvColorPrimaries(c.color_primaries)
}

func (c AvCodecParameters) ColorTrc() AvColorTransferCharacteristic {
	return AvColorTransferCharacteristic(c.color_trc)
}

func (c AvCodecParameters) ColorSpace() AvColorSpace {
	return AvColorSpace(c.color_space)
}

func (c AvCodecParameters) ChromaLocation() AvChromaLocation {
	return AvChromaLocation(c.chroma_location)
}

func (c AvCodecParameters) VideoDelay() int {
	return int(c.video_delay)
}

func (c AvCodecParameters) ChannelLayout() uint64 {
	return uint64(c.channel_layout)
}

func (c AvCodecParameters) Channels() int {
	return int(c.channels)
}

func (c AvCodecParameters) SampleRate() int {
	return int(c.sample_rate)
}

func (c AvCodecParameters) BlockAlign() int {
	return int(c.block_align)
}

func (c AvCodecParameters) FrameSize() int {
	return int(c.frame_size)
}

func (c AvCodecParameters) InitialPadding() int {
	return int(c.initial_padding)
}

func (c AvCodecParameters) TrailingPadding() int {
	return int(c.trailing_padding)
}

func (c AvCodecParameters) SeekPreroll() int {
	return int(c.seek_preroll)
}

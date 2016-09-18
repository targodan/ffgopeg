// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avcodec

// CodecType returns the general type of the encoded data.
//
// C-Attribute: codec_type
func (c CodecParameters) CodecType() MediaType {
	return MediaType(c.codec_type)
}

// CodecID returns the id of the codec.
//
// C-Attribute: codec_id
func (c CodecParameters) CodecID() CodecId {
	return CodecId(c.codec_id)
}

// CodecTag returns additional information about the codec (corrensponds to the AVI FOURCC).
//
// C-Attribute: codec_tag
func (c CodecParameters) CodecTag() uint32 {
	return uint32(c.codec_tag)
}

// Format returns the pixel/sample format of video/audio data.
//
// C-Attribute: format
func (c CodecParameters) Format() int {
	return int(c.format)
}

// BitRate returns the average bitrate of the encoded data (in bits per second).
//
// C-Attribute: bit_rate
func (c CodecParameters) BitRate() int64 {
	return int64(c.bit_rate)
}

// BitsPerCodedSample returns the number of bits per sample in the coded words.
//
// This is basically the bitrate per sample. It is mandatory for a bunch of formats to actually decode them. It's the number of bits for one sample.
//
// This could be for example 4 for ADPCM For PCM formats this matches bits_per_raw_sample. Can be 0
//
// C-Attribute: bits_per_coded_sample
func (c CodecParameters) BitsPerCodedSample() int {
	return int(c.bits_per_coded_sample)
}

// BitsPerRawSample returns the number of valid bits in each output sample.
//
// If the sample format has more bits, the least significant bits are additional padding bits, which are always 0. Use right shifts to reduce the sample to its actual size. For example, audio formats with 24 bit samples will have bits_per_raw_sample set to 24, and format set to AV_SAMPLE_FMT_S32. To get the original sample use "(int32_t)sample >> 8"."
//
// For ADPCM this might be 12 or 16 or similar Can be 0
//
// C-Attribute: bits_per_raw_sample
func (c CodecParameters) BitsPerRawSample() int {
	return int(c.bits_per_raw_sample)
}

// Profile returns the codec-specific bitstream restrictions that the stream conforms to.
//
// C-Attribute: profile
func (c CodecParameters) Profile() int {
	return int(c.profile)
}

// Level returns something undocumented...
//
// C-Attribute: level
func (c CodecParameters) Level() int {
	return int(c.level)
}

// Width returns the width of the video frame in pixels.
//
// C-Attribute: width
func (c CodecParameters) Width() int {
	return int(c.width)
}

// Height returns the height of the video frame in pixels.
//
// C-Attribute: height
func (c CodecParameters) Height() int {
	return int(c.height)
}

// SampleAspectRatio returns the aspect ratio (width / height) which a single pixel should have when displayed.
//
// When the aspect ratio is unknown / undefined, the numerator should be set to 0 (the denominator may have any value).
//
// C-Attribute: sample_aspect_ratio
func (c CodecParameters) SampleAspectRatio() Rational {
	return Rational(c.sample_aspect_ratio)
}

// FieldOrder returns the order of the fields in interlaced video.
//
// C-Attribute: field_order
func (c CodecParameters) FieldOrder() FieldOrder {
	return FieldOrder(c.field_order)
}

// ColorRange returns additional colorspace characteristics.
//
// C-Attribute: color_range
func (c CodecParameters) ColorRange() ColorRange {
	return ColorRange(c.color_range)
}

// ColorPrimaries returns something undocumented...
//
// C-Attribute: color_primaries
func (c CodecParameters) ColorPrimaries() ColorPrimaries {
	return ColorPrimaries(c.color_primaries)
}

// ColorTrc returns something undocumented...
//
// C-Attribute: color_trc
func (c CodecParameters) ColorTrc() ColorTransferCharacteristic {
	return ColorTransferCharacteristic(c.color_trc)
}

// ColorSpace returns something undocumented...
//
// C-Attribute: color_space
func (c CodecParameters) ColorSpace() ColorSpace {
	return ColorSpace(c.color_space)
}

// ChromaLocation returns something undocumented...
//
// C-Attribute: chroma_location
func (c CodecParameters) ChromaLocation() ChromaLocation {
	return ChromaLocation(c.chroma_location)
}

// VideoDelay returns the number of delayed frames.
//
// C-Attribute: video_delay
func (c CodecParameters) VideoDelay() int {
	return int(c.video_delay)
}

// ChannelLayout returns the channel layout bitmask.
// May be 0 if the channel layout is unknown or unspecified, otherwise the number of bits set bmust be equal to the channels field.
//
// C-Attribute: channel_layout
func (c CodecParameters) ChannelLayout() uint64 {
	return uint64(c.channel_layout)
}

// Channels return the number of audio channels.
//
// C-Attribute: channels
func (c CodecParameters) Channels() int {
	return int(c.channels)
}

// SampleRate returns the number of audio samples per second.
//
// C-Attribute: sample_rate
func (c CodecParameters) SampleRate() int {
	return int(c.sample_rate)
}

// BlockAlign returns the number of bytes per coded audio frame, required by some formats.
//
// Corresponds to nBlockAlign in WAVEFORMATEX.
//
// C-Attribute: block_align
func (c CodecParameters) BlockAlign() int {
	return int(c.block_align)
}

// FrameSize returns the audio frame size, if known.
// Required by some formats to be static.
//
// C-Attribute: frame_size
func (c CodecParameters) FrameSize() int {
	return int(c.frame_size)
}

// InitialPadding returns the amount of padding (in samples) inserted by the encoder at the beginning of the audio. I.e. this number of leading decoded samples must be discarded by the caller to get the original audio without leading padding.
//
// C-Attribute: initial_padding
func (c CodecParameters) InitialPadding() int {
	return int(c.initial_padding)
}

// TrailingPadding returns the amount of padding (in samples) appended by the encoder to the end of the audio. I.e. this number of decoded samples must be discarded by the caller from the end of the stream to get the original audio without any trailing padding.
//
// C-Attribute: trailing_padding
func (c CodecParameters) TrailingPadding() int {
	return int(c.trailing_padding)
}

// SeekPreroll returns the number of sample to skip after a discontinuity.
//
// C-Attribute: seek_preroll
func (c CodecParameters) SeekPreroll() int {
	return int(c.seek_preroll)
}

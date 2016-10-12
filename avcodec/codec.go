// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)
// Corbatto (luca@corbatto.de)

package avcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"
import (
	"unsafe"

	"github.com/targodan/goav/avutil"
)

// NewCodecContext allocates a Context and set its fields to default values.
//
// C-Function: avcodec_alloc_context3
func NewCodecContext(c *Codec) *CodecContext {
	return (*CodecContext)(C.avcodec_alloc_context3((*C.struct_AVCodec)(c)))
}

// FindDecoder finds a registered decoder with a matching codec ID.
//
// C-Function: avcodec_find_decoder
func FindDecoder(id CodecId) *Codec {
	return (*Codec)(C.avcodec_find_decoder((C.enum_AVCodecID)(id)))
}

// FindDecoderByName finds a registered decoder with the specified name.
//
// C-Function: avcodec_find_decoder_by_name
func FindDecoderByName(n string) *Codec {
	return (*Codec)(C.avcodec_find_decoder_by_name(C.CString(n)))
}

// FindEncoder finds a registered encoder with a matching codec ID.
//
// C-Function: avcodec_find_encoder
func FindEncoder(id CodecId) *Codec {
	return (*Codec)(C.avcodec_find_encoder((C.enum_AVCodecID)(id)))
}

// FindEncoderByName finds a registered encoder with the specified name.
//
// C-Function: avcodec_find_encoder_by_name
func FindEncoderByName(c string) *Codec {
	return (*Codec)(C.avcodec_find_encoder_by_name(C.CString(c)))
}

// MaxLowres returns the maximum lowres supported by the decoder.
//
// C-Function: av_codec_get_max_lowres
func (c *Codec) MaxLowres() int {
	return int(C.av_codec_get_max_lowres((*C.struct_AVCodec)(c)))
}

// ProfileName returns a name for the specified profile, if available.
//
// C-Function: av_get_profile_name
func (c *Codec) ProfileName(p int) string {
	return C.GoString(C.av_get_profile_name((*C.struct_AVCodec)(c), C.int(p)))
}

// IsEncoder returns true if the codec is an encoder.
//
// C-Function: av_codec_is_encoder
func (c *Codec) IsEncoder() bool {
	return int(C.av_codec_is_encoder((*C.struct_AVCodec)(c))) != 0
}

// IsDecoder returns true if the codec is an decoder.
//
// C-Function: av_codec_is_decoder
func (c *Codec) IsDecoder() bool {
	return int(C.av_codec_is_decoder((*C.struct_AVCodec)(c))) != 0
}

// Name returns the name of the codec.
//
// C-Variable: AVCodec::name
func (c *Codec) Name() string {
	return C.GoString(c.name)
}

// LongName returns the long, descriptive name of the codec.
//
// C-Variable: AVCodec::long_name
func (c *Codec) LongName() string {
	return C.GoString(c.long_name)
}

// SampleFmts returns the suppoted sample formats.
//
// C-Variable: AVCodec::sample_fmts
func (c *Codec) SampleFmts() []avutil.SampleFormat {
	arr := (*[1 << 30]avutil.SampleFormat)(unsafe.Pointer(c.sample_fmts))
	var len int
	for len = 0; arr[len] != -1; len++ {
	}
	return arr[:len:len]
}

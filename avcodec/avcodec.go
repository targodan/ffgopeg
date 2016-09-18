// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)
// Corbatto (luca@corbatto.de)

// Package avcodec contains the codecs (decoders and encoders) provided by the libavcodec library
package avcodec

//#cgo pkg-config: libavformat libavcodec libavutil libswresample
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
//#include <libavutil/avutil.h>
import "C"
import (
	"unsafe"

	"github.com/targodan/goav/avutil"
)

type (
	Codec                       C.struct_AVCodec
	CodecContext                C.struct_AVCodecContext
	CodecDescriptor             C.struct_AVCodecDescriptor
	CodecParser                 C.struct_AVCodecParser
	CodecParserContext          C.struct_AVCodecParserContext
	Dictionary                  C.struct_AVDictionary
	Frame                       C.struct_AVFrame
	MediaType                   C.enum_AVMediaType
	Packet                      C.struct_AVPacket
	BitStreamFilter             C.struct_AVBitStreamFilter
	BitStreamFilterContext      C.struct_AVBitStreamFilterContext
	Rational                    C.struct_AVRational
	Class                       C.struct_AVClass
	HWAccel                     C.struct_AVHWAccel
	PacketSideData              C.struct_AVPacketSideData
	PanScan                     C.struct_AVPanScan
	Profile                     C.struct_AVProfile
	Subtitle                    C.struct_AVSubtitle
	SubtitleRect                C.struct_AVSubtitleRect
	RcOverride                  C.struct_RcOverride
	BufferRef                   C.struct_AVBufferRef
	CodecParameters             C.struct_AVCodecParameters
	AudioServiceType            C.enum_AVAudioServiceType
	ChromaLocation              C.enum_AVChromaLocation
	CodecId                     C.enum_AVCodecID
	ColorPrimaries              C.enum_AVColorPrimaries
	ColorRange                  C.enum_AVColorRange
	ColorSpace                  C.enum_AVColorSpace
	ColorTransferCharacteristic C.enum_AVColorTransferCharacteristic
	Discard                     C.enum_AVDiscard
	FieldOrder                  C.enum_AVFieldOrder
	PacketSideDataType          C.enum_AVPacketSideDataType
	PixelFormat                 C.enum_AVPixelFormat
	SampleFormat                C.enum_AVSampleFormat
)

// nextRegisteredCodec returns the codec registered after the given codec, or the first one if nil is given.
//
// C-Function: av_codec_next
func nextRegisteredCodec(c *Codec) *Codec {
	return (*Codec)(C.av_codec_next((*C.struct_AVCodec)(c)))
}

// RegisteredCodecs returns a channel which can be used to iterate over the registered codecs.
//
// C-Function: av_codec_next
//
// Usage:
//
//     for codec := range avcodec.RegisteredCodecs() {
//         // ...
//     }
func RegisteredCodecs() <-chan *Codec {
	ch := make(chan *Codec)

	var c *Codec
	go func() {
		for {
			c = nextRegisteredCodec(c)
			if c == nil {
				break
			}
			ch <- c
		}
		close(ch)
	}()

	return ch
}

// RegisterCodec registers the codec and initializes libavcodec.
//
// C-Function: avcodec_register
func RegisterCodec(c *Codec) {
	C.avcodec_register((*C.struct_AVCodec)(c))
}

// FastPaddedMalloc allocates a buffer, reusing the given one if large enough.
// The buffer has additional FF_INPUT_BUFFER_PADDING_SIZE at the end which will always be 0.
// C-Funtion: av_fast_padded_malloc
func FastPaddedMalloc(p unsafe.Pointer, s *uint, t uintptr) {
	C.av_fast_padded_malloc(p, (*C.uint)(unsafe.Pointer(s)), (C.size_t)(t))
}

// Version returns the LIBAvCODEC_VERSION_INT constant.
//
// C-Function: avcodec_version
func Version() uint {
	return uint(C.avcodec_version())
}

// Configuration returns the libavcodec build-time configuration.
//
// C-Function: avcodec_configuration
func Configuration() string {
	return C.GoString(C.avcodec_configuration())

}

// License returns the libavcodec license.
//
// C-Function: avcodec_license
func License() string {
	return C.GoString(C.avcodec_license())
}

// RegisterAll registers all the codecs, parsers and bitstream filters which were enabled at configuration time.
//
// C-Function: avcodec_register_all
func RegisterAll() {
	C.avcodec_register_all()
}

// GetClass returns the Class for CodecContext.
//
// C-Function: avcodec_get_class
func GetClass() *Class {
	return (*Class)(C.avcodec_get_class())
}

// GetFrameClass returns the Class for Frame.
//
// C-Function: avcodec_get_frame_class
func GetFrameClass() *Class {
	return (*Class)(C.avcodec_get_frame_class())
}

// GetSubtitleRectClass returns the Class for SubtitleRect.
//
// C-Function: avcodec_get_subtitle_rect_class
func GetSubtitleRectClass() *Class {
	return (*Class)(C.avcodec_get_subtitle_rect_class())
}

// Free frees all allocated data of the Subtitle.
//
// C-Function: avsubtitle_free
func (s *Subtitle) Free() {
	C.avsubtitle_free((*C.struct_AVSubtitle)(s))
}

// Pack packs the dictionary for use in side_data.
//
// C-Function: av_packet_pack_dictionary
func (d *Dictionary) Pack(s *int) *uint8 {
	return (*uint8)(C.av_packet_pack_dictionary((*C.struct_AVDictionary)(d), (*C.int)(unsafe.Pointer(s))))
}

// UnpackDictionary unpacks a dictionary from side_data.
//
// C-Function: av_packet_unpack_dictionary
func UnpackDictionary(d *uint8, s int, dt **Dictionary) error {
	return avutil.CodeToError(int(C.av_packet_unpack_dictionary((*C.uint8_t)(d), C.int(s), (**C.struct_AVDictionary)(unsafe.Pointer(dt)))))
}

// ToChromaPos converts the ChromaLocation to swscale x/y chroma position.
//
// C-Function: avcodec_enum_to_chroma_pos
func (l ChromaLocation) ToChromaPos() (x, y int, err error) {
	err = avutil.CodeToError(int(C.avcodec_enum_to_chroma_pos((*C.int)(unsafe.Pointer(&x)), (*C.int)(unsafe.Pointer(&y)), (C.enum_AVChromaLocation)(l))))
	return
}

// ChromaPosToEnum converts swscale x/y chroma position to ChromaLocation.
//
// C-Function: avcodec_chroma_pos_to_enum
func ChromaPosToEnum(x, y int) ChromaLocation {
	return (ChromaLocation)(C.avcodec_chroma_pos_to_enum(C.int(x), C.int(y)))
}

// GetCodecTagString returns a string representing the codec tag codec_tag.
//
// C-Function: av_get_codec_tag_string
func GetCodecTagString(c uint) string {
	buffsize := 32
	var neededSize int
	var buf []C.char

	buf = make([]C.char, buffsize)
	neededSize = int(C.av_get_codec_tag_string((*C.char)(unsafe.Pointer(&buf[0])), C.size_t(buffsize), C.uint(c)))

	if neededSize+1 > buffsize {
		buffsize = neededSize + 1
		buf = make([]C.char, buffsize)
		_ = int(C.av_get_codec_tag_string((*C.char)(unsafe.Pointer(&buf[0])), C.size_t(buffsize), C.uint(c)))
	}

	return C.GoString((*C.char)(unsafe.Pointer(&buf[0])))
}

// FillAudioFrame fills the Frame audio data and linesize pointers.
//
// C-Function: avcodec_fill_audio_frame
func (f *Frame) FillAudioFrame(c int, s SampleFormat, b *uint8, bs, a int) int {
	return int(C.avcodec_fill_audio_frame((*C.struct_AVFrame)(f), C.int(c), (C.enum_AVSampleFormat)(s), (*C.uint8_t)(b), C.int(bs), C.int(a)))
}

// PcmCodec returns the PCM codec associated with the sample format.
//
// C-Function: av_get_pcm_codec
func (f SampleFormat) PcmCodec(b int) CodecId {
	return (CodecId)(C.av_get_pcm_codec((C.enum_AVSampleFormat)(f), C.int(b)))
}

// FastPaddedMallocz allocates a buffer, reusing the given one if large enough and initializes the data with 0.
// The buffer has additional FF_INPUT_BUFFER_PADDING_SIZE at the end which will always be 0.
//
// C-Function: av_fast_padded_mallocz
func FastPaddedMallocz(p unsafe.Pointer, s *uint, t uintptr) {
	C.av_fast_padded_mallocz(p, (*C.uint)(unsafe.Pointer(s)), (C.size_t)(t))
}

// Xiphlacing encodes extradata length to a buffer.
func Xiphlacing(v uint) []byte {
	buf := make([]byte, v/255+1)
	numBytes := uint(C.av_xiphlacing((*C.uchar)(unsafe.Pointer(&buf[0])), (C.uint)(v)))
	return buf[:numBytes]
}

// nextRegisteredCodec returns the HWAccel registered after the given HWAccel, or the first one if nil is given.
//
// C-Function: av_hwaccel_next
func nextHWAccel(a *HWAccel) *HWAccel {
	return (*HWAccel)(C.av_hwaccel_next((*C.struct_AVHWAccel)(a)))
}

// RegisteredHWAccels returns a channel which can be used to iterate over the registered HWAccel.
//
// C-Function: av_hwaccel_next
//
// Usage:
//
//     for hwa := range avcodec.RegisteredHWAccels() {
//         // ...
//     }
func RegisteredHWAccels() <-chan *HWAccel {
	ch := make(chan *HWAccel)

	var h *HWAccel
	go func() {
		for {
			h = nextHWAccel(h)
			if h == nil {
				break
			}
			ch <- h
		}
		close(ch)
	}()

	return ch
}

// nextDescriptor returns the CodecDescriptor registered after the given CodecDescriptor, or the first one if nil is given.
//
// C-Function: avcodec_descriptor_next
func nextDescriptor(d *CodecDescriptor) *CodecDescriptor {
	return (*CodecDescriptor)(C.avcodec_descriptor_next((*C.struct_AVCodecDescriptor)(d)))
}

// RegisteredCodecDescriptors returns a channel which can be used to iterate over the registered CodecDescriptor.
//
// C-Function: avcodec_descriptor_next
//
// Usage:
//
//     for cc := range avcodec.RegisteredCodecDescriptors() {
//         // ...
//     }
func RegisteredCodecDescriptors() <-chan *CodecDescriptor {
	ch := make(chan *CodecDescriptor)

	var cc *CodecDescriptor
	go func() {
		for {
			cc = nextDescriptor(cc)
			if cc == nil {
				break
			}
			ch <- cc
		}
		close(ch)
	}()

	return ch
}

// GetDescriptorByName returns the CodecDescriptor identified by the given name.
func GetDescriptorByName(n string) *CodecDescriptor {
	return (*CodecDescriptor)(C.avcodec_descriptor_get_by_name(C.CString(n)))
}

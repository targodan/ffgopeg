// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)
// Corbatto (luca@corbatto.de)

package avcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"
import (
	"unsafe"

	"github.com/colek42/ffgopeg/avutil"
)

// Open initializes the CodecContext to use the given Codec.
// C-Function: avcodec_open2
func (ctxt *CodecContext) Open(codec *Codec, d **Dictionary) avutil.ReturnCode {
	return avutil.NewReturnCode(int(C.avcodec_open2((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVCodec)(codec), (**C.struct_AVDictionary)(unsafe.Pointer(d)))))
}

// PktTimebase returns the packet timebase.
// C-Function: av_codec_get_pkt_timebase
func (ctxt *CodecContext) PktTimebase() Rational {
	return (Rational)(C.av_codec_get_pkt_timebase((*C.struct_AVCodecContext)(ctxt)))
}

// SetPktTimebase sets the packet timebase.
// C-Function: av_codec_set_pkt_timebase
func (ctxt *CodecContext) SetPktTimebase(r Rational) {
	C.av_codec_set_pkt_timebase((*C.struct_AVCodecContext)(ctxt), (C.struct_AVRational)(r))
}

// CodecDescriptor returns the CodecDescriptor.
// C-Function: av_codec_get_codec_descriptor
func (ctxt *CodecContext) CodecDescriptor() *CodecDescriptor {
	return (*CodecDescriptor)(C.av_codec_get_codec_descriptor((*C.struct_AVCodecContext)(ctxt)))
}

// SetCodecDescriptor sets the CodecDescriptor.
// C-Function: av_codec_set_codec_descriptor
func (ctxt *CodecContext) SetCodecDescriptor(d *CodecDescriptor) {
	C.av_codec_set_codec_descriptor((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVCodecDescriptor)(d))
}

// Lowres returns the low resolution decoding.
// 1 equals 1/2 size. 2 equals 1/4 size.
// C-Function: av_codec_get_lowres
func (ctxt *CodecContext) Lowres() int {
	return int(C.av_codec_get_lowres((*C.struct_AVCodecContext)(ctxt)))
}

// SetLowres sets the low resolution decoding.
// 1 equals 1/2 size. 2 equals 1/4 size.
// C-Function: av_codec_set_lowres
func (ctxt *CodecContext) SetLowres(i int) {
	C.av_codec_set_lowres((*C.struct_AVCodecContext)(ctxt), C.int(i))
}

// SeekPreroll returns something undocumented.
// C-Function: av_codec_get_seek_preroll
func (ctxt *CodecContext) SeekPreroll() int {
	return int(C.av_codec_get_seek_preroll((*C.struct_AVCodecContext)(ctxt)))
}

// SetSeekPreroll sets something undocumented.
// C-Function: av_codec_set_seek_preroll
func (ctxt *CodecContext) SetSeekPreroll(i int) {
	C.av_codec_set_seek_preroll((*C.struct_AVCodecContext)(ctxt), C.int(i))
}

// ChromaIntraMatrix returns something undocumented.
// C-Function: av_codec_get_chroma_intra_matrix
func (ctxt *CodecContext) ChromaIntraMatrix() *uint16 {
	return (*uint16)(C.av_codec_get_chroma_intra_matrix((*C.struct_AVCodecContext)(ctxt)))
}

// SetChromaIntraMatrix sets something undocumented.
// C-Function: av_codec_get_chroma_intra_matrix
func (ctxt *CodecContext) SetChromaIntraMatrix(t *uint16) {
	C.av_codec_set_chroma_intra_matrix((*C.struct_AVCodecContext)(ctxt), (*C.uint16_t)(t))
}

// Free frees the codec context and everything associated with it.
// C-Function: avcodec_free_context
func (ctxt *CodecContext) Free() {
	C.avcodec_free_context((**C.struct_AVCodecContext)(unsafe.Pointer(&ctxt)))
}

// Close closes a given Context and free all the data associated with it (but not the Context itself).
// C-Function: avcodec_close
func (ctxt *CodecContext) Close() int {
	return int(C.avcodec_close((*C.struct_AVCodecContext)(ctxt)))
}

// AlignDimensions modifies width and height values so that they will result in a memory buffer that is acceptable for the codec if you do not use any horizontal padding.
// May only be used if a codec with AV_CODEC_CAP_DR1 has been opened.
// C-Function: avcodec_align_dimensions
func (ctxt *CodecContext) AlignDimensions(w, h *int) {
	C.avcodec_align_dimensions((*C.struct_AVCodecContext)(ctxt), (*C.int)(unsafe.Pointer(w)), (*C.int)(unsafe.Pointer(h)))
}

// AlignDimensions2 modifies width and height values so that they will result in a memory buffer that is acceptable for the codec if you also ensure that all line sizes are a multiple of the respective linesize_align[i].
// May only be used if a codec with AV_CODEC_CAP_DR1 has been opened.
// C-Function: avcodec_align_dimensions2
func (ctxt *CodecContext) AlignDimensions2(w, h *int, l []int) {
	C.avcodec_align_dimensions2((*C.struct_AVCodecContext)(ctxt), (*C.int)(unsafe.Pointer(w)), (*C.int)(unsafe.Pointer(h)), (*C.int)(unsafe.Pointer(&l[0])))
}

// SendPacket sends a packet as input to the decoder.
// C-Function: avcodec_send_packet
func (ctxt *CodecContext) SendPacket(pkt *Packet) avutil.ReturnCode {
	return avutil.NewReturnCode(int(C.avcodec_send_packet((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVPacket)(pkt))))
}

// ReceiveFrame receives a frame as output from the decoder.
// C-Function: avcodec_receive_frame
func (ctxt *CodecContext) ReceiveFrame(frame *avutil.Frame) avutil.ReturnCode {
	return avutil.NewReturnCode(int(C.avcodec_receive_frame((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVFrame)(unsafe.Pointer(frame)))))
}

// SendFrame sends a frame as input to the encoder.
// C-Function: avcodec_send_frame
func (ctxt *CodecContext) SendFrame(frame *avutil.Frame) avutil.ReturnCode {
	return avutil.NewReturnCode(int(C.avcodec_send_frame((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVFrame)(unsafe.Pointer(frame)))))
}

// ReceivePacket receives a packet as output from the decoder.
// C-Function: avcodec_receive_packet
func (ctxt *CodecContext) ReceivePacket(pkt *Packet) avutil.ReturnCode {
	return avutil.NewReturnCode(int(C.avcodec_receive_packet((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVPacket)(pkt))))
}

// DecodeSubtitle decodes a subtitle message.
//
// Return a negative value on error, otherwise return the number of bytes used. If no subtitle could be decompressed, got_sub_ptr is zero. Otherwise, the subtitle is stored in *sub. Note that AV_CODEC_CAP_DR1 is not available for subtitle codecs. This is for simplicity, because the performance difference is expect to be negligible and reusing a get_buffer written for video codecs would probably perform badly due to a potentially very different allocation pattern.
//
// Some decoders (those marked with CODEC_CAP_DELAY) have a delay between input and output. This means that for some packets they will not immediately produce decoded output and need to be flushed at the end of decoding to get all the decoded data. Flushing is done by calling this function with packets with avpkt->data set to NULL and avpkt->size set to 0 until it stops returning subtitles. It is safe to flush even those decoders that are not marked with CODEC_CAP_DELAY, then no subtitles will be returned.
// C-Function: avcodec_decode_subtitle2
func (ctxt *CodecContext) DecodeSubtitle(s *Subtitle, g *int, a *Packet) int {
	return int(C.avcodec_decode_subtitle2((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVSubtitle)(s), (*C.int)(unsafe.Pointer(g)), (*C.struct_AVPacket)(a)))
}

// EncodeSubtitle encodes a subtitle message.
// C-Function: avcodec_encode_subtitle
func (ctxt *CodecContext) EncodeSubtitle(b *uint8, bs int, s *Subtitle) int {
	return int(C.avcodec_encode_subtitle((*C.struct_AVCodecContext)(ctxt), (*C.uint8_t)(b), C.int(bs), (*C.struct_AVSubtitle)(s)))
}

// DefaultGetFormat is undocumented...
// C-Function: avcodec_default_get_format
func (ctxt *CodecContext) DefaultGetFormat(f *avutil.PixelFormat) avutil.PixelFormat {
	return (avutil.PixelFormat)(C.avcodec_default_get_format((*C.struct_AVCodecContext)(ctxt), (*C.enum_AVPixelFormat)(f)))
}

// FlushBuffers resets the internal decoder state / flush internal buffers.
// Should be called e.g. when seeking or when switching to a different stream.
// C-Function: avcodec_flush_buffers
func (ctxt *CodecContext) FlushBuffers() {
	C.avcodec_flush_buffers((*C.struct_AVCodecContext)(ctxt))
}

// AudioFrameDuration returns audio frame duration, or 0 if not able to determine.
// C-Function: av_get_audio_frame_duration
func (ctxt *CodecContext) AudioFrameDuration(frameSize int) int {
	return int(C.av_get_audio_frame_duration((*C.struct_AVCodecContext)(ctxt), C.int(frameSize)))
}

// IsOpen returns true iff the CodecContext was opened and not yet closed.
// C-Function: avcodec_is_open
func (ctxt *CodecContext) IsOpen() bool {
	return int(C.avcodec_is_open((*C.struct_AVCodecContext)(ctxt))) != 0
}

// Parse is not yet completely implemented.
// C-Function: av_parser_parse2
// TODO
func (ctxt *CodecContext) Parse(ctxtp *CodecParserContext, p **uint8, ps *int, b *uint8, bs int, pt, dt, po int64) int {
	// TODO: implement correctly giving it a buffer, returning the output data and so on.
	return int(C.av_parser_parse2((*C.struct_AVCodecParserContext)(ctxtp), (*C.struct_AVCodecContext)(ctxt), (**C.uint8_t)(unsafe.Pointer(p)), (*C.int)(unsafe.Pointer(ps)), (*C.uint8_t)(b), C.int(bs), (C.int64_t)(pt), (C.int64_t)(dt), (C.int64_t)(po)))
}

// ParserChange is not yet completely implemented.
// C-Function: av_parser_change
// TODO
func (ctxt *CodecContext) ParserChange(ctxtp *CodecParserContext, pb **uint8, pbs *int, b *uint8, bs, k int) int {
	// TODO: implement correctly giving it a buffer, returning the output data and so on.
	return int(C.av_parser_change((*C.struct_AVCodecParserContext)(ctxtp), (*C.struct_AVCodecContext)(ctxt), (**C.uint8_t)(unsafe.Pointer(pb)), (*C.int)(unsafe.Pointer(pbs)), (*C.uint8_t)(b), C.int(bs), C.int(k)))
}

// NewCodecParserContext creates a new CodecParserContext.
// C-Function: av_parser_init
func NewCodecParserContext(c int) *CodecParserContext {
	return (*CodecParserContext)(C.av_parser_init(C.int(c)))
}

// Close closes the CodecParserContext.
// C-Function: av_parser_close
func (ctxtp *CodecParserContext) Close() {
	C.av_parser_close((*C.struct_AVCodecParserContext)(ctxtp))
}

// nextParser returns the CodecParser registered after the given CodecParser, or the first one if nil is given.
// C-Function: av_parser_next
func nextParser(p *CodecParser) *CodecParser {
	return (*CodecParser)(C.av_parser_next((*C.struct_AVCodecParser)(p)))
}

// RegisteredCodecParsers returns a channel which can be used to iterate over the registered CodecParser.
// C-Function: av_parser_next
//
// Usage:
//
//     for cc := range avcodec.RegisteredCodecDescriptors() {
//         // ...
//     }
func RegisteredCodecParsers() <-chan *CodecParser {
	ch := make(chan *CodecParser)

	var cp *CodecParser
	go func() {
		for {
			cp = nextParser(cp)
			if cp == nil {
				break
			}
			ch <- cp
		}
		close(ch)
	}()

	return ch
}

// RegisterCodecParser registeres a CodecParser.
func RegisterCodecParser(p *CodecParser) {
	C.av_register_codec_parser((*C.struct_AVCodecParser)(p))
}

// FromParameters fills the previously initializes CodecContext with the given parameters.
// C-Function: avcodec_parameters_to_context
func (ctxt *CodecContext) FromParameters(par *CodecParameters) avutil.ReturnCode {
	return avutil.NewReturnCode(int(C.avcodec_parameters_to_context((*C.AVCodecContext)(ctxt), (*C.AVCodecParameters)(par))))
}

// ToParameters returns a CodecParameters instance from a CodecContext
// C-Function: avcodec_parameters_from_contex
func (ctxt *CodecContext) ToParameters() (*CodecParameters, avutil.ReturnCode) {
	par := &CodecParameters{}
	ret := int(C.avcodec_parameters_from_context((*C.AVCodecParameters)(par), (*C.AVCodecContext)(ctxt)))
	return par, avutil.NewReturnCode(ret)
}

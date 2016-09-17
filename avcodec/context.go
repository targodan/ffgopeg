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

func (ctxt *CodecContext) Codec() *Codec {
	if ctxt.IsOpen() == false {
		codec := C.avcodec_find_decoder(ctxt.codec_id)
		if codec == nil {
			panic("Codec not found")
		}

		C.avcodec_open2((*C.struct_AVCodecContext)(ctxt), codec, nil)
	}

	return (*Codec)(ctxt.codec)
}

func (ctxt *CodecContext) GetPktTimebase() Rational {
	return (Rational)(C.av_codec_get_pkt_timebase((*C.struct_AVCodecContext)(ctxt)))
}

func (ctxt *CodecContext) SetPktTimebase(r Rational) {
	C.av_codec_set_pkt_timebase((*C.struct_AVCodecContext)(ctxt), (C.struct_AVRational)(r))
}

func (ctxt *CodecContext) GetCodecDescriptor() *CodecDescriptor {
	return (*CodecDescriptor)(C.av_codec_get_codec_descriptor((*C.struct_AVCodecContext)(ctxt)))
}

func (ctxt *CodecContext) SetCodecDescriptor(d *CodecDescriptor) {
	C.av_codec_set_codec_descriptor((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVCodecDescriptor)(d))
}

func (ctxt *CodecContext) GetLowres() int {
	return int(C.av_codec_get_lowres((*C.struct_AVCodecContext)(ctxt)))
}

func (ctxt *CodecContext) SetLowres(i int) {
	C.av_codec_set_lowres((*C.struct_AVCodecContext)(ctxt), C.int(i))
}

func (ctxt *CodecContext) GetSeekPreroll() int {
	return int(C.av_codec_get_seek_preroll((*C.struct_AVCodecContext)(ctxt)))
}

func (ctxt *CodecContext) SetSeekPreroll(i int) {
	C.av_codec_set_seek_preroll((*C.struct_AVCodecContext)(ctxt), C.int(i))
}

func (ctxt *CodecContext) GetChromaIntraMatrix() *uint16 {
	return (*uint16)(C.av_codec_get_chroma_intra_matrix((*C.struct_AVCodecContext)(ctxt)))
}

func (ctxt *CodecContext) SetChromaIntraMatrix(t *uint16) {
	C.av_codec_set_chroma_intra_matrix((*C.struct_AVCodecContext)(ctxt), (*C.uint16_t)(t))
}

//Free the codec context and everything associated with it and write NULL to the provided pointer.
func (ctxt *CodecContext) FreeContext() {
	C.avcodec_free_context((**C.struct_AVCodecContext)(unsafe.Pointer(ctxt)))
}

//Set the fields of the given Context to default values corresponding to the given codec (defaults may be codec-dependent).
func (ctxt *CodecContext) GetContextDefaults3(c *Codec) int {
	return int(C.avcodec_get_context_defaults3((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVCodec)(c)))
}

//Copy the settings of the source Context into the destination Context.
func (ctxt *CodecContext) CopyContext(ctxt2 *CodecContext) int {
	return int(C.avcodec_copy_context((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVCodecContext)(ctxt2)))
}

//Initialize the Context to use the given Codec
func (ctxt *CodecContext) Open2(c *Codec, d **Dictionary) int {
	return int(C.avcodec_open2((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVCodec)(c), (**C.struct_AVDictionary)(unsafe.Pointer(d))))
}

//Close a given Context and free all the data associated with it (but not the Context itself).
func (ctxt *CodecContext) Close() int {
	return int(C.avcodec_close((*C.struct_AVCodecContext)(ctxt)))
}

//The default callback for Context.get_buffer2().
func (s *CodecContext) DefaultGetBuffer2(f *Frame, l int) int {
	return int(C.avcodec_default_get_buffer2((*C.struct_AVCodecContext)(s), (*C.struct_AVFrame)(f), C.int(l)))
}

//Modify width and height values so that they will result in a memory buffer that is acceptable for the codec if you do not use any horizontal padding.
func (ctxt *CodecContext) AlignDimensions(w, h *int) {
	C.avcodec_align_dimensions((*C.struct_AVCodecContext)(ctxt), (*C.int)(unsafe.Pointer(w)), (*C.int)(unsafe.Pointer(h)))
}

//Modify width and height values so that they will result in a memory buffer that is acceptable for the codec if you also ensure that all line sizes are a multiple of the respective linesize_align[i].
func (ctxt *CodecContext) AlignDimensions2(w, h *int, l int) {
	C.avcodec_align_dimensions2((*C.struct_AVCodecContext)(ctxt), (*C.int)(unsafe.Pointer(w)), (*C.int)(unsafe.Pointer(h)), (*C.int)(unsafe.Pointer(&l)))
}

//Decode the audio frame of size avpkt->size from avpkt->data into frame.
func (ctxt *CodecContext) DecodeAudio4(f *Frame, g *int, a *Packet) int {
	return int(C.avcodec_decode_audio4((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVFrame)(f), (*C.int)(unsafe.Pointer(g)), (*C.struct_AVPacket)(a)))
}

//Decode the video frame of size avpkt->size from avpkt->data into picture.
func (ctxt *CodecContext) DecodeVideo2(p *Frame, g *int, a *Packet) int {
	return int(C.avcodec_decode_video2((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVFrame)(p), (*C.int)(unsafe.Pointer(g)), (*C.struct_AVPacket)(a)))
}

//Decode a subtitle message.
func (ctxt *CodecContext) DecodeSubtitle2(s *Subtitle, g *int, a *Packet) int {
	return int(C.avcodec_decode_subtitle2((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVSubtitle)(s), (*C.int)(unsafe.Pointer(g)), (*C.struct_AVPacket)(a)))
}

//Encode a frame of audio.
func (ctxt *CodecContext) EncodeAudio2(p *Packet, f *Frame, gp *int) int {
	return int(C.avcodec_encode_audio2((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVPacket)(p), (*C.struct_AVFrame)(f), (*C.int)(unsafe.Pointer(gp))))
}

//Encode a frame of video
func (ctxt *CodecContext) EncodeVideo2(p *Packet, f *Frame, gp *int) int {
	return int(C.avcodec_encode_video2((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVPacket)(p), (*C.struct_AVFrame)(f), (*C.int)(unsafe.Pointer(gp))))
}

func (ctxt *CodecContext) EncodeSubtitle(b *uint8, bs int, s *Subtitle) int {
	return int(C.avcodec_encode_subtitle((*C.struct_AVCodecContext)(ctxt), (*C.uint8_t)(b), C.int(bs), (*C.struct_AVSubtitle)(s)))
}

func (ctxt *CodecContext) DefaultGetFormat(f *PixelFormat) PixelFormat {
	return (PixelFormat)(C.avcodec_default_get_format((*C.struct_AVCodecContext)(ctxt), (*C.enum_AVPixelFormat)(f)))
}

//Reset the internal decoder state / flush internal buffers.
func (ctxt *CodecContext) FlushBuffers() {
	C.avcodec_flush_buffers((*C.struct_AVCodecContext)(ctxt))
}

//Return audio frame duration.
func (ctxt *CodecContext) GetAudioFrameDuration(f int) int {
	return int(C.av_get_audio_frame_duration((*C.struct_AVCodecContext)(ctxt), C.int(f)))
}

func (ctxt *CodecContext) IsOpen() bool {
	return int(C.avcodec_is_open((*C.struct_AVCodecContext)(ctxt))) == 1
}

//Parse a packet.
func (ctxt *CodecContext) Parse2(ctxtp *CodecParserContext, p **uint8, ps *int, b *uint8, bs int, pt, dt, po int64) int {
	return int(C.av_parser_parse2((*C.struct_AVCodecParserContext)(ctxtp), (*C.struct_AVCodecContext)(ctxt), (**C.uint8_t)(unsafe.Pointer(p)), (*C.int)(unsafe.Pointer(ps)), (*C.uint8_t)(b), C.int(bs), (C.int64_t)(pt), (C.int64_t)(dt), (C.int64_t)(po)))
}

func (ctxt *CodecContext) ParserChange(ctxtp *CodecParserContext, pb **uint8, pbs *int, b *uint8, bs, k int) int {
	return int(C.av_parser_change((*C.struct_AVCodecParserContext)(ctxtp), (*C.struct_AVCodecContext)(ctxt), (**C.uint8_t)(unsafe.Pointer(pb)), (*C.int)(unsafe.Pointer(pbs)), (*C.uint8_t)(b), C.int(bs), C.int(k)))
}

func ParserInit(c int) *CodecParserContext {
	return (*CodecParserContext)(C.av_parser_init(C.int(c)))
}

func ParserClose(ctxtp *CodecParserContext) {
	C.av_parser_close((*C.struct_AVCodecParserContext)(ctxtp))
}

func (p *CodecParser) ParserNext() *CodecParser {
	return (*CodecParser)(C.av_parser_next((*C.struct_AVCodecParser)(p)))
}

func (p *CodecParser) RegisterCodecParser() {
	C.av_register_codec_parser((*C.struct_AVCodecParser)(p))
}

// FromParameters fills the previously initializes CodecContext with the given parameters.
func (ctxt *CodecContext) FromParameters(par *CodecParameters) error {
	return avutil.CodeToError(int(C.avcodec_parameters_to_context((*C.AVCodecContext)(ctxt), (*C.AVCodecParameters)(par))))
}

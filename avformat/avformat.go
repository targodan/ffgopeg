// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)
// Corbatto (luca@corbatto.de)

//Package avformat provides some generic global options, which can be set on all the muxers and demuxers.
//In addition each muxer or demuxer may support so-called private options, which are specific for that component.
//Supported formats (muxers and demuxers) provided by the libavformat library
package avformat

//#cgo pkg-config: libavformat libavcodec libavutil libavdevice libavfilter libswresample libswscale
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
//#include <libavutil/avutil.h>
//#include <libavutil/opt.h>
//#include <libavdevice/avdevice.h>
import "C"
import (
	"unsafe"

	"gopkg.in/targodan/ffgopeg.v0/avutil"
)

type (
	ProbeData                C.struct_AVProbeData
	InputFormat              C.struct_AVInputFormat
	OutputFormat             C.struct_AVOutputFormat
	FormatContext            C.struct_AVFormatContext
	Frame                    C.struct_AVFrame
	CodecFormatContext       C.struct_AVCodecFormatContext
	CodecParserContext       C.struct_AVCodecParserContext
	Dictionary               C.struct_AVDictionary
	DictionaryEntry          C.struct_AVDictionaryEntry
	IndexEntry               C.struct_AVIndexEntry
	Stream                   C.struct_AVStream
	Program                  C.struct_AVProgram
	Chapter                  C.struct_AVChapter
	PacketList               C.struct_AVPacketList
	Packet                   C.struct_AVPacket
	CodecParserFormatContext C.struct_AVCodecParserFormatContext
	IOContext                C.struct_AVIOContext
	IOFormatContext          C.struct_AVIOFormatContext
	Rational                 C.struct_AVRational
	Codec                    C.struct_AVCodec
	CodecTag                 C.struct_AVCodecTag
	Class                    C.struct_AVClass
	FormatInternal           C.struct_AVFormatInternal
	IOInterruptCB            C.struct_AVIOInterruptCB
	PacketSideData           C.struct_AVPacketSideData
	FFFrac                   C.struct_FFFrac
	StreamParseType          C.enum_AVStreamParseType
	Discard                  C.enum_AVDiscard
	MediaType                C.enum_AVMediaType
	DurationEstimationMethod C.enum_AVDurationEstimationMethod
	PacketSideDataType       C.enum_AVPacketSideDataType
	CodecId                  C.enum_AVCodecID
)

type File C.FILE

// TODO: Find a nice way to convert os.File into C.FILE.

// GetPacket allocates and reads the payload of a packet and initialize its fields with default values.
//
// C-Function: av_get_packet
func (ctxt *IOContext) GetPacket(size int) (*Packet, avutil.ReturnCode) {
	var pkt *Packet
	code := avutil.NewReturnCode(int(C.av_get_packet((*C.struct_AVIOContext)(ctxt), (*C.struct_AVPacket)(pkt), C.int(size))))
	return pkt, code
}

//Read data and append it to the current content of the Packet.
//
// C-Function: av_append_packet
func (ctxt *IOContext) AvAppendPacket(pkt *Packet, s int) int {
	return int(C.av_append_packet((*C.struct_AVIOContext)(ctxt), (*C.struct_AVPacket)(pkt), C.int(s)))
}

// Register registers the InputFormat.
//
// C-Function: av_register_input_format
func (f *InputFormat) Register() {
	C.av_register_input_format((*C.struct_AVInputFormat)(f))
}

// Register registers the OutputFormat.
//
// C-Function: av_register_output_format
func (f *OutputFormat) Register() {
	C.av_register_output_format((*C.struct_AVOutputFormat)(f))
}

func informatNext(f *InputFormat) *InputFormat {
	return (*InputFormat)(C.av_iformat_next((*C.struct_AVInputFormat)(f)))
}

// InputFormats returns a channel which can be used to iterate over the input formats.
//
// Usage:
//
//     for in := range avformat.InputFormats() {
//         // ...
//     }
//
// C-Function: av_iformat_next
func InputFormats() <-chan *InputFormat {
	ch := make(chan *InputFormat)

	var e *InputFormat
	go func() {
		for {
			e = informatNext(e)
			if e == nil {
				break
			}
			ch <- e
		}
		close(ch)
	}()

	return ch
}

func oformatNext(f *OutputFormat) *OutputFormat {
	return (*OutputFormat)(C.av_oformat_next((*C.struct_AVOutputFormat)(f)))
}

// OutputFormats returns a channel which can be used to iterate over the output formats.
//
// Usage:
//
//     for out := range avformat.OutputFormats() {
//         // ...
//     }
//
// C-Function: av_oformat_next
func OutputFormats() <-chan *OutputFormat {
	ch := make(chan *OutputFormat)

	var e *OutputFormat
	go func() {
		for {
			e = oformatNext(e)
			if e == nil {
				break
			}
			ch <- e
		}
		close(ch)
	}()

	return ch
}

// Version returns the LIBAVFORMAT_VERSION_INT constant.
//
// C-Function: avformat_version
func Version() uint {
	return uint(C.avformat_version())
}

// Configuration returns the libavformat build-time configuration.
//
// C-Function: avformat_configuration
func Configuration() string {
	return C.GoString(C.avformat_configuration())
}

// License returns the libavformat license.
//
// C-Function: avformat_license
func License() string {
	return C.GoString(C.avformat_license())
}

// RegisterAll initializes libavformat and register all the muxers, demuxers and protocols.
//
// C-Function: av_register_all
func RegisterAll() {
	C.av_register_all()
}

// NetworkInit does global initialization of network components.
//
// This is optional, but recommended, since it avoids the overhead of implicitly doing the setup for each session.
//
// Calling this function will become mandatory if using network protocols at some major version bump.
//
// C-Function: avformat_network_init
func NetworkInit() avutil.ReturnCode {
	return avutil.NewReturnCode(int(C.avformat_network_init()))
}

// NetworkDeinit undoes the initialization done by avformat_network_init.
//
// C-Function: avformat_network_deinit
func NetworkDeinit() int {
	return int(C.avformat_network_deinit())
}

// NewFormatContext allocates an FormatContext.
//
// C-Function: avformat_alloc_context
func NewFormatContext() *FormatContext {
	return (*FormatContext)(C.avformat_alloc_context())
}

// GetClass gets the Class for FormatContext.
//
// C-Function: avformat_get_class
func GetClass() *Class {
	return (*Class)(C.avformat_get_class())
}

// SideData returns side information from stream.
//
// C-Function: av_stream_get_side_data
func (s *Stream) SideData(t PacketSideDataType) []byte {
	var size int
	arr := (*[1 << 30]byte)(unsafe.Pointer(C.av_stream_get_side_data((*C.struct_AVStream)(s), (C.enum_AVPacketSideDataType)(t), (*C.int)(unsafe.Pointer(&size)))))
	return arr[:size:size]
}

// NewOutputFormatContext allocates a FormatContext for an output format.
//
// C-Function: avformat_alloc_output_context2
func NewOutputFormatContext(outFormat *OutputFormat, formatName, fileName string) (*FormatContext, avutil.ReturnCode) {
	var ctx *FormatContext
	code := avutil.NewReturnCode(int(C.avformat_alloc_output_context2((**C.struct_AVFormatContext)(unsafe.Pointer(&ctx)), (*C.struct_AVOutputFormat)(outFormat), C.CString(formatName), C.CString(fileName))))
	return ctx, code
}

// FindInputFormat finds InputFormat based on the short name of the input format.
//
// C-Function: av_find_input_format
func FindInputFormat(shortName string) *InputFormat {
	return (*InputFormat)(C.av_find_input_format(C.CString(shortName)))
}

// ProbeInputFormat guesses the file format.
//
// C-Function: av_probe_input_format
func ProbeInputFormat(pd *ProbeData, isOpened bool) *InputFormat {
	var iOpened int
	if isOpened {
		iOpened = 1
	} else {
		iOpened = 0
	}
	return (*InputFormat)(C.av_probe_input_format((*C.struct_AVProbeData)(pd), C.int(iOpened)))
}

// ProbeInputFormat2 guesses the file format.
//
// C-Function: av_probe_input_format2
func ProbeInputFormat2(pd *ProbeData, isOpened bool, scoreMax int) (*InputFormat, int) {
	var iOpened int
	if isOpened {
		iOpened = 1
	} else {
		iOpened = 0
	}
	fmt := (*InputFormat)(C.av_probe_input_format2((*C.struct_AVProbeData)(pd), C.int(iOpened), (*C.int)(unsafe.Pointer(&scoreMax))))
	return fmt, scoreMax
}

// ProbeInputFormat3 guesses the file format.
//
// C-Function: av_probe_input_format3
func ProbeInputFormat3(pd *ProbeData, isOpened bool) (*InputFormat, int) {
	var iOpened int
	if isOpened {
		iOpened = 1
	} else {
		iOpened = 0
	}
	var score int
	fmt := (*InputFormat)(C.av_probe_input_format3((*C.struct_AVProbeData)(pd), C.int(iOpened), (*C.int)(unsafe.Pointer(&score))))
	return fmt, score
}

// ProbeInputBuffer probes a bytestream to determine the input format.
//
// C-Function: av_probe_input_buffer2
func (pb *IOContext) ProbeInputBuffer(url string, offset, maxProbeSize uint) (*InputFormat, avutil.ReturnCode) {
	var f *InputFormat
	var l int // TODO: this doesn't make any sense
	code := avutil.NewReturnCode(int(C.av_probe_input_buffer2((*C.struct_AVIOContext)(pb), (**C.struct_AVInputFormat)(unsafe.Pointer(&f)), C.CString(url), unsafe.Pointer(&l), C.uint(offset), C.uint(maxProbeSize))))
	return f, code
}

// OpenInput opens an input stream and reads the header.
//
// C-Function: avformat_open_input
func OpenInput(url string, fmt *InputFormat, options **Dictionary) (*FormatContext, avutil.ReturnCode) {
	var ctxt *FormatContext
	err := avutil.NewReturnCode(int(C.avformat_open_input((**C.struct_AVFormatContext)(unsafe.Pointer(&ctxt)), C.CString(url), (*C.struct_AVInputFormat)(fmt), (**C.struct_AVDictionary)(unsafe.Pointer(options)))))
	return ctxt, err
}

//Return the output format in the list of registered output formats which best matches the provided parameters, or return NULL if there is no match.
//
// C-Function: av_guess_format
func AvGuessFormat(sn, f, mt string) *OutputFormat {
	return (*OutputFormat)(C.av_guess_format(C.CString(sn), C.CString(f), C.CString(mt)))
}

//Guess the codec ID based upon muxer and filename.
//
// C-Function: av_guess_codec
func AvGuessCodec(fmt *OutputFormat, sn, f, mt string, t avutil.MediaType) CodecId {
	return (CodecId)(C.av_guess_codec((*C.struct_AVOutputFormat)(fmt), C.CString(sn), C.CString(f), C.CString(mt), (C.enum_AVMediaType)(t)))
}

//Send a nice hexadecimal dump of a buffer to the specified file stream.
//
// C-Function: av_hex_dump
func AvHexDump(f *File, b *uint8, s int) {
	C.av_hex_dump((*C.FILE)(f), (*C.uint8_t)(b), C.int(s))
}

//Send a nice hexadecimal dump of a buffer to the log.
//
// C-Function: av_hex_dump_log
func AvHexDumpLog(a, l int, b *uint8, s int) {
	C.av_hex_dump_log(unsafe.Pointer(&a), C.int(l), (*C.uint8_t)(b), C.int(s))
}

//Send a nice dump of a packet to the specified file stream.
//
// C-Function: av_pkt_dump2
func AvPktDump2(f *File, pkt *Packet, dp int, st *Stream) {
	C.av_pkt_dump2((*C.FILE)(f), (*C.struct_AVPacket)(pkt), C.int(dp), (*C.struct_AVStream)(st))
}

//Send a nice dump of a packet to the log.
//
// C-Function: av_pkt_dump_log2
func AvPktDumpLog2(a int, l int, pkt *Packet, dp int, st *Stream) {
	C.av_pkt_dump_log2(unsafe.Pointer(&a), C.int(l), (*C.struct_AVPacket)(pkt), C.int(dp), (*C.struct_AVStream)(st))
}

//enum CodecId av_codec_get_id (const struct CodecTag *const *tags, unsigned int tag)
//Get the CodecId for the given codec tag tag.
//
// C-Function: av_codec_get_id
func AvCodecGetId(t **CodecTag, tag uint) CodecId {
	return (CodecId)(C.av_codec_get_id((**C.struct_AVCodecTag)(unsafe.Pointer(t)), C.uint(tag)))
}

//Get the codec tag for the given codec id id.
//
// C-Function: av_codec_get_tag
func AvCodecGetTag(t **CodecTag, id CodecId) uint {
	return uint(C.av_codec_get_tag((**C.struct_AVCodecTag)(unsafe.Pointer(t)), (C.enum_AVCodecID)(id)))
}

//Get the codec tag for the given codec id.
//
// C-Function: av_codec_get_tag2
func AvCodecGetTag2(t **CodecTag, id CodecId, tag *uint) int {
	return int(C.av_codec_get_tag2((**C.struct_AVCodecTag)(unsafe.Pointer(t)), (C.enum_AVCodecID)(id), (*C.uint)(unsafe.Pointer(tag))))
}

//Get the index for a specific timestamp.
//
// C-Function: av_index_search_timestamp
func AvIndexSearchTimestamp(st *Stream, t int64, f int) int {
	return int(C.av_index_search_timestamp((*C.struct_AVStream)(st), C.int64_t(t), C.int(f)))
}

//Add an index entry into a sorted list.
//
// C-Function: av_add_index_entry
func AvAddIndexEntry(st *Stream, pos, t, int64, s, d, f int) int {
	return int(C.av_add_index_entry((*C.struct_AVStream)(st), C.int64_t(pos), C.int64_t(t), C.int(s), C.int(d), C.int(f)))
}

//Split a URL string into components.
//
// C-Function: av_url_split
func AvUrlSplit(p string, ps int, a string, as int, h string, hs int, pp *int, path string, psize int, url string) {
	C.av_url_split(C.CString(p), C.int(ps), C.CString(a), C.int(as), C.CString(h), C.int(hs), (*C.int)(unsafe.Pointer(pp)), C.CString(path), C.int(psize), C.CString(url))
}

//int av_get_frame_filename (char *buf, int buf_size, const char *path, int number)
//Return in 'buf' the path with 'd' replaced by a number.
//
// C-Function: av_get_frame_filename
func AvGetFrameFilename(b string, bs int, pa string, n int) int {
	return int(C.av_get_frame_filename(C.CString(b), C.int(bs), C.CString(pa), C.int(n)))
}

//Check whether filename actually is a numbered sequence generator.
//
// C-Function: av_filename_number_test
func AvFilenameNumberTest(f string) int {
	return int(C.av_filename_number_test(C.CString(f)))
}

//Generate an SDP for an RTP session.
//
// C-Function: av_sdp_create
func AvSdpCreate(ac **FormatContext, nf int, b string, s int) int {
	return int(C.av_sdp_create((**C.struct_AVFormatContext)(unsafe.Pointer(ac)), C.int(nf), C.CString(b), C.int(s)))
}

//int av_match_ext (const char *filename, const char *extensions)
//Return a positive value if the given filename has one of the given extensions, 0 otherwise.
//
// C-Function: av_match_ext
func AvMatchExt(f, e string) int {
	return int(C.av_match_ext(C.CString(f), C.CString(e)))
}

//Test if the given container can store a codec.
//
// C-Function: avformat_query_codec
func AvformatQueryCodec(o *OutputFormat, cd CodecId, sc int) int {
	return int(C.avformat_query_codec((*C.struct_AVOutputFormat)(o), (C.enum_AVCodecID)(cd), C.int(sc)))
}

//
// C-Function: avformat_get_riff_video_tags
func AvformatGetRiffVideoTags() *CodecTag {
	return (*CodecTag)(C.avformat_get_riff_video_tags())
}

//struct CodecTag * avformat_get_riff_audio_tags (void)
//
// C-Function: avformat_get_riff_audio_tags
func AvformatGetRiffAudioTags() *CodecTag {
	return (*CodecTag)(C.avformat_get_riff_audio_tags())
}

//
// C-Function: avformat_get_mov_video_tags
func AvformatGetMovVideoTags() *CodecTag {
	return (*CodecTag)(C.avformat_get_mov_video_tags())
}

//
// C-Function: avformat_get_mov_audio_tags
func AvformatGetMovAudioTags() *CodecTag {
	return (*CodecTag)(C.avformat_get_mov_audio_tags())
}

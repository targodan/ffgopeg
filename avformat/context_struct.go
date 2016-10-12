// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)
// Corbatto (luca@corbatto.de)

package avformat

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"
import "unsafe"

// Chapters returns chapters.
//
// C-Variable: AVFormatContext::chapters
func (ctxt *FormatContext) Chapters() []*Chapter {
	return (*[1 << 30]*Chapter)(unsafe.Pointer(ctxt.chapters))[:ctxt.NbChapters():ctxt.NbChapters()]
}

//
// C-Variable: AVFormatContext::metadata
func (ctxt *FormatContext) Metadata() *Dictionary {
	return (*Dictionary)(unsafe.Pointer(ctxt.metadata))
}

//
// C-Variable: AVFormatContext::internal
func (ctxt *FormatContext) Internal() *FormatInternal {
	return (*FormatInternal)(unsafe.Pointer(ctxt.internal))
}

//
// C-Variable: AVFormatContext::pb
func (ctxt *FormatContext) Pb() *IOContext {
	return (*IOContext)(unsafe.Pointer(ctxt.pb))
}

//
// C-Variable: AVFormatContext::interrupt_callback
func (ctxt *FormatContext) InterruptCallback() IOInterruptCB {
	return IOInterruptCB(ctxt.interrupt_callback)
}

//
// C-Variable: AVFormatContext::programs
func (ctxt *FormatContext) Programs() **Program {
	return (**Program)(unsafe.Pointer(ctxt.programs))
}

//
// C-Variable: AVFormatContext::streams
func (ctxt *FormatContext) Streams() []*Stream {
	length := ctxt.NbStreams()
	return (*[1 << 30]*Stream)(unsafe.Pointer(ctxt.streams))[:length:length]
}

//
// C-Variable: AVFormatContext::filename
func (ctxt *FormatContext) Filename() string {
	return C.GoString((*C.char)(unsafe.Pointer(&ctxt.filename[0])))
}

// //
// C-Variable: AVFormatContext::codec_whitelist
func (ctxt *FormatContext) CodecWhitelist() string {
	return C.GoString(ctxt.codec_whitelist)
}

// //
// C-Variable: AVFormatContext::format_whitelist
func (ctxt *FormatContext) FormatWhitelist() string {
	return C.GoString(ctxt.format_whitelist)
}

//
// C-Variable: AVFormatContext::audio_codec_id
func (ctxt *FormatContext) AudioCodecId() CodecId {
	return CodecId(ctxt.audio_codec_id)
}

//
// C-Variable: AVFormatContext::subtitle_codec_id
func (ctxt *FormatContext) SubtitleCodecId() CodecId {
	return CodecId(ctxt.subtitle_codec_id)
}

//
// C-Variable: AVFormatContext::video_codec_id
func (ctxt *FormatContext) VideoCodecId() CodecId {
	return CodecId(ctxt.video_codec_id)
}

//
// C-Variable: AVFormatContext::audio_preload
func (ctxt *FormatContext) AudioPreload() int {
	return int(ctxt.audio_preload)
}

//
// C-Variable: AVFormatContext::avio_flags
func (ctxt *FormatContext) AvioFlags() int {
	return int(ctxt.avio_flags)
}

//
// C-Variable: AVFormatContext::avoid_negative_ts
func (ctxt *FormatContext) AvoidNegativeTs() int {
	return int(ctxt.avoid_negative_ts)
}

//
// C-Variable: AVFormatContext::bit_rate
func (ctxt *FormatContext) BitRate() int {
	return int(ctxt.bit_rate)
}

//
// C-Variable: AVFormatContext::ctx_flags
func (ctxt *FormatContext) CtxFlags() int {
	return int(ctxt.ctx_flags)
}

//
// C-Variable: AVFormatContext::debug
func (ctxt *FormatContext) Debug() int {
	return int(ctxt.debug)
}

//
// C-Variable: AVFormatContext::error_recognition
func (ctxt *FormatContext) ErrorRecognition() int {
	return int(ctxt.error_recognition)
}

//
// C-Variable: AVFormatContext::event_flags
func (ctxt *FormatContext) EventFlags() int {
	return int(ctxt.event_flags)
}

//
// C-Variable: AVFormatContext::flags
func (ctxt *FormatContext) Flags() int {
	return int(ctxt.flags)
}

//
// C-Variable: AVFormatContext::flush_packets
func (ctxt *FormatContext) FlushPackets() int {
	return int(ctxt.flush_packets)
}

//
// C-Variable: AVFormatContext::format_probesize
func (ctxt *FormatContext) FormatProbesize() int {
	return int(ctxt.format_probesize)
}

//
// C-Variable: AVFormatContext::fps_probe_size
func (ctxt *FormatContext) FpsProbeSize() int {
	return int(ctxt.fps_probe_size)
}

//
// C-Variable: AVFormatContext::io_repositioned
func (ctxt *FormatContext) IoRepositioned() int {
	return int(ctxt.io_repositioned)
}

//
// C-Variable: AVFormatContext::keylen
func (ctxt *FormatContext) Keylen() int {
	return int(ctxt.keylen)
}

//
// C-Variable: AVFormatContext::max_chunk_duration
func (ctxt *FormatContext) MaxChunkDuration() int {
	return int(ctxt.max_chunk_duration)
}

//
// C-Variable: AVFormatContext::max_chunk_size
func (ctxt *FormatContext) MaxChunkSize() int {
	return int(ctxt.max_chunk_size)
}

//
// C-Variable: AVFormatContext::max_delay
func (ctxt *FormatContext) MaxDelay() int {
	return int(ctxt.max_delay)
}

//
// C-Variable: AVFormatContext::max_ts_probe
func (ctxt *FormatContext) MaxTsProbe() int {
	return int(ctxt.max_ts_probe)
}

//
// C-Variable: AVFormatContext::seek2any
func (ctxt *FormatContext) Seek2any() int {
	return int(ctxt.seek2any)
}

//
// C-Variable: AVFormatContext::strict_std_compliance
func (ctxt *FormatContext) StrictStdCompliance() int {
	return int(ctxt.strict_std_compliance)
}

//
// C-Variable: AVFormatContext::ts_id
func (ctxt *FormatContext) TsId() int {
	return int(ctxt.ts_id)
}

//
// C-Variable: AVFormatContext::use_wallclock_as_timestamps
func (ctxt *FormatContext) UseWallclockAsTimestamps() int {
	return int(ctxt.use_wallclock_as_timestamps)
}

//
// C-Variable: AVFormatContext::duration
func (ctxt *FormatContext) Duration() int64 {
	return int64(ctxt.duration)
}

//
// C-Variable: AVFormatContext::max_analyze_duration
func (ctxt *FormatContext) MaxAnalyzeDuration2() int64 {
	return int64(ctxt.max_analyze_duration)
}

//
// C-Variable: AVFormatContext::max_interleave_delta
func (ctxt *FormatContext) MaxInterleaveDelta() int64 {
	return int64(ctxt.max_interleave_delta)
}

//
// C-Variable: AVFormatContext::output_ts_offset
func (ctxt *FormatContext) OutputTsOffset() int64 {
	return int64(ctxt.output_ts_offset)
}

//
// C-Variable: AVFormatContext::probesize
func (ctxt *FormatContext) Probesize2() int64 {
	return int64(ctxt.probesize)
}

//
// C-Variable: AVFormatContext::skip_initial_bytes
func (ctxt *FormatContext) SkipInitialBytes() int64 {
	return int64(ctxt.skip_initial_bytes)
}

//
// C-Variable: AVFormatContext::start_time
func (ctxt *FormatContext) StartTime() int64 {
	return int64(ctxt.start_time)
}

//
// C-Variable: AVFormatContext::start_time_realtime
func (ctxt *FormatContext) StartTimeRealtime() int64 {
	return int64(ctxt.start_time_realtime)
}

//
// C-Variable: AVFormatContext::iformat
func (ctxt *FormatContext) Iformat() *InputFormat {
	return (*InputFormat)(unsafe.Pointer(ctxt.iformat))
}

//
// C-Variable: AVFormatContext::oformat
func (ctxt *FormatContext) Oformat() *OutputFormat {
	return (*OutputFormat)(unsafe.Pointer(ctxt.oformat))
}

// //
// C-Variable: AVFormatContext::dump_separator
func (ctxt *FormatContext) DumpSeparator() uint8 {
	return uint8(*ctxt.dump_separator)
}

//
// C-Variable: AVFormatContext::correct_ts_overflow
func (ctxt *FormatContext) CorrectTsOverflow() int {
	return int(ctxt.correct_ts_overflow)
}

//
// C-Variable: AVFormatContext::max_index_size
func (ctxt *FormatContext) MaxIndexSize() uint {
	return uint(ctxt.max_index_size)
}

//
// C-Variable: AVFormatContext::max_picture_buffer
func (ctxt *FormatContext) MaxPictureBuffer() uint {
	return uint(ctxt.max_picture_buffer)
}

//
// C-Variable: AVFormatContext::nb_chapters
func (ctxt *FormatContext) NbChapters() uint {
	return uint(ctxt.nb_chapters)
}

//
// C-Variable: AVFormatContext::nb_programs
func (ctxt *FormatContext) NbPrograms() uint {
	return uint(ctxt.nb_programs)
}

//
// C-Variable: AVFormatContext::nb_streams
func (ctxt *FormatContext) NbStreams() uint {
	return uint(ctxt.nb_streams)
}

//
// C-Variable: AVFormatContext::packet_size
func (ctxt *FormatContext) PacketSize() uint {
	return uint(ctxt.packet_size)
}

//
// C-Variable: AVFormatContext::probesize
func (ctxt *FormatContext) Probesize() uint {
	return uint(ctxt.probesize)
}

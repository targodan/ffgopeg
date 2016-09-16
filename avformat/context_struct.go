// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avformat

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"
import "unsafe"

func (ctxt *FormatContext) Chapters() **Chapter {
	return (**Chapter)(unsafe.Pointer(ctxt.chapters))
}

func (ctxt *FormatContext) AudioCodec() *Codec {
	return (*Codec)(unsafe.Pointer(ctxt.audio_codec))
}

func (ctxt *FormatContext) SubtitleCodec() *Codec {
	return (*Codec)(unsafe.Pointer(ctxt.subtitle_codec))
}

func (ctxt *FormatContext) VideoCodec() *Codec {
	return (*Codec)(unsafe.Pointer(ctxt.video_codec))
}

func (ctxt *FormatContext) Metadata() *Dictionary {
	return (*Dictionary)(unsafe.Pointer(ctxt.metadata))
}

func (ctxt *FormatContext) Internal() *FormatInternal {
	return (*FormatInternal)(unsafe.Pointer(ctxt.internal))
}

func (ctxt *FormatContext) Pb() *IOContext {
	return (*IOContext)(unsafe.Pointer(ctxt.pb))
}

func (ctxt *FormatContext) InterruptCallback() IOInterruptCB {
	return IOInterruptCB(ctxt.interrupt_callback)
}

func (ctxt *FormatContext) Programs() **Program {
	return (**Program)(unsafe.Pointer(ctxt.programs))
}

func (ctxt *FormatContext) Streams() []*Stream {
	length := ctxt.NbStreams()
	return (*[1 << 30]*Stream)(unsafe.Pointer(*ctxt.streams))[:length:length]
}

func (ctxt *FormatContext) Filename() string {
	return C.GoString((*C.char)(unsafe.Pointer(&ctxt.filename[0])))
}

// func (ctxt *FormatContext) CodecWhitelist() string {
// 	return C.GoString(ctxt.codec_whitelist)
// }

// func (ctxt *FormatContext) FormatWhitelist() string {
// 	return C.GoString(ctxt.format_whitelist)
// }

func (ctxt *FormatContext) AudioCodecId() CodecId {
	return CodecId(ctxt.audio_codec_id)
}

func (ctxt *FormatContext) SubtitleCodecId() CodecId {
	return CodecId(ctxt.subtitle_codec_id)
}

func (ctxt *FormatContext) VideoCodecId() CodecId {
	return CodecId(ctxt.video_codec_id)
}

func (ctxt *FormatContext) DurationEstimationMethod() DurationEstimationMethod {
	return DurationEstimationMethod(ctxt.duration_estimation_method)
}

func (ctxt *FormatContext) AudioPreload() int {
	return int(ctxt.audio_preload)
}

func (ctxt *FormatContext) AvioFlags() int {
	return int(ctxt.avio_flags)
}

func (ctxt *FormatContext) AvoidNegativeTs() int {
	return int(ctxt.avoid_negative_ts)
}

func (ctxt *FormatContext) BitRate() int {
	return int(ctxt.bit_rate)
}

func (ctxt *FormatContext) CtxFlags() int {
	return int(ctxt.ctx_flags)
}

func (ctxt *FormatContext) Debug() int {
	return int(ctxt.debug)
}

func (ctxt *FormatContext) ErrorRecognition() int {
	return int(ctxt.error_recognition)
}

func (ctxt *FormatContext) EventFlags() int {
	return int(ctxt.event_flags)
}

func (ctxt *FormatContext) Flags() int {
	return int(ctxt.flags)
}

func (ctxt *FormatContext) FlushPackets() int {
	return int(ctxt.flush_packets)
}

func (ctxt *FormatContext) FormatProbesize() int {
	return int(ctxt.format_probesize)
}

func (ctxt *FormatContext) FpsProbeSize() int {
	return int(ctxt.fps_probe_size)
}

func (ctxt *FormatContext) IoRepositioned() int {
	return int(ctxt.io_repositioned)
}

func (ctxt *FormatContext) Keylen() int {
	return int(ctxt.keylen)
}

func (ctxt *FormatContext) MaxChunkDuration() int {
	return int(ctxt.max_chunk_duration)
}

func (ctxt *FormatContext) MaxChunkSize() int {
	return int(ctxt.max_chunk_size)
}

func (ctxt *FormatContext) MaxDelay() int {
	return int(ctxt.max_delay)
}

func (ctxt *FormatContext) MaxTsProbe() int {
	return int(ctxt.max_ts_probe)
}

func (ctxt *FormatContext) MetadataHeaderPadding() int {
	return int(ctxt.metadata_header_padding)
}

func (ctxt *FormatContext) ProbeScore() int {
	return int(ctxt.probe_score)
}

func (ctxt *FormatContext) Seek2any() int {
	return int(ctxt.seek2any)
}

func (ctxt *FormatContext) StrictStdCompliance() int {
	return int(ctxt.strict_std_compliance)
}

func (ctxt *FormatContext) TsId() int {
	return int(ctxt.ts_id)
}

func (ctxt *FormatContext) UseWallclockAsTimestamps() int {
	return int(ctxt.use_wallclock_as_timestamps)
}

func (ctxt *FormatContext) Duration() int64 {
	return int64(ctxt.duration)
}

func (ctxt *FormatContext) MaxAnalyzeDuration2() int64 {
	return int64(ctxt.max_analyze_duration)
}

func (ctxt *FormatContext) MaxInterleaveDelta() int64 {
	return int64(ctxt.max_interleave_delta)
}

func (ctxt *FormatContext) OutputTsOffset() int64 {
	return int64(ctxt.output_ts_offset)
}

func (ctxt *FormatContext) Probesize2() int64 {
	return int64(ctxt.probesize)
}

func (ctxt *FormatContext) SkipInitialBytes() int64 {
	return int64(ctxt.skip_initial_bytes)
}

func (ctxt *FormatContext) StartTime() int64 {
	return int64(ctxt.start_time)
}

func (ctxt *FormatContext) StartTimeRealtime() int64 {
	return int64(ctxt.start_time_realtime)
}

func (ctxt *FormatContext) Iformat() *InputFormat {
	return (*InputFormat)(unsafe.Pointer(ctxt.iformat))
}

func (ctxt *FormatContext) Oformat() *OutputFormat {
	return (*OutputFormat)(unsafe.Pointer(ctxt.oformat))
}

// func (ctxt *FormatContext) DumpSeparator() uint8 {
// 	return uint8(ctxt.dump_separator)
// }

func (ctxt *FormatContext) CorrectTsOverflow() int {
	return int(ctxt.correct_ts_overflow)
}

func (ctxt *FormatContext) MaxIndexSize() uint {
	return uint(ctxt.max_index_size)
}

func (ctxt *FormatContext) MaxPictureBuffer() uint {
	return uint(ctxt.max_picture_buffer)
}

func (ctxt *FormatContext) NbChapters() uint {
	return uint(ctxt.nb_chapters)
}

func (ctxt *FormatContext) NbPrograms() uint {
	return uint(ctxt.nb_programs)
}

func (ctxt *FormatContext) NbStreams() uint {
	return uint(ctxt.nb_streams)
}

func (ctxt *FormatContext) PacketSize() uint {
	return uint(ctxt.packet_size)
}

func (ctxt *FormatContext) Probesize() uint {
	return uint(ctxt.probesize)
}

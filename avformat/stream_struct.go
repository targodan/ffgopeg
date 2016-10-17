// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)
// Corbatto (luca@corbatto.de)

package avformat

//#cgo pkg-config: libavformat libavcodec
//#include <libavformat/avformat.h>
import "C"
import (
	"unsafe"

	"github.com/targodan/ffgopeg/avcodec"
)

//
// C-Variable: AVStream::codec
func (avs *Stream) CodecContext() *avcodec.CodecContext {
	return (*avcodec.CodecContext)(unsafe.Pointer(avs.codec))
}

//
// C-Variable: AVStream::metadata
func (avs *Stream) Metadata() *Dictionary {
	return (*Dictionary)(unsafe.Pointer(avs.metadata))
}

//
// C-Variable: AVStream::index_entries
func (avs *Stream) IndexEntries() *IndexEntry {
	return (*IndexEntry)(unsafe.Pointer(avs.index_entries))
}

//
// C-Variable: AVStream::attached_pic
func (avs *Stream) AttachedPic() Packet {
	return Packet(avs.attached_pic)
}

//
// C-Variable: AVStream::probe_data
func (avs *Stream) ProbeData() ProbeData {
	return ProbeData(avs.probe_data)
}

//
// C-Variable: AVStream::avg_frame_rate
func (avs *Stream) AvgFrameRate() Rational {
	return Rational(avs.avg_frame_rate)
}

//
// C-Variable: AVStream::r_frame_rate
func (avs *Stream) RFrameRate() Rational {
	return Rational(avs.r_frame_rate)
}

//
// C-Variable: AVStream::sample_aspect_ratio
func (avs *Stream) SampleAspectRatio() Rational {
	return Rational(avs.sample_aspect_ratio)
}

//
// C-Variable: AVStream::time_base
func (avs *Stream) TimeBase() Rational {
	return Rational(avs.time_base)
}

//
// C-Variable: AVStream::discard
func (avs *Stream) Discard() Discard {
	return Discard(avs.discard)
}

//
// C-Variable: AVStream::need_parsing
func (avs *Stream) NeedParsing() StreamParseType {
	return StreamParseType(avs.need_parsing)
}

//
// C-Variable: AVStream::codec_info_nb_frames
func (avs *Stream) CodecInfoNbFrames() int {
	return int(avs.codec_info_nb_frames)
}

//
// C-Variable: AVStream::disposition
func (avs *Stream) Disposition() int {
	return int(avs.disposition)
}

//
// C-Variable: AVStream::event_flags
func (avs *Stream) EventFlags() int {
	return int(avs.event_flags)
}

//
// C-Variable: AVStream::id
func (avs *Stream) Id() int {
	return int(avs.id)
}

//
// C-Variable: AVStream::index
func (avs *Stream) Index() int {
	return int(avs.index)
}

//
// C-Variable: AVStream::inject_global_side_data
func (avs *Stream) InjectGlobalSideData() int {
	return int(avs.inject_global_side_data)
}

//
// C-Variable: AVStream::last_
func (avs *Stream) LastIpDuration() int {
	return int(avs.last_IP_duration)
}

//
// C-Variable: AVStream::nb_decoded_frames
func (avs *Stream) NbDecodedFrames() int {
	return int(avs.nb_decoded_frames)
}

//
// C-Variable: AVStream::nb_index_entries
func (avs *Stream) NbIndexEntries() int {
	return int(avs.nb_index_entries)
}

//
// C-Variable: AVStream::nb_side_data
func (avs *Stream) NbSideData() int {
	return int(avs.nb_side_data)
}

//
// C-Variable: AVStream::probe_packets
func (avs *Stream) ProbePackets() int {
	return int(avs.probe_packets)
}

//
// C-Variable: AVStream::pts_wrap_behavior
func (avs *Stream) PtsWrapBehavior() int {
	return int(avs.pts_wrap_behavior)
}

//
// C-Variable: AVStream::request_probe
func (avs *Stream) RequestProbe() int {
	return int(avs.request_probe)
}

//
// C-Variable: AVStream::skip_samples
func (avs *Stream) SkipSamples() int {
	return int(avs.skip_samples)
}

//
// C-Variable: AVStream::skip_to_keyframe
func (avs *Stream) SkipToKeyframe() int {
	return int(avs.skip_to_keyframe)
}

//
// C-Variable: AVStream::stream_identifier
func (avs *Stream) StreamIdentifier() int {
	return int(avs.stream_identifier)
}

//
// C-Variable: AVStream::update_initial_durations_done
func (avs *Stream) UpdateInitialDurationsDone() int {
	return int(avs.update_initial_durations_done)
}

//
// C-Variable: AVStream::cur_dts
func (avs *Stream) CurDts() int64 {
	return int64(avs.cur_dts)
}

//
// C-Variable: AVStream::duration
func (avs *Stream) Duration() int64 {
	return int64(avs.duration)
}

//
// C-Variable: AVStream::first_dts
func (avs *Stream) FirstDts() int64 {
	return int64(avs.first_dts)
}

//
// C-Variable: AVStream::interleaver_chunk_duration
func (avs *Stream) InterleaverChunkDuration() int64 {
	return int64(avs.interleaver_chunk_duration)
}

//
// C-Variable: AVStream::interleaver_chunk_size
func (avs *Stream) InterleaverChunkSize() int64 {
	return int64(avs.interleaver_chunk_size)
}

//
// C-Variable: AVStream::last_dts_for_order_check
func (avs *Stream) LastDtsForOrderCheck() int64 {
	return int64(avs.last_dts_for_order_check)
}

//
// C-Variable: AVStream::last_
func (avs *Stream) LastIpPts() int64 {
	return int64(avs.last_IP_pts)
}

//
// C-Variable: AVStream::mux_ts_offset
func (avs *Stream) MuxTsOffset() int64 {
	return int64(avs.mux_ts_offset)
}

//
// C-Variable: AVStream::nb_frames
func (avs *Stream) NbFrames() int64 {
	return int64(avs.nb_frames)
}

//
// C-Variable: AVStream::pts_buffer
func (avs *Stream) PtsBuffer() int64 {
	return int64(avs.pts_buffer[0])
}

//
// C-Variable: AVStream::pts_reorder_error
func (avs *Stream) PtsReorderError() int64 {
	return int64(avs.pts_reorder_error[0])
}

//
// C-Variable: AVStream::pts_wrap_reference
func (avs *Stream) PtsWrapReference() int64 {
	return int64(avs.pts_wrap_reference)
}

//
// C-Variable: AVStream::start_time
func (avs *Stream) StartTime() int64 {
	return int64(avs.start_time)
}

//
// C-Variable: AVStream::parser
func (avs *Stream) Parser() *CodecParserContext {
	return (*CodecParserContext)(unsafe.Pointer(avs.parser))
}

//
// C-Variable: AVStream::last_in_packet_buffer
func (avs *Stream) LastInPacketBuffer() *PacketList {
	return (*PacketList)(unsafe.Pointer(avs.last_in_packet_buffer))
}

//
// C-Variable: AVStream::dts_misordered
func (avs *Stream) DtsMisordered() uint8 {
	return uint8(avs.dts_misordered)
}

//
// C-Variable: AVStream::dts_ordered
func (avs *Stream) DtsOrdered() uint8 {
	return uint8(avs.dts_ordered)
}

//
// C-Variable: AVStream::pts_reorder_error_count
func (avs *Stream) PtsReorderErrorCount() uint8 {
	return uint8(avs.pts_reorder_error_count[0])
}

//
// C-Variable: AVStream::index_entries_allocated_size
func (avs *Stream) IndexEntriesAllocatedSize() uint {
	return uint(avs.index_entries_allocated_size)
}

//
// C-Variable: AVStream::codecpar
func (avs *Stream) CodecPar() *avcodec.CodecParameters {
	return (*avcodec.CodecParameters)(unsafe.Pointer(avs.codecpar))
}

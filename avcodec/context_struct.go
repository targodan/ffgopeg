// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"
import (
	"github.com/targodan/goav/avutil"
)

// ActiveThreadType returns which multithreading methods are used by the codec.
//
// C-Field: AVCodecContext::active_thread_type
func (ctxt *CodecContext) ActiveThreadType() int {
	return int(ctxt.active_thread_type)
}

// BQuantFactor returns the qscale factor between IP and B-frames.
//
// C-Field: AVCodecContext::b_quant_factor
func (ctxt *CodecContext) BQuantFactor() float32 {
	return float32(ctxt.b_quant_factor)
}

// SetBQuantFactor sets the qscale factor between IP and B-frames.
// If > 0 then the last P-frame quantizer will be used (q = lastp_q*factor+offset).
// If < 0 then normal ratecontrol will be done (q = -normal_q*factor+offset)
//
// C-Field: AVCodecContext::b_quant_factor
func (ctxt *CodecContext) SetBQuantFactor(v float32) {
	ctxt.b_quant_factor = C.float(v)
}

// BQuantOffset return the qscale offset between IP and B-frames.
//
// C-Field: AVCodecContext::b_quant_offset
func (ctxt *CodecContext) BQuantOffset() float32 {
	return float32(ctxt.b_quant_offset)
}

// SetBQuantOffset return the qscale offset between IP and B-frames.
//
// C-Field: AVCodecContext::b_quant_offset
func (ctxt *CodecContext) SetBQuantOffset(v float32) {
	ctxt.b_quant_offset = C.float(v)
}

// BidirRefine returns something undocumented...
//
// C-Field: AVCodecContext::bidir_refine
func (ctxt *CodecContext) BidirRefine() int {
	return int(ctxt.bidir_refine)
}

// SetBidirRefine sets something undocumented...
//
// C-Field: AVCodecContext::bidir_refine
func (ctxt *CodecContext) SetBidirRefine(v int) {
	ctxt.bidir_refine = C.int(v)
}

// BitRate returns the average bit rate.
//
// C-Field: AVCodecContext::bit_rate
func (ctxt *CodecContext) BitRate() int {
	return int(ctxt.bit_rate)
}

// SetBitRate sets the average bit rate.
//
// C-Field: AVCodecContext::bit_rate
func (ctxt *CodecContext) SetBitRate(v int) {
	ctxt.bit_rate = C.int64_t(v)
}

// BitRateTolerance returns the number of bits the bitstream is allowed to diverge from the reference.
//
// C-Field: AVCodecContext::bit_rate_tolerance
func (ctxt *CodecContext) BitRateTolerance() int {
	return int(ctxt.bit_rate_tolerance)
}

// SetBitRateTolerance sets the number of bits the bitstream is allowed to diverge from the reference.
// The reference can be CBR (for CBR pass 1) or VBR (for VBR pass 2)
//
// C-Field: AVCodecContext::bit_rate_tolerance
func (ctxt *CodecContext) SetBitRateTolerance(v int) {
	ctxt.bit_rate_tolerance = C.int(v)
}

// BitsPerCodedSample returns the bits per sample/pixel from the demuxer (needed for huffyuv).
//
// C-Field: AVCodecContext::bits_per_coded_sample
func (ctxt *CodecContext) BitsPerCodedSample() int {
	return int(ctxt.bits_per_coded_sample)
}

// SetBitsPerCodedSample set the bits per sample/pixel from the demuxer (needed for huffyuv).
//
// C-Field: AVCodecContext::bits_per_coded_sample
func (ctxt *CodecContext) SetBitsPerCodedSample(v int) {
	ctxt.bits_per_coded_sample = C.int(v)
}

// BitsPerRawSample returns the bits per sample/pixel of internal libavcodec pixel/sample format.
//
// C-Field: AVCodecContext::bits_per_raw_sample
func (ctxt *CodecContext) BitsPerRawSample() int {
	return int(ctxt.bits_per_raw_sample)
}

// SetBitsPerRawSample sets the bits per sample/pixel of internal libavcodec pixel/sample format.
//
// C-Field: AVCodecContext::bits_per_raw_sample
func (ctxt *CodecContext) SetBitsPerRawSample(v int) {
	ctxt.bits_per_raw_sample = C.int(v)
}

// BlockAlign returns the number of bytes per packet if constant and known or 0.
// Used by some WAV based audio codecs.
//
// C-Field: AVCodecContext::block_align
func (ctxt *CodecContext) BlockAlign() int {
	return int(ctxt.block_align)
}

// Channels returns the number of audio channels.
//
// C-Field: AVCodecContext::channels
func (ctxt *CodecContext) Channels() int {
	return int(ctxt.channels)
}

// CodedHeight returns the bitstream height.
//
// C-Field: AVCodecContext::coded_height
func (ctxt *CodecContext) CodedHeight() int {
	return int(ctxt.coded_height)
}

// SetCodedHeight sets the bitstream height.
//
// C-Field: AVCodecContext::coded_height
func (ctxt *CodecContext) SetCodedHeight(v int) {
	ctxt.coded_height = C.int(v)
}

// CodedWidth returns the bitstream width.
// It may be different from the raw width. When the decoded frame is cropped before being output or lowres is enabled.
//
// C-Field: AVCodecContext::coded_width
func (ctxt *CodecContext) CodedWidth() int {
	return int(ctxt.coded_width)
}

// SetCodedWidth sets the bitstream width.
// It may be different from the raw width. When the decoded frame is cropped before being output or lowres is enabled.
//
// C-Field: AVCodecContext::coded_width
func (ctxt *CodecContext) SetCodedWidth(v int) {
	ctxt.coded_width = C.int(v)
}

// Height returns the raw image height.
//
// C-Field: AVCodecContext::height
func (ctxt *CodecContext) Height() int {
	return int(ctxt.height)
}

// SetHeight sets the raw image height.
//
// C-Field: AVCodecContext::coded_height
func (ctxt *CodecContext) SetHeight(v int) {
	ctxt.height = C.int(v)
}

// Width returns the raw image width.
//
// C-Field: AVCodecContext::width
func (ctxt *CodecContext) Width() int {
	return int(ctxt.width)
}

// SetWidth sets the raw image width.
//
// C-Field: AVCodecContext::width
func (ctxt *CodecContext) SetWidth(v int) {
	ctxt.width = C.int(v)
}

// CompressionLevel returns the compression level.
//
// C-Field: AVCodecContext::compression_level
func (ctxt *CodecContext) CompressionLevel() int {
	return int(ctxt.compression_level)
}

// SetCompressionLevel sets the compression level.
//
// C-Field: AVCodecContext::compression_level
func (ctxt *CodecContext) SetCompressionLevel(v int) {
	ctxt.compression_level = C.int(v)
}

// Cutoff returns the audio cutoff bandwidth (0 means "automatic")
//
// C-Field: AVCodecContext::cutoff
func (ctxt *CodecContext) Cutoff() int {
	return int(ctxt.cutoff)
}

// SetCutoff sets the audio cutoff bandwidth (0 means "automatic")
//
// C-Field: AVCodecContext::cutoff
func (ctxt *CodecContext) SetCutoff(v int) {
	ctxt.cutoff = C.int(v)
}

// DarkMasking returns the darkness masking (0 means "disabled")
//
// C-Field: AVCodecContext::dark_masking
func (ctxt *CodecContext) DarkMasking() float32 {
	return float32(ctxt.dark_masking)
}

// SetDarkMasking sets the darkness masking (0 means "disabled")
//
// C-Field: AVCodecContext::dark_masking
func (ctxt *CodecContext) SetDarkMasking(v float32) {
	ctxt.dark_masking = C.float(v)
}

// DctAlgo returns the DCT algorithm, see FF_DCT_*.
//
// C-Field: AVCodecContext::dct_algo
func (ctxt *CodecContext) DctAlgo() int {
	return int(ctxt.dct_algo)
}

// SetDctAlgo sets the DCT algorithm, see FF_DCT_*.
//
// C-Field: AVCodecContext::dct_algo
func (ctxt *CodecContext) SetDctAlgo(v int) {
	ctxt.dct_algo = C.int(v)
}

// Debug returns... the documentation only says "debug"
//
// C-Field: AVCodecContext::debug
func (ctxt *CodecContext) Debug() int {
	return int(ctxt.debug)
}

// SetDebug sets... the documentation only says "debug"
//
// C-Field: AVCodecContext::debug
func (ctxt *CodecContext) SetDebug(v int) {
	ctxt.debug = C.int(v)
}

// DebugMv returns... the documentation only says "debug"
// Code outside libavcodec should access this field using AVOptions.
//
// C-Field: AVCodecContext::debug_mv
func (ctxt *CodecContext) DebugMv() int {
	return int(ctxt.debug_mv)
}

// SetDebugMv sets... the documentation only says "debug"
// Code outside libavcodec should access this field using AVOptions.
//
// C-Field: AVCodecContext::debug_mv
func (ctxt *CodecContext) SetDebugMv(v int) {
	ctxt.debug_mv = C.int(v)
}

// Delay returns the codec delay.
//
// Encoding: Number of frames delay there will be from the encoder input to the decoder output. (we assume the decoder matches the spec) Decoding: Number of frames delay in addition to what a standard decoder as specified in the spec would produce.
//
// Video: Number of frames the decoded output will be delayed relative to the encoded input.
//
// Audio: For encoding, this field is unused (see initial_padding).
//
// For decoding, this is the number of samples the decoder needs to output before the decoder's output is valid. When seeking, you should start decoding this many samples prior to your desired seek point.
//
// C-Field: AVCodecContext::delay
func (ctxt *CodecContext) Delay() int {
	return int(ctxt.delay)
}

// SetDelay sets the codec delay.
//
// Encoding: Number of frames delay there will be from the encoder input to the decoder output. (we assume the decoder matches the spec) Decoding: Number of frames delay in addition to what a standard decoder as specified in the spec would produce.
//
// Video: Number of frames the decoded output will be delayed relative to the encoded input.
//
// Audio: For encoding, this field is unused (see initial_padding).
//
// For decoding, this is the number of samples the decoder needs to output before the decoder's output is valid. When seeking, you should start decoding this many samples prior to your desired seek point.
//
// C-Field: AVCodecContext::delay
func (ctxt *CodecContext) SetDelay(v int) {
	ctxt.delay = C.int(v)
}

// DiaSize returns ME diamod size and shape.
//
// C-Field: AVCodecContext::dia_size
func (ctxt *CodecContext) DiaSize() int {
	return int(ctxt.dia_size)
}

// SetDiaSize sets ME diamod size and shape.
//
// C-Field: AVCodecContext::dia_size
func (ctxt *CodecContext) SetDiaSize(v int) {
	ctxt.dia_size = C.int(v)
}

// ErrRecognition returns the error recognition.
// It may misdetect some more or less valid parts as errors.
//
// C-Field: AVCodecContext::err_recognition
func (ctxt *CodecContext) ErrRecognition() int {
	return int(ctxt.err_recognition)
}

// SetErrRecognition sets the error recognition.
// It may misdetect some more or less valid parts as errors.
//
// C-Field: AVCodecContext::err_recognition
func (ctxt *CodecContext) SetErrRecognition(v int) {
	ctxt.err_recognition = C.int(v)
}

// ErrorConcealment returns error concealment flags.
//
// C-Field: AVCodecContext::error_concealment
func (ctxt *CodecContext) ErrorConcealment() int {
	return int(ctxt.error_concealment)
}

// SetErrorConcealment sets error concealment flags.
//
// C-Field: AVCodecContext::error_concealment
func (ctxt *CodecContext) SetErrorConcealment(v int) {
	ctxt.error_concealment = C.int(v)
}

// ExtradataSize returns the extradata size.
//
// C-Field: AVCodecContext::extradata_size
func (ctxt *CodecContext) ExtradataSize() int {
	return int(ctxt.extradata_size)
}

// Flags returns the flags AV_CODEC_FLAG_*.
//
// C-Field: AVCodecContext::flags
func (ctxt *CodecContext) Flags() int {
	return int(ctxt.flags)
}

// SetFlags set the flags AV_CODEC_FLAG_*.
//
// C-Field: AVCodecContext::flags
func (ctxt *CodecContext) SetFlags() int {
	return int(ctxt.flags)
}

// Flags2 returns the flags AV_CODEC_FLAG2_*.
//
// C-Field: AVCodecContext::flags2
func (ctxt *CodecContext) Flags2() int {
	return int(ctxt.flags2)
}

// SetFlags2 sets the flags AV_CODEC_FLAG2_*.
//
// C-Field: AVCodecContext::flags2
func (ctxt *CodecContext) SetFlags2(v int) {
	ctxt.flags2 = C.int(v)
}

// FrameNumber returns the Frame counter, set by libavcodec.
//
// Decoding: Total number of frames returned from the decoder so far.
// Encoding: Total number of frames passed to the encoder so far.
//
// C-Field: AVCodecContext::frame_number
func (ctxt *CodecContext) FrameNumber() int {
	return int(ctxt.frame_number)
}

// FrameSize returns the number of samples per channel in an audio frame.
//
// Encoding: Set by libavcodec in avcodec_open2(). Each submitted frame except the last must contain exactly frame_size samples per channel. May be 0 when the codec has AV_CODEC_CAP_VARIABLE_FRAME_SIZE set, then the frame size is not restricted.
// Decoding: May be set by some decoders to indicate constant frame size.
//
// C-Field: AVCodecContext::frame_size
func (ctxt *CodecContext) FrameSize() int {
	return int(ctxt.frame_size)
}

// GlobalQuality returns the global quality for codecs which cannot change it per frame.
// This should be proportional to MPEG-1/2/4 qscale.
func (ctxt *CodecContext) GlobalQuality() int {
	return int(ctxt.global_quality)
}

// SetGlobalQuality sets the global quality for codecs which cannot change it per frame.
// This should be proportional to MPEG-1/2/4 qscale.
func (ctxt *CodecContext) SetGlobalQuality(v int) {
	ctxt.global_quality = C.int(v)
}

// GopSize returns the number of pictures in a group of pictures, or 0 for intra_only.
//
// C-Field: AVCodecContext::gop_size
func (ctxt *CodecContext) GopSize() int {
	return int(ctxt.gop_size)
}

// SetGopSize sets the number of pictures in a group of pictures, or 0 for intra_only.
//
// C-Field: AVCodecContext::gop_size
func (ctxt *CodecContext) SetGopSize(v int) {
	ctxt.gop_size = C.int(v)
}

// HasBFrames returns the size of the reordering buffer in the decoder.
//
// For MPEG-2 it is 1 IPB or 0 low delay IP.
//
// C-Field: AVCodecContext::has_b_frames
func (ctxt *CodecContext) HasBFrames() int {
	return int(ctxt.has_b_frames)
}

// IQuantFactor returns the qscale factor between P- and I-frames.
// If > 0 then the last P-frame quantizer will be used (q = lastp_q * factor + offset).
// If < 0 then normal ratecontrol will be done (q = -normal_q * factor + offset).
//
// C-Field: AVCodecContext::i_quant_factor
func (ctxt *CodecContext) IQuantFactor() float32 {
	return float32(ctxt.i_quant_factor)
}

// SetIQuantFactor sets the qscale factor between P- and I-frames.
// If > 0 then the last P-frame quantizer will be used (q = lastp_q * factor + offset).
// If < 0 then normal ratecontrol will be done (q = -normal_q * factor + offset).
//
// C-Field: AVCodecContext::i_quant_factor
func (ctxt *CodecContext) SetIQuantFactor(v float32) {
	ctxt.i_quant_factor = C.float(v)
}

// IQuantOffset returns the qscale offset between P and I-frames.
//
// C-Field: AVCodecContext::i_quant_offset
func (ctxt *CodecContext) IQuantOffset() float32 {
	return float32(ctxt.i_quant_offset)
}

// SetIQuantOffset sets the qscale offset between P and I-frames.
//
// C-Field: AVCodecContext::i_quant_offset
func (ctxt *CodecContext) SetIQuantOffset(v float32) {
	ctxt.i_quant_offset = C.float(v)
}

// IdctAlgo returns the IDCT algorithm, see FF_IDCT_*.
//
// C-Field: AVCodecContext::idct_algo
func (ctxt *CodecContext) IdctAlgo() int {
	return int(ctxt.idct_algo)
}

// SetIdctAlgo sets the IDCT algorithm, see FF_IDCT_*.
//
// C-Field: AVCodecContext::idct_algo
func (ctxt *CodecContext) SetIdctAlgo(v int) {
	ctxt.idct_algo = C.int(v)
}

// IldctCmp returns the interlaced DCT comparison function.
//
// C-Field: AVCodecContext::ildct_cmp
func (ctxt *CodecContext) IldctCmp() int {
	return int(ctxt.ildct_cmp)
}

// SetIldctCmp sets the interlaced DCT comparison function.
//
// C-Field: AVCodecContext::ildct_cmp
func (ctxt *CodecContext) SetIldctCmp(v int) {
	ctxt.ildct_cmp = C.int(v)
}

// IntraDcPrecision returns the precision of the intra DC coefficient - 8.
//
// C-Field: AVCodecContext::intra_dc_precision
func (ctxt *CodecContext) IntraDcPrecision() int {
	return int(ctxt.intra_dc_precision)
}

// SetIntraDcPrecision sets the precision of the intra DC coefficient - 8.
//
// C-Field: AVCodecContext::intra_dc_precision
func (ctxt *CodecContext) SetIntraDcPrecision(v int) {
	ctxt.intra_dc_precision = C.int(v)
}

// TODO: continue here

func (ctxt *CodecContext) KeyintMin() int {
	return int(ctxt.keyint_min)
}

func (ctxt *CodecContext) LastPredictorCount() int {
	return int(ctxt.last_predictor_count)
}

func (ctxt *CodecContext) Level() int {
	return int(ctxt.level)
}

func (ctxt *CodecContext) LogLevelOffset() int {
	return int(ctxt.log_level_offset)
}

func (ctxt *CodecContext) Lowres() int {
	return int(ctxt.lowres)
}

func (ctxt *CodecContext) LumiMasking() float64 {
	return float64(ctxt.lumi_masking)
}

func (ctxt *CodecContext) MaxBFrames() int {
	return int(ctxt.max_b_frames)
}

func (ctxt *CodecContext) MaxPredictionOrder() int {
	return int(ctxt.max_prediction_order)
}

func (ctxt *CodecContext) MaxQdiff() int {
	return int(ctxt.max_qdiff)
}

func (ctxt *CodecContext) MbCmp() int {
	return int(ctxt.mb_cmp)
}

func (ctxt *CodecContext) MbDecision() int {
	return int(ctxt.mb_decision)
}

func (ctxt *CodecContext) MbLmax() int {
	return int(ctxt.mb_lmax)
}

func (ctxt *CodecContext) MbLmin() int {
	return int(ctxt.mb_lmin)
}

func (ctxt *CodecContext) MeCmp() int {
	return int(ctxt.me_cmp)
}

func (ctxt *CodecContext) MePenaltyCompensation() int {
	return int(ctxt.me_penalty_compensation)
}

func (ctxt *CodecContext) MePreCmp() int {
	return int(ctxt.me_pre_cmp)
}

func (ctxt *CodecContext) MeRange() int {
	return int(ctxt.me_range)
}

func (ctxt *CodecContext) MeSubCmp() int {
	return int(ctxt.me_sub_cmp)
}

func (ctxt *CodecContext) MeSubpelQuality() int {
	return int(ctxt.me_subpel_quality)
}

func (ctxt *CodecContext) MinPredictionOrder() int {
	return int(ctxt.min_prediction_order)
}

func (ctxt *CodecContext) MiscBits() int {
	return int(ctxt.misc_bits)
}

func (ctxt *CodecContext) MpegQuant() int {
	return int(ctxt.mpeg_quant)
}

func (ctxt *CodecContext) Mv0Threshold() int {
	return int(ctxt.mv0_threshold)
}

func (ctxt *CodecContext) MvBits() int {
	return int(ctxt.mv_bits)
}

func (ctxt *CodecContext) NoiseReduction() int {
	return int(ctxt.noise_reduction)
}

func (ctxt *CodecContext) NsseWeight() int {
	return int(ctxt.nsse_weight)
}

func (ctxt *CodecContext) PCount() int {
	return int(ctxt.p_count)
}

func (ctxt *CodecContext) PMasking() float64 {
	return float64(ctxt.p_masking)
}

func (ctxt *CodecContext) PTexBits() int {
	return int(ctxt.p_tex_bits)
}

func (ctxt *CodecContext) PreDiaSize() int {
	return int(ctxt.pre_dia_size)
}

func (ctxt *CodecContext) PreMe() int {
	return int(ctxt.pre_me)
}

func (ctxt *CodecContext) PredictionMethod() int {
	return int(ctxt.prediction_method)
}

func (ctxt *CodecContext) Profile() int {
	return int(ctxt.profile)
}

func (ctxt *CodecContext) Qblur() float64 {
	return float64(ctxt.qblur)
}

func (ctxt *CodecContext) Qcompress() float64 {
	return float64(ctxt.qcompress)
}

func (ctxt *CodecContext) Qmax() int {
	return int(ctxt.qmax)
}

func (ctxt *CodecContext) Qmin() int {
	return int(ctxt.qmin)
}

func (ctxt *CodecContext) RcBufferSize() int {
	return int(ctxt.rc_buffer_size)
}

func (ctxt *CodecContext) RcInitialBufferOccupancy() int {
	return int(ctxt.rc_initial_buffer_occupancy)
}

func (ctxt *CodecContext) RcMaxAvailableVbvUse() float64 {
	return float64(ctxt.rc_max_available_vbv_use)
}

func (ctxt *CodecContext) RcMaxRate() int {
	return int(ctxt.rc_max_rate)
}

func (ctxt *CodecContext) RcMinRate() int {
	return int(ctxt.rc_min_rate)
}

func (ctxt *CodecContext) RcMinVbvOverflowUse() float64 {
	return float64(ctxt.rc_min_vbv_overflow_use)
}

func (ctxt *CodecContext) RcOverrideCount() int {
	return int(ctxt.rc_override_count)
}

func (ctxt *CodecContext) RefcountedFrames() int {
	return int(ctxt.refcounted_frames)
}

func (ctxt *CodecContext) Refs() int {
	return int(ctxt.refs)
}

func (ctxt *CodecContext) RtpPayloadSize() int {
	return int(ctxt.rtp_payload_size)
}

func (ctxt *CodecContext) SampleRate() int {
	return int(ctxt.sample_rate)
}

func (ctxt *CodecContext) ScenechangeThreshold() int {
	return int(ctxt.scenechange_threshold)
}

func (ctxt *CodecContext) SeekPreroll() int {
	return int(ctxt.seek_preroll)
}

func (ctxt *CodecContext) SideDataOnlyPackets() int {
	return int(ctxt.side_data_only_packets)
}

func (ctxt *CodecContext) SkipAlpha() int {
	return int(ctxt.skip_alpha)
}

func (ctxt *CodecContext) SkipBottom() int {
	return int(ctxt.skip_bottom)
}

func (ctxt *CodecContext) SkipCount() int {
	return int(ctxt.skip_count)
}

func (ctxt *CodecContext) SkipTop() int {
	return int(ctxt.skip_top)
}

func (ctxt *CodecContext) SliceCount() int {
	return int(ctxt.slice_count)
}

func (ctxt *CodecContext) SliceFlags() int {
	return int(ctxt.slice_flags)
}

func (ctxt *CodecContext) Slices() int {
	return int(ctxt.slices)
}

func (ctxt *CodecContext) SpatialCplxMasking() float64 {
	return float64(ctxt.spatial_cplx_masking)
}

func (ctxt *CodecContext) StrictStdCompliance() int {
	return int(ctxt.strict_std_compliance)
}

func (ctxt *CodecContext) SubCharencMode() int {
	return int(ctxt.sub_charenc_mode)
}

func (ctxt *CodecContext) SubtitleHeaderSize() int {
	return int(ctxt.subtitle_header_size)
}

func (ctxt *CodecContext) TemporalCplxMasking() float64 {
	return float64(ctxt.temporal_cplx_masking)
}

func (ctxt *CodecContext) ThreadCount() int {
	return int(ctxt.thread_count)
}

func (ctxt *CodecContext) ThreadSafeCallbacks() int {
	return int(ctxt.thread_safe_callbacks)
}

func (ctxt *CodecContext) ThreadType() int {
	return int(ctxt.thread_type)
}

func (ctxt *CodecContext) TicksPerFrame() int {
	return int(ctxt.ticks_per_frame)
}

func (ctxt *CodecContext) Trellis() int {
	return int(ctxt.trellis)
}

func (ctxt *CodecContext) WorkaroundBugs() int {
	return int(ctxt.workaround_bugs)
}

func (ctxt *CodecContext) AudioServiceType() AudioServiceType {
	return (AudioServiceType)(ctxt.audio_service_type)
}

func (ctxt *CodecContext) ChromaSampleLocation() ChromaLocation {
	return (ChromaLocation)(ctxt.chroma_sample_location)
}

func (ctxt *CodecContext) CodecDescriptor() *CodecDescriptor {
	return (*CodecDescriptor)(ctxt.codec_descriptor)
}

func (ctxt *CodecContext) CodecId() CodecId {
	return (CodecId)(ctxt.codec_id)
}

func (ctxt *CodecContext) CodecType() avutil.MediaType {
	return (avutil.MediaType)(ctxt.codec_type)
}

func (ctxt *CodecContext) ColorPrimaries() ColorPrimaries {
	return (ColorPrimaries)(ctxt.color_primaries)
}

func (ctxt *CodecContext) ColorRange() ColorRange {
	return (ColorRange)(ctxt.color_range)
}

func (ctxt *CodecContext) ColorTrc() ColorTransferCharacteristic {
	return (ColorTransferCharacteristic)(ctxt.color_trc)
}

func (ctxt *CodecContext) Colorspace() ColorSpace {
	return (ColorSpace)(ctxt.colorspace)
}

func (ctxt *CodecContext) FieldOrder() FieldOrder {
	return (FieldOrder)(ctxt.field_order)
}

func (ctxt *CodecContext) PixFmt() PixelFormat {
	return (PixelFormat)(ctxt.pix_fmt)
}

func (ctxt *CodecContext) RequestSampleFmt() SampleFormat {
	return (SampleFormat)(ctxt.request_sample_fmt)
}

func (ctxt *CodecContext) SetRequestSampleFmt(fmt SampleFormat) {
	ctxt.request_sample_fmt = C.enum_AVSampleFormat(fmt)
}

func (ctxt *CodecContext) SampleFmt() SampleFormat {
	return (SampleFormat)(ctxt.sample_fmt)
}

func (ctxt *CodecContext) SkipFrame() Discard {
	return (Discard)(ctxt.skip_frame)
}

func (ctxt *CodecContext) SkipIdct() Discard {
	return (Discard)(ctxt.skip_idct)
}

func (ctxt *CodecContext) SkipLoopFilter() Discard {
	return (Discard)(ctxt.skip_loop_filter)
}

// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)
// Corbatto (luca@corbatto.de)

package avcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"
import (
	"github.com/colek42/ffgopeg/avutil"
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

// KeyintMin returns the minimum GOP size.
//
// C-Field: AVCodecContext::keyint_min
func (ctxt *CodecContext) KeyintMin() int {
	return int(ctxt.keyint_min)
}

// SetKeyintMin sets the minimum GOP size.
//
// C-Field: AVCodecContext::keyint_min
func (ctxt *CodecContext) SetKeyintMin(v int) {
	ctxt.keyint_min = C.int(v)
}

// LastPredictorCount returns the amount of previous MV predictors (2a+1 x 2a+1)
//
// C-Field: AVCodecContext::last_predictor_count
func (ctxt *CodecContext) LastPredictorCount() int {
	return int(ctxt.last_predictor_count)
}

// SetLastPredictorCount sets the amount of previous MV predictors (2a+1 x 2a+1)
//
// C-Field: AVCodecContext::last_predictor_count
func (ctxt *CodecContext) SetLastPredictorCount(v int) {
	ctxt.last_predictor_count = C.int(v)
}

// Level returns the level.
//
// C-Field: AVCodecContext::level
func (ctxt *CodecContext) Level() int {
	return int(ctxt.level)
}

// SetLevel sets the level.
//
// C-Field: AVCodecContext::level
func (ctxt *CodecContext) SetLevel(v int) {
	ctxt.level = C.int(v)
}

// LogLevelOffset returns something undocumented...
//
// C-Field: AVCodecContext::log_level_offset
func (ctxt *CodecContext) LogLevelOffset() int {
	return int(ctxt.log_level_offset)
}

// SetLogLevelOffset sets something undocumented...
//
// C-Field: AVCodecContext::log_level_offset
func (ctxt *CodecContext) SetLogLevelOffset(v int) {
	ctxt.log_level_offset = C.int(v)
}

// LumiMasking retunrs the level of luminance masking.
// 0 means disabled.
//
// C-Field: AVCodecContext::lumi_masking
func (ctxt *CodecContext) LumiMasking() float32 {
	return float32(ctxt.lumi_masking)
}

// SetLumiMasking retunrs the level of luminance masking.
// 0 means disabled.
//
// C-Field: AVCodecContext::lumi_masking
func (ctxt *CodecContext) SetLumiMasking(v float32) {
	ctxt.lumi_masking = C.float(v)
}

// MaxBFrames returns the maximum nomber of B-frames between non-B-frames.
//
// Note: The output will be delayed by max_b_frames+1 relative to the input.
//
// C-Field: AVCodecContext::max_b_frames
func (ctxt *CodecContext) MaxBFrames() int {
	return int(ctxt.max_b_frames)
}

// SetMaxBFrames sets the maximum nomber of B-frames between non-B-frames.
//
// Note: The output will be delayed by max_b_frames+1 relative to the input.
//
// C-Field: AVCodecContext::max_b_frames
func (ctxt *CodecContext) SetMaxBFrames(v int) {
	ctxt.max_b_frames = C.int(v)
}

// MaxQDiff returns the maximum quantizer difference between frames.
//
// C-Field: AVCodecContext::max_qdiff
func (ctxt *CodecContext) MaxQDiff() int {
	return int(ctxt.max_qdiff)
}

// SetMaxQDiff sets the maximum quantizer difference between frames.
//
// C-Field: AVCodecContext::max_qdiff
func (ctxt *CodecContext) SetMaxQDiff(v int) {
	ctxt.max_qdiff = C.int(v)
}

// MbCmp returns the macroblock comparison function.
// (not supported yet)
//
// C-Field: AVCodecContext::mb_cmp
func (ctxt *CodecContext) MbCmp() int {
	return int(ctxt.mb_cmp)
}

// SetMbCmp sets the macroblock comparison function.
// (not supported yet)
//
// C-Field: AVCodecContext::mb_cmp
func (ctxt *CodecContext) SetMbCmp(v int) {
	ctxt.mb_cmp = C.int(v)
}

// MbDecision returns the macroblock decision mode.
//
// C-Field: AVCodecContext::mb_decision
func (ctxt *CodecContext) MbDecision() int {
	return int(ctxt.mb_decision)
}

// SetMbDecision sets the macroblock decision mode.
//
// C-Field: AVCodecContext::mb_decision
func (ctxt *CodecContext) SetMbDecision(v int) {
	ctxt.mb_decision = C.int(v)
}

// MbLMax returns the maximum MB Lagrange multiplier.
//
// C-Field: AVCodecContext::mb_lmax
func (ctxt *CodecContext) MbLMax() int {
	return int(ctxt.mb_lmax)
}

// SetMbLMax returns the maximum MB Lagrange multiplier.
//
// C-Field: AVCodecContext::mb_lmax
func (ctxt *CodecContext) SetMbLMax(v int) {
	ctxt.mb_lmax = C.int(v)
}

// MbLMin returns the minimum MB Lagrange multiplier.
//
// C-Field: AVCodecContext::mb_lmin
func (ctxt *CodecContext) MbLMin() int {
	return int(ctxt.mb_lmin)
}

// SetMbLMin returns the minimum MB Lagrange multiplier.
//
// C-Field: AVCodecContext::mb_lmin
func (ctxt *CodecContext) SetMbLMin(v int) {
	ctxt.mb_lmin = C.int(v)
}

// MeCmp returns the motion estimation comparison function.
//
// C-Field: AVCodecContext::me_cmp
func (ctxt *CodecContext) MeCmp() int {
	return int(ctxt.me_cmp)
}

// SetMeCmp sets the motion estimation comparison function.
//
// C-Field: AVCodecContext::me_cmp
func (ctxt *CodecContext) SetMeCmp(v int) {
	ctxt.me_cmp = C.int(v)
}

// MePreCmp returns the motion estimation prepass comparison function.
//
// C-Field: AVCodecContext::me_pre_cmp
func (ctxt *CodecContext) MePreCmp() int {
	return int(ctxt.me_pre_cmp)
}

// SetMePreCmp sets the motion estimation prepass comparison function.
//
// C-Field: AVCodecContext::me_pre_cmp
func (ctxt *CodecContext) SetMePreCmp(v int) {
	ctxt.me_pre_cmp = C.int(v)
}

// MeRange returns the maximum motion estimation search range in subpel units.
// If 0 then no limit.
//
// C-Field: AVCodecContext::me_range
func (ctxt *CodecContext) MeRange() int {
	return int(ctxt.me_range)
}

// SetMeRange sets the maximum motion estimation search range in subpel units.
// If 0 then no limit.
//
// C-Field: AVCodecContext::me_range
func (ctxt *CodecContext) SetMeRange(v int) {
	ctxt.me_range = C.int(v)
}

// MeSubCmp returns the subpixel motion comparison function.
//
// C-Field: AVCodecContext::me_sub_cmp
func (ctxt *CodecContext) MeSubCmp() int {
	return int(ctxt.me_sub_cmp)
}

// SetMeSubCmp sets the subpixel motion comparison function.
//
// C-Field: AVCodecContext::me_sub_cmp
func (ctxt *CodecContext) SetMeSubCmp(v int) {
	ctxt.me_sub_cmp = C.int(v)
}

// MeSubpelQuality returns the subpel ME quality.
//
// C-Field: AVCodecContext::me_subpel_quality
func (ctxt *CodecContext) MeSubpelQuality() int {
	return int(ctxt.me_subpel_quality)
}

// SetMeSubpelQuality sets the subpel ME quality.
//
// C-Field: AVCodecContext::me_subpel_quality
func (ctxt *CodecContext) SetMeSubpelQuality(v int) {
	ctxt.me_subpel_quality = C.int(v)
}

// Mv0Threshold returns the mv0 threshold.
// Value depends on the compare function used for follpel ME.
//
// C-Field: AVCodecContext::mv0_threshold
func (ctxt *CodecContext) Mv0Threshold() int {
	return int(ctxt.mv0_threshold)
}

// SetMv0Threshold sets the mv0 threshold.
// Value depends on the compare function used for follpel ME.
//
// C-Field: AVCodecContext::mv0_threshold
func (ctxt *CodecContext) SetMv0Threshold(v int) {
	ctxt.mv0_threshold = C.int(v)
}

// NsseWeight returns the noise vs sse weight for the nsse comparison function.
//
// C-Field: AVCodecContext::nsse_weight
func (ctxt *CodecContext) NsseWeight() int {
	return int(ctxt.nsse_weight)
}

// PMasking returns the spatial complexity masking. 0 equals disabled.
//
// C-Field: AVCodecContext::p_masking
func (ctxt *CodecContext) PMasking() float32 {
	return float32(ctxt.p_masking)
}

// SetPMasking sets the spatial complexity masking. 0 equals disabled.
//
// C-Field: AVCodecContext::p_masking
func (ctxt *CodecContext) SetPMasking(v float32) {
	ctxt.p_masking = C.float(v)
}

// PreDiaSize returns the ME prepass diamond size & shape.
//
// C-Field: AVCodecContext::pre_dia_size
func (ctxt *CodecContext) PreDiaSize() int {
	return int(ctxt.pre_dia_size)
}

// SetPreDiaSize sets the ME prepass diamond size & shape.
//
// C-Field: AVCodecContext::pre_dia_size
func (ctxt *CodecContext) SetPreDiaSize(v int) {
	ctxt.pre_dia_size = C.int(v)
}

// Profile returns somthing undocumented...
//
// C-Field: AVCodecContext::profile
func (ctxt *CodecContext) Profile() int {
	return int(ctxt.profile)
}

// SetProfile sets somthing undocumented...
//
// C-Field: AVCodecContext::profile
func (ctxt *CodecContext) SetProfile(v int) {
	ctxt.profile = C.int(v)
}

// QBlur returns the amount of qscale smoothing over time.
// (0.0 - 1.0)
//
// C-Field: AVCodecContext::qblur
func (ctxt *CodecContext) QBlur() float32 {
	return float32(ctxt.qblur)
}

// SetQBlur sets the amount of qscale smoothing over time.
// (0.0 - 1.0)
//
// C-Field: AVCodecContext::qblur
func (ctxt *CodecContext) SetQBlur(v float32) {
	ctxt.qblur = C.float(v)
}

// QCompress returns the amount of qscale change between easy & hard scenes. (0.0 - 1.0)
//
// C-Field: AVCodecContext::qcompress
func (ctxt *CodecContext) QCompress() float32 {
	return float32(ctxt.qcompress)
}

// SetQCompress sets the amount of qscale change between easy & hard scenes. (0.0 - 1.0)
//
// C-Field: AVCodecContext::qcompress
func (ctxt *CodecContext) SetQCompress(v float32) {
	ctxt.qcompress = C.float(v)
}

// QMax returns the maximum quantizer.
//
// C-Field: AVCodecContext::qmax
func (ctxt *CodecContext) QMax() int {
	return int(ctxt.qmax)
}

// SetQMax sets the maximum quantizer.
//
// C-Field: AVCodecContext::qmax
func (ctxt *CodecContext) SetQMax(v int) {
	ctxt.qmax = C.int(v)
}

// QMin returns the minimum quantizer.
//
// C-Field: AVCodecContext::qmin
func (ctxt *CodecContext) QMin() int {
	return int(ctxt.qmin)
}

// SetQMin sets the minimum quantizer.
//
// C-Field: AVCodecContext::qmin
func (ctxt *CodecContext) SetQMin(v int) {
	ctxt.qmin = C.int(v)
}

// RcBufferSize returns the decoder bitstream buffer size.
//
// C-Field: AVCodecContext::rc_buffer_size
func (ctxt *CodecContext) RcBufferSize() int {
	return int(ctxt.rc_buffer_size)
}

// SetRcBufferSize sets the decoder bitstream buffer size.
//
// C-Field: AVCodecContext::rc_buffer_size
func (ctxt *CodecContext) SetRcBufferSize(v int) {
	ctxt.rc_buffer_size = C.int(v)
}

// RcInitialBufferOccupancy returns the number of bits which should be loaded into the rx buffer before decoding starts.
//
// C-Field: AVCodecContext::rc_initial_buffer_occupancy
func (ctxt *CodecContext) RcInitialBufferOccupancy() int {
	return int(ctxt.rc_initial_buffer_occupancy)
}

// SetRcInitialBufferOccupancy sets the number of bits which should be loaded into the rx buffer before decoding starts.
//
// C-Field: AVCodecContext::rc_initial_buffer_occupancy
func (ctxt *CodecContext) SetRcInitialBufferOccupancy(v int) {
	ctxt.rc_initial_buffer_occupancy = C.int(v)
}

// RcMaxAvailableVbvUse returns the ratecontrol attempt to use, at maximum, of what can be used without an underflow.
//
// C-Field: AVCodecContext::rc_max_available_vbv_use
func (ctxt *CodecContext) RcMaxAvailableVbvUse() float32 {
	return float32(ctxt.rc_max_available_vbv_use)
}

// SetRcMaxAvailableVbvUse sets the ratecontrol attempt to use, at maximum, of what can be used without an underflow.
//
// C-Field: AVCodecContext::rc_max_available_vbv_use
func (ctxt *CodecContext) SetRcMaxAvailableVbvUse(v float32) {
	ctxt.rc_max_available_vbv_use = C.float(v)
}

// RcMaxRate returns the maximum bitrate.
//
// C-Field: AVCodecContext::rc_max_rate
func (ctxt *CodecContext) RcMaxRate() int64 {
	return int64(ctxt.rc_max_rate)
}

// SetRcMaxRate sets the maximum bitrate.
//
// C-Field: AVCodecContext::rc_max_rate
func (ctxt *CodecContext) SetRcMaxRate(v int64) {
	ctxt.rc_max_rate = C.int64_t(v)
}

// RcMinRate returns the minimum bitrate.
//
// C-Field: AVCodecContext::rc_min_rate
func (ctxt *CodecContext) RcMinRate() int64 {
	return int64(ctxt.rc_min_rate)
}

// SetRcMinRate sets the minimum bitrate.
//
// C-Field: AVCodecContext::rc_min_rate
func (ctxt *CodecContext) SetRcMinRate(v int64) {
	ctxt.rc_min_rate = C.int64_t(v)
}

// RcMinVbvOverflowUse returns the ratecontrol attempt to use, at least, times the amount needed to prevent a vbv overflow.
//
// C-Field: AVCodecContext::rc_min_vbv_overflow_use
func (ctxt *CodecContext) RcMinVbvOverflowUse() float32 {
	return float32(ctxt.rc_min_vbv_overflow_use)
}

// SetRcMinVbvOverflowUse sets the ratecontrol attempt to use, at least, times the amount needed to prevent a vbv overflow.
//
// C-Field: AVCodecContext::rc_min_vbv_overflow_use
func (ctxt *CodecContext) SetRcMinVbvOverflowUse(v float32) {
	ctxt.rc_min_vbv_overflow_use = C.float(v)
}

// RcOverrideCount returns ratecontrol override, see RcOverride.
//
// C-Field: AVCodecContext::rc_override_count
func (ctxt *CodecContext) RcOverrideCount() int {
	return int(ctxt.rc_override_count)
}

// SetRcOverrideCount sets ratecontrol override, see RcOverride.
//
// C-Field: AVCodecContext::rc_override_count
func (ctxt *CodecContext) SetRcOverrideCount(v int) {
	ctxt.rc_override_count = C.int(v)
}

// RefcountedFrames returns the number of references of audio and video frames.
//
// If non-zero, the decoded audio and video frames returned from avcodec_decode_video2() and avcodec_decode_audio4() are reference-counted and are valid indefinitely.
//
// The caller must free them with av_frame_unref() when they are not needed anymore. Otherwise, the decoded frames must not be freed by the caller and are only valid until the next decode call.
//
// This is always automatically enabled if avcodec_receive_frame() is used.
//
// C-Field: AVCodecContext::refcounted_frames
func (ctxt *CodecContext) RefcountedFrames() int {
	return int(ctxt.refcounted_frames)
}

// Refs returns the number reference frames.
//
// C-Field: AVCodecContext::refs
func (ctxt *CodecContext) Refs() int {
	return int(ctxt.refs)
}

// SetRefs sets the number reference frames.
//
// C-Field: AVCodecContext::refs
func (ctxt *CodecContext) SetRefs(v int) {
	ctxt.refs = C.int(v)
}

// SampleRate returns the number of samples per second.
//
// C-Field: AVCodecContext::sample_rate
func (ctxt *CodecContext) SampleRate() int {
	return int(ctxt.sample_rate)
}

// SetSampleRate sets the number of samples per second.
//
// C-Field: AVCodecContext::sample_rate
func (ctxt *CodecContext) SetSampleRate(sr int) {
	ctxt.sample_rate = C.int(sr)
}

// SkipBottom returns the number of macroblocks rows at the bottom which are skipped.
//
// C-Field: AVCodecContext::skip_bottom
func (ctxt *CodecContext) SkipBottom() int {
	return int(ctxt.skip_bottom)
}

// SetSkipBottom sets the number of macroblocks rows at the bottom which are skipped.
//
// C-Field: AVCodecContext::skip_bottom
func (ctxt *CodecContext) SetSkipBottom(v int) {
	ctxt.skip_bottom = C.int(v)
}

// SkipTop returns the number of macroblock rows at the top which are skipped.
//
// C-Field: AVCodecContext::skip_top
func (ctxt *CodecContext) SkipTop() int {
	return int(ctxt.skip_top)
}

// SetSkipTop sets the number of macroblock rows at the top which are skipped.
//
// C-Field: AVCodecContext::skip_top
func (ctxt *CodecContext) SetSkipTop(v int) {
	ctxt.skip_top = C.int(v)
}

// SliceCount returns the slice count.
//
// C-Field: AVCodecContext::slice_count
func (ctxt *CodecContext) SliceCount() int {
	return int(ctxt.slice_count)
}

// SetSliceCount sets the slice count.
//
// C-Field: AVCodecContext::slice_count
func (ctxt *CodecContext) SetSliceCount(v int) {
	ctxt.slice_count = C.int(v)
}

// SliceFlags returns the slice flags.
//
// C-Field: AVCodecContext::slice_flags
func (ctxt *CodecContext) SliceFlags() int {
	return int(ctxt.slice_flags)
}

// SetSliceFlags sets the slice flags.
//
// C-Field: AVCodecContext::slice_flags
func (ctxt *CodecContext) SetSliceFlags(v int) {
	ctxt.slice_flags = C.int(v)
}

// Slices returns the number of slices.
//
// Indicates number of picture subdicisions. Used for parallelized decoding.
//
// C-Field: AVCodecContext::slices
func (ctxt *CodecContext) Slices() int {
	return int(ctxt.slices)
}

// SetSlices sets the number of slices.
//
// Indicates number of picture subdicisions. Used for parallelized decoding.
//
// C-Field: AVCodecContext::slices
func (ctxt *CodecContext) SetSlices(v int) {
	ctxt.slices = C.int(v)
}

// SpatialCplxMasking returns the spatial complexity masking. 0 means disabled.
//
// C-Field: AVCodecContext::spatial_cplx_masking
func (ctxt *CodecContext) SpatialCplxMasking() float32 {
	return float32(ctxt.spatial_cplx_masking)
}

// SetSpatialCplxMasking sets the spatial complexity masking. 0 means disabled.
//
// C-Field: AVCodecContext::spatial_cplx_masking
func (ctxt *CodecContext) SetSpatialCplxMasking(v float32) {
	ctxt.spatial_cplx_masking = C.float(v)
}

// StrictStdCompliance returns the strictness to follow the standard (MPEG-4, ...).
//
// C-Field: AVCodecContext::strict_std_compliance
func (ctxt *CodecContext) StrictStdCompliance() int {
	return int(ctxt.strict_std_compliance)
}

// SetStrictStdCompliance sets the strictness to follow the standard (MPEG-4, ...).
//
// C-Field: AVCodecContext::strict_std_compliance
func (ctxt *CodecContext) SetStrictStdCompliance(v int) {
	ctxt.strict_std_compliance = C.int(v)
}

// SubCharencMode returns the subtitles character encoding mode.
//
// Formats or codecs might be adjusting this settings (if they are doing the conversion themselves for instance).
//
// C-Field: AVCodecContext::sub_charenc_mode
func (ctxt *CodecContext) SubCharencMode() int {
	return int(ctxt.sub_charenc_mode)
}

// SubtitleHeaderSize returns something undocumented...
//
// C-Field: AVCodecContext::subtitle_header_size
func (ctxt *CodecContext) SubtitleHeaderSize() int {
	return int(ctxt.subtitle_header_size)
}

// TemporalCplxMasking returns the temporal complexity masking. 0 means disabled.
//
// C-Field: AVCodecContext::temporal_cplx_masking
func (ctxt *CodecContext) TemporalCplxMasking() float32 {
	return float32(ctxt.temporal_cplx_masking)
}

// SetTemporalCplxMasking returns the temporal complexity masking. 0 means disabled.
//
// C-Field: AVCodecContext::temporal_cplx_masking
func (ctxt *CodecContext) SetTemporalCplxMasking(v float32) {
	ctxt.temporal_cplx_masking = C.float(v)
}

// ThreadCount returns the thread count, which is used to decide how many independent tasks should be passed to execute().
//
// C-Field: AVCodecContext::thread_count
func (ctxt *CodecContext) ThreadCount() int {
	return int(ctxt.thread_count)
}

// SetThreadCount sets the thread count, which is used to decide how many independent tasks should be passed to execute().
//
// C-Field: AVCodecContext::thread_count
func (ctxt *CodecContext) SetThreadCount(v int) {
	ctxt.thread_count = C.int(v)
}

// ThreadSafeCallbacks returns whether or not the custom get_buffer() callback can be called synchronously from another threa, which allows faster multithreaded decoding.
//
// C-Field: AVCodecContext::thread_safe_callbacks
func (ctxt *CodecContext) ThreadSafeCallbacks() bool {
	return int(ctxt.thread_safe_callbacks) != 0
}

// SetThreadSafeCallbacks sets whether or not the custom get_buffer() callback can be called synchronously from another threa, which allows faster multithreaded decoding.
//
// C-Field: AVCodecContext::thread_safe_callbacks
func (ctxt *CodecContext) SetThreadSafeCallbacks(v bool) {
	if v {
		ctxt.thread_safe_callbacks = C.int(1)
	} else {
		ctxt.thread_safe_callbacks = C.int(0)
	}
}

// ThreadType returns which multithreading methods to use.
//
// C-Field: AVCodecContext::thread_type
func (ctxt *CodecContext) ThreadType() int {
	return int(ctxt.thread_type)
}

// SetThreadType sets which multithreading methods to use.
//
// C-Field: AVCodecContext::thread_type
func (ctxt *CodecContext) SetThreadType(v int) {
	ctxt.thread_type = C.int(v)
}

// TicksPerFrame returns the ticks per frame.
//
// For some codecs, the time base is closer to the field rate than the frame rate.
//
// Most notably, H.264 and MPEG-2 specify time_base as half of frame duration if no telecine is used ...
//
// Set to time_base ticks per frame. Default 1, e.g., H.264/MPEG-2 set it to 2.
//
// C-Field: AVCodecContext::ticks_per_frame
func (ctxt *CodecContext) TicksPerFrame() int {
	return int(ctxt.ticks_per_frame)
}

// SetTicksPerFrame sets the ticks per frame.
//
// For some codecs, the time base is closer to the field rate than the frame rate.
//
// Most notably, H.264 and MPEG-2 specify time_base as half of frame duration if no telecine is used ...
//
// Set to time_base ticks per frame. Default 1, e.g., H.264/MPEG-2 set it to 2.
//
// C-Field: AVCodecContext::ticks_per_frame
func (ctxt *CodecContext) SetTicksPerFrame(v int) {
	ctxt.ticks_per_frame = C.int(v)
}

// Trellis returns the trellis RD quantization.
//
// C-Field: AVCodecContext::trellis
func (ctxt *CodecContext) Trellis() int {
	return int(ctxt.trellis)
}

// SetTrellis sets the trellis RD quantization.
//
// C-Field: AVCodecContext::trellis
func (ctxt *CodecContext) SetTrellis(v int) {
	ctxt.trellis = C.int(v)
}

// WorkaroundBugs returns which bufs to work around.
//
// C-Field: AVCodecContext::workaround_bugs
func (ctxt *CodecContext) WorkaroundBugs() int {
	return int(ctxt.workaround_bugs)
}

// SetWorkaroundBugs sets which bufs to work around.
//
// C-Field: AVCodecContext::workaround_bugs
func (ctxt *CodecContext) SetWorkaroundBugs(v int) {
	ctxt.workaround_bugs = C.int(v)
}

// AudioServiceType returns the type of service that the audio stream conveys.
//
// C-Field: AVCodecContext::audio_service_type
func (ctxt *CodecContext) AudioServiceType() AudioServiceType {
	return (AudioServiceType)(ctxt.audio_service_type)
}

// SetAudioServiceType sets the type of service that the audio stream conveys.
//
// C-Field: AVCodecContext::audio_service_type
func (ctxt *CodecContext) SetAudioServiceType(v AudioServiceType) {
	ctxt.audio_service_type = C.enum_AVAudioServiceType(v)
}

// ChromaSampleLocation returns the location of chroma samples.
//
// C-Field: AVCodecContext::chroma_sample_location
func (ctxt *CodecContext) ChromaSampleLocation() ChromaLocation {
	return (ChromaLocation)(ctxt.chroma_sample_location)
}

// SetChromaSampleLocation sets the location of chroma samples.
//
// C-Field: AVCodecContext::chroma_sample_location
func (ctxt *CodecContext) SetChromaSampleLocation(v ChromaLocation) {
	ctxt.chroma_sample_location = C.enum_AVChromaLocation(v)
}

// CodecID returns the codec id.
//
// C-Field: AVCodecContext::codec_id
func (ctxt *CodecContext) CodecID() CodecId {
	return (CodecId)(ctxt.codec_id)
}

// CodecType return the codec type.
//
// C-Field: AVCodecContext::codec_type
func (ctxt *CodecContext) CodecType() avutil.MediaType {
	return (avutil.MediaType)(ctxt.codec_type)
}

// ColorPrimaries returns the chromaticity coordinates of the source primaries.
//
// C-Field: AVCodecContext::color_primaries
func (ctxt *CodecContext) ColorPrimaries() ColorPrimaries {
	return (ColorPrimaries)(ctxt.color_primaries)
}

// SetColorPrimaries sets the chromaticity coordinates of the source primaries.
//
// C-Field: AVCodecContext::color_primaries
func (ctxt *CodecContext) SetColorPrimaries(v ColorPrimaries) {
	ctxt.color_primaries = C.enum_AVColorPrimaries(v)
}

// ColorRange returns the MPEG vs JPEG YUV range.
//
// C-Field: AVCodecContext::color_range
func (ctxt *CodecContext) ColorRange() ColorRange {
	return (ColorRange)(ctxt.color_range)
}

// SetColorRange sets the MPEG vs JPEG YUV range.
//
// C-Field: AVCodecContext::color_range
func (ctxt *CodecContext) SetColorRange(v ColorRange) {
	ctxt.color_range = C.enum_AVColorRange(v)
}

// ColorTrc returns the color transfer characteristic.
//
// C-Field: AVCodecContext::color_trc
func (ctxt *CodecContext) ColorTrc() ColorTransferCharacteristic {
	return (ColorTransferCharacteristic)(ctxt.color_trc)
}

// SetColorTrc sets the color transfer characteristic.
//
// C-Field: AVCodecContext::color_trc
func (ctxt *CodecContext) SetColorTrc(v ColorTransferCharacteristic) {
	ctxt.color_trc = C.enum_AVColorTransferCharacteristic(v)
}

// Colorspace returns the YUV colorspace type.
//
// C-Field: AVCodecContext::colorspace
func (ctxt *CodecContext) Colorspace() ColorSpace {
	return (ColorSpace)(ctxt.colorspace)
}

// SetColorspace sets the YUV colorspace type.
//
// C-Field: AVCodecContext::colorspace
func (ctxt *CodecContext) SetColorspace(v ColorSpace) {
	ctxt.colorspace = C.enum_AVColorSpace(v)
}

// FieldOrder returns the field order.
//
// C-Field: AVCodecContext::field_order
func (ctxt *CodecContext) FieldOrder() FieldOrder {
	return (FieldOrder)(ctxt.field_order)
}

// SetFieldOrder sets the field order.
//
// C-Field: AVCodecContext::field_order
func (ctxt *CodecContext) SetFieldOrder(v FieldOrder) {
	ctxt.field_order = C.enum_AVFieldOrder(v)
}

// PixFmt returns the pixel format, see AV_PIX_FMT_xxx.
//
// C-Field: AVCodecContext::pix_fmt
func (ctxt *CodecContext) PixFmt() avutil.PixelFormat {
	return (avutil.PixelFormat)(ctxt.pix_fmt)
}

// SetPixFmt sets the pixel format, see AV_PIX_FMT_xxx.
//
// C-Field: AVCodecContext::pix_fmt
func (ctxt *CodecContext) SetPixFmt(v avutil.PixelFormat) {
	ctxt.pix_fmt = C.enum_AVPixelFormat(v)
}

// RequestSampleFmt returns the desired sample format.
//
// C-Field: AVCodecContext::request_sample_fmt
func (ctxt *CodecContext) RequestSampleFmt() avutil.SampleFormat {
	return (avutil.SampleFormat)(ctxt.request_sample_fmt)
}

// SetRequestSampleFmt sets the desired sample format.
//
// C-Field: AVCodecContext::request_sample_fmt
func (ctxt *CodecContext) SetRequestSampleFmt(fmt avutil.SampleFormat) {
	ctxt.request_sample_fmt = C.enum_AVSampleFormat(fmt)
}

// SampleFmt returns the audio sample format.
//
// C-Field: AVCodecContext::sample_fmt
func (ctxt *CodecContext) SampleFmt() avutil.SampleFormat {
	return (avutil.SampleFormat)(ctxt.sample_fmt)
}

// SetSampleFmt sets the audio sample format.
//
// C-Field: AVCodecContext::sample_fmt
func (ctxt *CodecContext) SetSampleFmt(v avutil.SampleFormat) {
	ctxt.sample_fmt = C.enum_AVSampleFormat(v)
}

// SkipFrame returns the skip decoding for selected frames.
//
// C-Field: AVCodecContext::skip_frame
func (ctxt *CodecContext) SkipFrame() Discard {
	return (Discard)(ctxt.skip_frame)
}

// SetSkipFrame sets the skip decoding for selected frames.
//
// C-Field: AVCodecContext::skip_frame
func (ctxt *CodecContext) SetSkipFrame(v Discard) {
	ctxt.skip_frame = C.enum_AVDiscard(v)
}

// SkipIdct returns the skip IDCT/dequantization for selected frames.
//
// C-Field: AVCodecContext::skip_idct
func (ctxt *CodecContext) SkipIdct() Discard {
	return (Discard)(ctxt.skip_idct)
}

// SetSkipIdct sets the skip IDCT/dequantization for selected frames.
//
// C-Field: AVCodecContext::skip_idct
func (ctxt *CodecContext) SetSkipIdct(v Discard) {
	ctxt.skip_idct = C.enum_AVDiscard(v)
}

// SkipLoopFilter returns the skip loop filtering for selected frames.
//
// C-Field: AVCodecContext::skip_loop_filter
func (ctxt *CodecContext) SkipLoopFilter() Discard {
	return (Discard)(ctxt.skip_loop_filter)
}

// SetSkipLoopFilter sets the skip loop filtering for selected frames.
//
// C-Field: AVCodecContext::skip_loop_filter
func (ctxt *CodecContext) SetSkipLoopFilter(v Discard) {
	ctxt.skip_loop_filter = C.enum_AVDiscard(v)
}

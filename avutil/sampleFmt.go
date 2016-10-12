// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)
// Corbatto (luca@corbatto.de)

package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/samplefmt.h>
import "C"

type SampleFormat C.enum_AVSampleFormat

const (
	AV_SAMPLE_FMT_NONE = SampleFormat(C.AV_SAMPLE_FMT_NONE)
	AV_SAMPLE_FMT_U8   = SampleFormat(C.AV_SAMPLE_FMT_U8)
	AV_SAMPLE_FMT_S16  = SampleFormat(C.AV_SAMPLE_FMT_S16)
	AV_SAMPLE_FMT_S32  = SampleFormat(C.AV_SAMPLE_FMT_S32)
	AV_SAMPLE_FMT_FLT  = SampleFormat(C.AV_SAMPLE_FMT_FLT)
	AV_SAMPLE_FMT_DBL  = SampleFormat(C.AV_SAMPLE_FMT_DBL)
	AV_SAMPLE_FMT_U8P  = SampleFormat(C.AV_SAMPLE_FMT_U8P)
	AV_SAMPLE_FMT_S16P = SampleFormat(C.AV_SAMPLE_FMT_S16P)
	AV_SAMPLE_FMT_S32P = SampleFormat(C.AV_SAMPLE_FMT_S32P)
	AV_SAMPLE_FMT_FLTP = SampleFormat(C.AV_SAMPLE_FMT_FLTP)
	AV_SAMPLE_FMT_DBLP = SampleFormat(C.AV_SAMPLE_FMT_DBLP)
	AV_SAMPLE_FMT_NB   = SampleFormat(C.AV_SAMPLE_FMT_NB)
)

// Alternative returns the planar<->packed alternative form of the given sample format, or AV_SAMPLE_FMT_NONE on error.
//
// C-Function: av_get_alt_sample_fmt
func (s SampleFormat) Alternative(planar bool) SampleFormat {
	p := 0
	if planar {
		p = 1
	}
	return SampleFormat(C.av_get_alt_sample_fmt(C.enum_AVSampleFormat(s), C.int(p)))
}

// Packed returns the packed variant.
//
// C-Function: av_get_packed_sample_fmt
func (s SampleFormat) Packed() SampleFormat {
	return SampleFormat(C.av_get_packed_sample_fmt(C.enum_AVSampleFormat(s)))
}

// Planar returns the planar variant.
//
// C-Function: av_get_planar_sample_fmt
func (s SampleFormat) Planar() SampleFormat {
	return SampleFormat(C.av_get_planar_sample_fmt(C.enum_AVSampleFormat(s)))
}

// Name returns the sample format name.
//
// C-Function: av_get_sample_fmt_name
func (s SampleFormat) Name() string {
	return C.GoString(C.av_get_sample_fmt_name(C.enum_AVSampleFormat(s)))
}

// BytesPerSample returns the bytes per sample.
//
// C-Function: av_get_bytes_per_sample
func (s SampleFormat) BytesPerSample() int {
	return int(C.av_get_bytes_per_sample(C.enum_AVSampleFormat(s)))
}

// IsPlanar returns true if the format is planar.
//
// C-Function: av_get_bytes_per_sample
func (s SampleFormat) IsPlanar() bool {
	return int(C.av_sample_fmt_is_planar(C.enum_AVSampleFormat(s))) != 0
}

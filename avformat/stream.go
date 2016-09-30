// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)
// Dmitry Patsura <talk@dmtry.me> https://github.com/ovr
// Corbatto (luca@corbatto.de)

package avformat

//#cgo pkg-config: libavformat
//#cgo pkg-config: libavutil
//#include <libavformat/avformat.h>
//#include <libavutil/display.h>
import "C"
import (
	"strconv"
	"strings"
	"unsafe"
)

// StreamGetRFrameRate does something undocumented.
//
// C-Function: av_stream_get_r_frame_rate
func (s *Stream) StreamGetRFrameRate() Rational {
	return (Rational)(C.av_stream_get_r_frame_rate((*C.struct_AVStream)(s)))
}

// StreamSetRFrameRate does something undocumented.
//
// C-Function: av_stream_set_r_frame_rate
func (s *Stream) StreamSetRFrameRate(r Rational) {
	C.av_stream_set_r_frame_rate((*C.struct_AVStream)(s), (C.struct_AVRational)(r))
}

// GetRotation does something undocumented.
//
// C-Function: av_display_rotation_get
func (s *Stream) GetRotation() int64 {
	dictionaryEntry := s.Metadata().Get("rotate")
	if dictionaryEntry != nil {
		value := dictionaryEntry.Value()
		strings.TrimRight(value, "%")

		rotation, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			return rotation
		}
	}

	var displaymatrix = C.av_stream_get_side_data((*C.struct_AVStream)(s), C.AV_PKT_DATA_DISPLAYMATRIX, nil)
	if displaymatrix != nil {
		return -int64(C.av_display_rotation_get((*C.int32_t)(unsafe.Pointer(displaymatrix))))
	}

	return 0
}

// StreamGetParser does something undocumented.
//
// C-Function: av_stream_get_parser
func (s *Stream) StreamGetParser() *CodecParserContext {
	return (*CodecParserContext)(C.av_stream_get_parser((*C.struct_AVStream)(s)))
}

// StreamGetEndPts returns the pts of the last muxed packet + its duration.
//
// C-Function: av_stream_get_end_pts
func (s *Stream) StreamGetEndPts() int64 {
	return int64(C.av_stream_get_end_pts((*C.struct_AVStream)(s)))
}

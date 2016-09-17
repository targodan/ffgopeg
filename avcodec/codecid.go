// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)
// Corbatto (luca@corbatto.de)

package avcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"
import "github.com/targodan/goav/avutil"

// BitsPerSample returns codec bits per sample.
// C-Function: av_get_bits_per_sample
func (c CodecId) BitsPerSample() int {
	return int(C.av_get_bits_per_sample((C.enum_AVCodecID)(c)))
}

// ExactBitsPerSample returns codec bits per sample.
// C-Function: av_get_exact_bits_per_sample
func (c CodecId) ExactBitsPerSample() int {
	return int(C.av_get_exact_bits_per_sample((C.enum_AVCodecID)(c)))
}

// Type returns the type of the codec.
// C-Function: avcodec_get_type
func (c CodecId) Type() avutil.MediaType {
	return (avutil.MediaType)(C.avcodec_get_type((C.enum_AVCodecID)(c)))
}

// Name returns the name of the codec.
// C-Function: avcodec_get_name
func (c CodecId) Name() string {
	return C.GoString(C.avcodec_get_name((C.enum_AVCodecID)(c)))
}

// Descriptor returns the CodecDescriptor of the codec.
// C-Function: avcodec_descriptor_get
func (c CodecId) Descriptor() *CodecDescriptor {
	return (*CodecDescriptor)(C.avcodec_descriptor_get((C.enum_AVCodecID)(c)))
}

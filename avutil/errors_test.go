// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Corbatto (luca@corbatto.de)

package avutil

import (
	"fmt"
	"testing"
)

func TestPrintErrors(t *testing.T) {
	fmt.Printf("%s: %d | %s\n", "AVERROR_BSF_NOT_FOUND", AVERROR_BSF_NOT_FOUND(), Strerror(AVERROR_BSF_NOT_FOUND()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_BUG", AVERROR_BUG(), Strerror(AVERROR_BUG()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_BUFFER_TOO_SMALL", AVERROR_BUFFER_TOO_SMALL(), Strerror(AVERROR_BUFFER_TOO_SMALL()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_DECODER_NOT_FOUND", AVERROR_DECODER_NOT_FOUND(), Strerror(AVERROR_DECODER_NOT_FOUND()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_DEMUXER_NOT_FOUND", AVERROR_DEMUXER_NOT_FOUND(), Strerror(AVERROR_DEMUXER_NOT_FOUND()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_ENCODER_NOT_FOUND", AVERROR_ENCODER_NOT_FOUND(), Strerror(AVERROR_ENCODER_NOT_FOUND()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_EOF", AVERROR_EOF(), Strerror(AVERROR_EOF()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_EXIT", AVERROR_EXIT(), Strerror(AVERROR_EXIT()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_EXTERNAL", AVERROR_EXTERNAL(), Strerror(AVERROR_EXTERNAL()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_FILTER_NOT_FOUND", AVERROR_FILTER_NOT_FOUND(), Strerror(AVERROR_FILTER_NOT_FOUND()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_INVALIDDATA", AVERROR_INVALIDDATA(), Strerror(AVERROR_INVALIDDATA()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_MUXER_NOT_FOUND", AVERROR_MUXER_NOT_FOUND(), Strerror(AVERROR_MUXER_NOT_FOUND()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_OPTION_NOT_FOUND", AVERROR_OPTION_NOT_FOUND(), Strerror(AVERROR_OPTION_NOT_FOUND()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_PATCHWELCOME", AVERROR_PATCHWELCOME(), Strerror(AVERROR_PATCHWELCOME()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_PROTOCOL_NOT_FOUND", AVERROR_PROTOCOL_NOT_FOUND(), Strerror(AVERROR_PROTOCOL_NOT_FOUND()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_STREAM_NOT_FOUND", AVERROR_STREAM_NOT_FOUND(), Strerror(AVERROR_STREAM_NOT_FOUND()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_UNKNOWN", AVERROR_UNKNOWN(), Strerror(AVERROR_UNKNOWN()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_EXPERIMENTAL", AVERROR_EXPERIMENTAL(), Strerror(AVERROR_EXPERIMENTAL()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_INPUT_CHANGED", AVERROR_INPUT_CHANGED(), Strerror(AVERROR_INPUT_CHANGED()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_OUTPUT_CHANGED", AVERROR_OUTPUT_CHANGED(), Strerror(AVERROR_OUTPUT_CHANGED()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_HTTP_BAD_REQUEST", AVERROR_HTTP_BAD_REQUEST(), Strerror(AVERROR_HTTP_BAD_REQUEST()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_HTTP_UNAUTHORIZED", AVERROR_HTTP_UNAUTHORIZED(), Strerror(AVERROR_HTTP_UNAUTHORIZED()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_HTTP_FORBIDDEN", AVERROR_HTTP_FORBIDDEN(), Strerror(AVERROR_HTTP_FORBIDDEN()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_HTTP_NOT_FOUND", AVERROR_HTTP_NOT_FOUND(), Strerror(AVERROR_HTTP_NOT_FOUND()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_HTTP_OTHER_4XX", AVERROR_HTTP_OTHER_4XX(), Strerror(AVERROR_HTTP_OTHER_4XX()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_HTTP_SERVER_ERROR", AVERROR_HTTP_SERVER_ERROR(), Strerror(AVERROR_HTTP_SERVER_ERROR()))
	fmt.Printf("%s: %d\n", "AV_ERROR_MAX_STRING_SIZE", AV_ERROR_MAX_STRING_SIZE())
}

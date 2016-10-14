package main

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math"
	"os"

	"github.com/targodan/native"

	"gopkg.in/targodan/ffgopeg.v1/avcodec"
	"gopkg.in/targodan/ffgopeg.v1/avformat"
	"gopkg.in/targodan/ffgopeg.v1/avutil"
)

const rawOutOnPlanar = true

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func panicOnCode(err avutil.ReturnCode) {
	if !err.Ok() {
		panic(err)
	}
}

func panicOnNil(err interface{}, msg string) {
	if err == nil {
		panic(msg)
	}
}

func findFirstAudioStream(ctxt *avformat.FormatContext) (int, error) {
	for i, s := range ctxt.Streams() {
		// Use the first audio stream we can find.
		// NOTE: There may be more than one, depending on the file.
		if s.CodecPar().CodecType() == avutil.AVMEDIA_TYPE_AUDIO {
			return i, nil
		}
	}
	return -1, errors.New("Could not find an audio stream.")
}

func printStreamInformation(codec *avcodec.Codec, codecCtxt *avcodec.CodecContext, audioStreamIndex int) {
	fmt.Fprintf(os.Stderr, "Codec: %s\n", codec.LongName())
	fmt.Fprint(os.Stderr, "Supported sample formats: ")
	for _, f := range codec.SampleFmts() {
		fmt.Fprintf(os.Stderr, "%s, ", f.Name())
	}
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "----------\n")
	fmt.Fprintf(os.Stderr, "%15s: %7d\n", "Stream", audioStreamIndex)
	fmt.Fprintf(os.Stderr, "%15s: %7s\n", "Sample Format", codecCtxt.SampleFmt().Name())
	fmt.Fprintf(os.Stderr, "%15s: %7d\n", "Sample rate", codecCtxt.SampleRate())
	fmt.Fprintf(os.Stderr, "%15s: %7d\n", "Sample size", codecCtxt.SampleFmt().BytesPerSample())
	fmt.Fprintf(os.Stderr, "%15s: %7v\n", "Planar", codecCtxt.SampleFmt().IsPlanar())
	fmt.Fprintf(os.Stderr, "%15s: %7d\n", "Channels", codecCtxt.Channels())
	fmt.Fprintf(os.Stderr, "%15s: %7v\n", "Float output", !rawOutOnPlanar || codecCtxt.SampleFmt().IsPlanar())
}

func receiveAndHandle(codecCtxt *avcodec.CodecContext, frame *avutil.Frame, out io.Writer) (err avutil.ReturnCode) {
	// Read the packets from the decoder.
	// NOTE: Each packet may generate more than one frame, depending on the codec.
	err = codecCtxt.ReceiveFrame(frame)
	for err.Ok() {
		// Let's handle the frame in a function.
		handleFrame(codecCtxt, frame, out)
		// Free any buffers and reset the fields to default values.
		frame.Unref()
		err = codecCtxt.ReceiveFrame(frame)
	}
	return
}

func handleFrame(codecCtxt *avcodec.CodecContext, frame *avutil.Frame, out io.Writer) {
	if codecCtxt.SampleFmt().IsPlanar() {
		// This means that the data of each channel is in its own buffer.
		// => frame->extended_data[i] contains data for the i-th channel.
		for s := 0; s < frame.NbSamples(); s++ {
			for c := 0; c < codecCtxt.Channels(); c++ {
				sample := getSample(codecCtxt, frame.ExtendedData(c, frame.Linesize(0)), s)
				binary.Write(out, binary.LittleEndian, sample)
			}
		}
	} else {
		// This means that the data of each channel is in the same buffer.
		// => frame->extended_data[0] contains data of all channels.
		if rawOutOnPlanar {
			out.Write(frame.ExtendedData(0, frame.Linesize(0)))
		} else {
			for s := 0; s < frame.NbSamples(); s++ {
				for c := 0; c < codecCtxt.Channels(); c++ {
					sample := getSample(codecCtxt, frame.ExtendedData(0, frame.Linesize(0)), s*codecCtxt.Channels()+c)
					binary.Write(out, binary.LittleEndian, sample)
				}
			}
		}
	}
}

func getSample(codecCtxt *avcodec.CodecContext, buffer []byte, sampleIndex int) float32 {
	sampleSize := codecCtxt.SampleFmt().BytesPerSample()
	byteIndex := sampleSize * sampleIndex
	var val int64
	switch sampleSize {
	case 1:
		// 8bit samples are always unsigned
		val = int64(buffer[byteIndex]) - 127

	case 2:
		val = int64(int16(native.ByteOrder.Uint16(buffer[byteIndex : byteIndex+sampleSize])))

	case 4:
		val = int64(int32(native.ByteOrder.Uint32(buffer[byteIndex : byteIndex+sampleSize])))

	case 8:
		val = int64(native.ByteOrder.Uint64(buffer[byteIndex : byteIndex+sampleSize]))

	default:
		panic(fmt.Sprintf("Invalid sample size %d.", sampleSize))
	}

	var ret float32
	switch codecCtxt.SampleFmt() {
	case avutil.AV_SAMPLE_FMT_U8,
		avutil.AV_SAMPLE_FMT_S16,
		avutil.AV_SAMPLE_FMT_S32,
		avutil.AV_SAMPLE_FMT_U8P,
		avutil.AV_SAMPLE_FMT_S16P,
		avutil.AV_SAMPLE_FMT_S32P:
		// integer => Scale to [-1, 1] and convert to float.
		div := ((1 << (uint(sampleSize)*8 - 1)) - 1)
		ret = float32(val) / float32(div)
		break

	case avutil.AV_SAMPLE_FMT_FLT,
		avutil.AV_SAMPLE_FMT_FLTP:
		// float => reinterpret
		ret = math.Float32frombits(uint32(val))
		break

	case avutil.AV_SAMPLE_FMT_DBL,
		avutil.AV_SAMPLE_FMT_DBLP:
		// double => reinterpret and then static cast down
		ret = float32(math.Float64frombits(uint64(val)))
		break

	default:
		panic(fmt.Sprintf("Invalid sample format %s.", codecCtxt.SampleFmt().Name()))
	}

	return ret
}

func drainDecoder(codecCtxt *avcodec.CodecContext, frame *avutil.Frame, out io.Writer) {
	// Some codecs may buffer frames. Sending NULL activates drain-mode.
	if err := codecCtxt.SendPacket(nil); err.Ok() {
		// Read the remaining packets from the decoder.
		err = receiveAndHandle(codecCtxt, frame, out)
		if !err.IsOneOf(avutil.AVERROR_EAGAIN(), avutil.AVERROR_EOF()) {
			// Neither EAGAIN nor EOF => Something went wrong.
			panicOnCode(err)
		}
	} else {
		// Something went wrong.
		panicOnCode(err)
	}
}

func main() {
	// Handle missuse.
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Fprintln(os.Stderr, "Usage: go run decode.go <audiofile> [outfile]")
		os.Exit(-1)
	}

	inFilename := os.Args[1]
	var outFilename string
	if len(os.Args) == 3 {
		outFilename = os.Args[2]
	} else {
		outFilename = os.Args[1] + ".raw"
	}

	// Open the outFile
	outFile, err := os.OpenFile(outFilename, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	panicOnErr(err)
	// Remember to close it afterwards.
	defer outFile.Close()
	bufFile := bufio.NewWriter(outFile)
	defer bufFile.Flush()

	// Initialize the libavformat. This registers all muxers, demuxers and protocols.
	avformat.RegisterAll()

	formatCtx, code := avformat.OpenInput(inFilename, nil, nil)
	panicOnCode(code)
	// Remember to clean up.
	defer formatCtx.Close()

	// In case the file had no header, read some frames and find out which format and codecs are used.
	// This does not consume any data. Any read packets are buffered for later use.
	formatCtx.FindStreamInfo(nil)

	// Try to find an audio stream.
	audioStreamIndex, err := findFirstAudioStream(formatCtx)
	panicOnErr(err)

	// Find the correct decoder for the codec.
	codec := avcodec.FindDecoder(formatCtx.Streams()[audioStreamIndex].CodecPar().CodecID())
	panicOnNil(codec, "Decoder not found. The codec is not supported.")

	// Initialize codec context for the decoder.
	codecCtxt := avcodec.NewCodecContext(codec)
	panicOnNil(codecCtxt, "Could not allocate a decoding context.")
	// Remember to clean up.
	defer codecCtxt.Free()
	defer codecCtxt.Close()

	// Fill the codecCtx with the parameters of the codec used in the read file.
	code = codecCtxt.FromParameters(formatCtx.Streams()[audioStreamIndex].CodecPar())
	panicOnCode(code)

	// Explicitly request non planar data.
	codecCtxt.SetRequestSampleFmt(codecCtxt.SampleFmt().Packed())

	// Initialize the decoder.
	code = codecCtxt.Open(codec, nil)
	panicOnCode(code)

	// Print some intersting file information.
	printStreamInformation(codec, codecCtxt, audioStreamIndex)

	frame := avutil.NewFrame()
	panicOnNil(frame, "Could not allocate frame.")
	// Remember to clean up.
	defer frame.Free()

	var packet avcodec.Packet
	packet.Init()

	for {
		// Read next frame
		code = formatCtx.ReadFrame(&packet)
		if code.IsOneOf(avutil.AVERROR_EOF()) {
			break
		}
		panicOnCode(code)

		// Does the packet belong to the correct stream?
		if packet.StreamIndex() != audioStreamIndex {
			// Free the buffers used by the frame and reset all fields.
			packet.Unref()
			continue
		}

		// We have a valid packet => send it to the decoder.
		code = codecCtxt.SendPacket(&packet)
		if code.Ok() {
			// The packet was sent successfully. We don't need it anymore.
			// => Free the buffers used by the frame and reset all fields.
			packet.Unref()
		} else {
			// Something went wrong.
			// EAGAIN is technically no error here but if it occurs we would need to buffer
			// the packet and send it again after receiving more frames. Thus we handle it as an error here.
			panic(code)
		}

		// Receive and handle frames.
		// EAGAIN means we need to send before receiving again. So thats not an error.
		code = receiveAndHandle(codecCtxt, frame, bufFile)
		if !code.IsOneOf(avutil.AVERROR_EAGAIN()) {
			// Not EAGAIN => Something went wrong.
			panic(code)
		}
	}

	// Drain the decoder.
	drainDecoder(codecCtxt, frame, bufFile)

	// Cleaning up is done by the go defer statements.
}

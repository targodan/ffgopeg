// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)
// Corbatto (luca@corbatto.de)

package avdevice

/*
	#cgo pkg-config: libavdevice
	#include <libavdevice/avdevice.h>
*/
import "C"

func inputAudioDeviceNext(d *InputFormat) *InputFormat {
	return (*InputFormat)(C.av_input_audio_device_next((*C.struct_AVInputFormat)(d)))
}

// InputAudioDevices returns a channel which can be used to iterate over the input audio devices.
//
// Usage:
//
//     for d := range avdevice.InputAudioDevices() {
//         // ...
//     }
//
// C-Function: av_input_audio_device_next
func InputAudioDevices() <-chan *InputFormat {
	ch := make(chan *InputFormat)

	var e *InputFormat
	go func() {
		for {
			e = inputAudioDeviceNext(e)
			if e == nil {
				break
			}
			ch <- e
		}
		close(ch)
	}()

	return ch
}

func inputVideoDeviceNext(d *InputFormat) *InputFormat {
	return (*InputFormat)(C.av_input_video_device_next((*C.struct_AVInputFormat)(d)))
}

// InputVideoDevices returns a channel which can be used to iterate over the input video devices.
//
// Usage:
//
//     for d := range avdevice.InputVideoDevices() {
//         // ...
//     }
//
// C-Function: av_input_video_device_next
func InputVideoDevices() <-chan *InputFormat {
	ch := make(chan *InputFormat)

	var e *InputFormat
	go func() {
		for {
			e = inputVideoDeviceNext(e)
			if e == nil {
				break
			}
			ch <- e
		}
		close(ch)
	}()

	return ch
}

func outputAudioDeviceNext(d *OutputFormat) *OutputFormat {
	return (*OutputFormat)(C.av_output_audio_device_next((*C.struct_AVOutputFormat)(d)))
}

// OutputAudioDevices returns a channel which can be used to iterate over the output audio devices.
//
// Usage:
//
//     for d := range avdevice.OutputAudioDevices() {
//         // ...
//     }
//
// C-Function: av_output_audio_device_next
func OutputAudioDevices() <-chan *OutputFormat {
	ch := make(chan *OutputFormat)

	var e *OutputFormat
	go func() {
		for {
			e = outputAudioDeviceNext(e)
			if e == nil {
				break
			}
			ch <- e
		}
		close(ch)
	}()

	return ch
}

func outputVideoDeviceNext(d *OutputFormat) *OutputFormat {
	return (*OutputFormat)(C.av_output_video_device_next((*C.struct_AVOutputFormat)(d)))
}

// OutputVideoDevices returns a channel which can be used to iterate over the output video devices.
//
// Usage:
//
//     for d := range avdevice.OutputVideoDevices() {
//         // ...
//     }
//
// C-Function: av_output_video_device_next
func OutputVideoDevices() <-chan *OutputFormat {
	ch := make(chan *OutputFormat)

	var e *OutputFormat
	go func() {
		for {
			e = outputVideoDeviceNext(e)
			if e == nil {
				break
			}
			ch <- e
		}
		close(ch)
	}()

	return ch
}

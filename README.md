# goav

This go library provides bindings for ffmpeg.
This is a permanent fork of [giorgisio/goav](https://github.com/giorgisio/goav) and will differ in the following ways.

- Target ffmpeg version: 3.1.n
- No deprecated functions.
- OOP-Style where applicable, see [Mapping ffmpeg function names to go](#mapping-ffmpeg-function-names-to-go).
- Renaming of functions to make sense in the go context, again see [Mapping ffmpeg function names to go](#mapping-ffmpeg-function-names-to-go).
- I merged some pending pull request on the original.

## Current state

The master branch for now is on the same level as the original repo (except this readme file).
There are some feature branches from when I opened pull requests.
The delvelop branch is under heavy development.

## Usage

**Don't use it yet.**

As it is under massive development and anything not only can but *will* change at any moment, I would advise to not use it yet.
Stay tuned though, I'll give my best to release a 1.0 ASAP.

## Mapping ffmpeg function names to go

The ffmpeg C-Functions are renamed following these rules:

- Everything will be camel case: `foo_bar` ➡️ `FooBar`
- Library-level functions like `avcodec_register_all` are going to be package level functions with the prefix removed: `avcodec.RegisterAll`
- Functions that return something which is technically a `bool` will now return a go `bool`: `see next example`
- Functions that modify or access a struct in a sort-of-OOP way will become methods: `int av_codec_is_encoder(const AVCodec* codec)` ➡️ `func (codec *Codec) IsEncoder() bool`
- Allocating functions that are used as a constructor are going to be called `NewXxx`: `avcodec_alloc_context3` ➡️ `func NewContext() *CodecContext`
- Instead of `int` return codes go `error`s will be returned containing an error description (as returned by `av_strerror`) and the code.

In the future I will provide a little tool wich lets you search for the go-representation of a given ffmpeg C-function.

## Contribute

Until I released version 1.0 and defined where exactly I want to go with this project, please don't send any pull requests or open any issues.

There is one point where I could use some help though:
If anyone has a good idea how to test the bindings I would be glad to hear it!
I don't think it makes sense to test the behaviour of simple C calls as I assume these functions are already tested by ffmpeg.
I don't know how I could test if the binding was successful without testing the behaviour though.
Apart from that in order to test the behaviour I would need to dive way way way deeper into ffmpeg than I anticipate.

## My TODO list

- [ ] Do the actual renames
- [ ] Document the whole thing
- [ ] Implement the replacements for deprecated functions
- [ ] Setup TravisCI
- [ ] Test the library

## Libraries

- `avcodec` corresponds to the ffmpeg library: libavcodec [provides implementation of a wider range of codecs]
- `avformat` corresponds to the ffmpeg library: libavformat [implements streaming protocols, container formats and basic I/O access]
- `avutil` corresponds to the ffmpeg library: libavutil [includes hashers, decompressors and miscellaneous utility functions]
- `avfilter` corresponds to the ffmpeg library: libavfilter [provides a mean to alter decoded Audio and Video through chain of filters]
- `avdevice` corresponds to the ffmpeg library: libavdevice [provides an abstraction to access capture and playback devices]
- `swresample` corresponds to the ffmpeg library: libswresample [implements audio mixing and resampling routines]
- `swscale` corresponds to the ffmpeg library: libswscale [implements color conversion and scaling routines]

## License

This library is under the [MIT License](http://opensource.org/licenses/MIT)

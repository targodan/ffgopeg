# ffgopeg [![Build Status](https://travis-ci.org/targodan/ffgopeg.svg?branch=develop)](https://travis-ci.org/targodan/ffgopeg)

This go library provides bindings for ffmpeg.
Its main goal is to make the ffmpeg library accessable in a way that feels natural in the go language.
This library is based on [giorgisio/goav](https://github.com/giorgisio/goav) and will differ in the following ways.

- Target ffmpeg version: 3.1.n
- No deprecated functions.
- OOP-Style where applicable, see [Mapping ffmpeg function names to go](#mapping-ffmpeg-function-names-to-go).
- Renaming of functions to make sense in the go context, again see [Mapping ffmpeg function names to go](#mapping-ffmpeg-function-names-to-go).
- I merged some pending pull request on the original.

## Current state

TL; DR: *not stable*

The renaming of funcitons is mostly complete.
Although I have released version 1 I wouldn't call it stable yet.
It is largely untested and many functions may not work as intended (meaning propable segmentation faults).
If you discover any errors and/or have functions that are missing feel free to [contribute](#contribute)! :grinning:

If you just want to decode audio files into `float32` samples [this project of mine](https://github.com/targodan/goad) could be for you.
It provides a high level API for just that usecase, based on the ffgopeg library.
A first version is going to come soon but feel free to checkout the develop-branch in the meantime ("checkout"... get it? :laughing:).

## Usage

Installing or updating.

```
go get -u gopkg.in/targodan/ffgopeg.v1
```

This will install the latest version 1.n.m release.
See [Versioning](#versioning).

The ffmpeg libraries are in theire own packages.
If you know which ffmpeg function you want to use you can use [this little webtool](https://targodan.github.io/ffgopeg)

## Versioning

This project follows the [semantic versioning scheme](http://semver.org/).
That means if you use the `gopkg.in` link from above using `go get -u` should be safe and not break anything.

## Mapping ffmpeg function names to go

The ffmpeg C-Functions are renamed following these rules:

- Everything will be camel case: `foo_bar`  `FooBar`
- Library-level functions like `avcodec_register_all` are going to be package level functions with the prefix removed: `avcodec.RegisterAll`
- Functions that return something which is technically a `bool` will now return a go `bool`: `see next example`
- Functions that modify or access a struct in a sort-of-OOP way will become methods: `int av_codec_is_encoder(const AVCodec* codec)`  `func (codec *Codec) IsEncoder() bool`
- Allocating functions that are used as a constructor are going to be called `NewXxx`: `avcodec_alloc_context3`  `func NewContext() *CodecContext`
- Instead of `int` return codes go `error`s will be returned containing an error description (as returned by `av_strerror`) and the code.

In the future I will provide a little tool wich lets you search for the go-representation of a given ffmpeg C-function.

## Contribute

If you get any errors using this library or some functions are missing please open an issue and/or fork, branch and file a pull request.
Please only file pull requests on the develop branch.
This project losely followes the [git-flow](http://danielkummer.github.io/git-flow-cheatsheet/) workflow.

After opening the pull request travis ci should start compiling and testing it.
I won't react to your pull request untill travis *successfully* builds the request.
If travis errors just add commits to your fork that will fix it.
Travis will then rebuild the request.
If your git history becomes messy by wildly trying to fix the build, please do a `git rebase -i HEAD~n` (with n as number of commits) and clean up your history.

Also I could use some help with these things:
If anyone has a good idea how to test the bindings I would be glad to hear it!
I don't think it makes sense to test the behaviour of simple C calls as I assume these functions are already tested by ffmpeg.
I don't know how I could test if the binding was successful without testing the behaviour though.
Apart from that in order to test the behaviour I would need to dive way way way deeper into ffmpeg than I anticipate.

## My TODO list

- [x] Do the actual renames
- [ ] Document the whole thing
- [ ] Implement the replacements for deprecated functions
- [x] Setup TravisCI
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

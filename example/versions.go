package main

import (
	"github.com/targodan/goav/avcodec"
	"github.com/targodan/goav/avdevice"
	"github.com/targodan/goav/avfilter"
	"github.com/targodan/goav/avformat"
	"github.com/targodan/goav/avutil"
	"github.com/targodan/goav/swresample"
	"github.com/targodan/goav/swscale"
	"log"
)

func main() {

	// Register all formats and codecs
	avformat.AvRegisterAll()
	avcodec.AvcodecRegisterAll()

	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())

}

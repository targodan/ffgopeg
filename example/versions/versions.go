package main

import (
	"log"

	"gopkg.in/targodan/ffgopeg.v1/avcodec"
	"gopkg.in/targodan/ffgopeg.v1/avdevice"
	"gopkg.in/targodan/ffgopeg.v1/avfilter"
	"gopkg.in/targodan/ffgopeg.v1/avformat"
	"gopkg.in/targodan/ffgopeg.v1/avutil"
	"gopkg.in/targodan/ffgopeg.v1/swresample"
	"gopkg.in/targodan/ffgopeg.v1/swscale"
)

func main() {
	log.Printf("AvCodec Version:\t%v", avcodec.Version())
	log.Printf("AvCodec License:\t%v", avcodec.License())
	log.Printf("AvDevice Version:\t%v", avdevice.Version())
	log.Printf("AvDevice License:\t%v", avdevice.License())
	log.Printf("AvFilter Version:\t%v", avfilter.Version())
	log.Printf("AvFilter License:\t%v", avfilter.License())
	log.Printf("AvFormat Version:\t%v", avformat.Version())
	log.Printf("AvFormat License:\t%v", avformat.License())
	log.Printf("AvUtil Version:\t%v", avutil.Version())
	log.Printf("AvUtil License:\t%v", avutil.License())
	log.Printf("SWResample Version:\t%v", swresample.Version())
	log.Printf("SWResample License:\t%v", swresample.License())
	log.Printf("SWScale Version:\t%v", swscale.Version())
	log.Printf("SWScale License:\t%v", swscale.License())
}

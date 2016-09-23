// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)
// Corbatto (luca@corbatto.de)

// Package avdevice deals with the input and output devices provided by the libavdevice library
// The libavdevice library provides the same interface as libavformat.
// Namely, an input device is considered like a demuxer, and an output device like a muxer,
// and the interface and generic device options are the same provided by libavformat
package avdevice

/*
	#cgo pkg-config: libavdevice
	#include <libavdevice/avdevice.h>
*/
import "C"
import (
	"unsafe"

	"github.com/targodan/goav/avutil"
)

type (
	DeviceRect              C.struct_AVDeviceRect
	DeviceCapabilitiesQuery C.struct_AVDeviceCapabilitiesQuery
	DeviceInfo              C.struct_AVDeviceInfo
	DeviceInfoList          C.struct_AVDeviceInfoList
	InputFormat             C.struct_AVInputFormat
	OutputFormat            C.struct_AVOutputFormat
	FormatContext           C.struct_AVFormatContext
	Dictionary              C.struct_AVDictionary
	AppToDevMessageType     C.enum_AVAppToDevMessageType
	DevToAppMessageType     C.enum_AVDevToAppMessageType
)

// Version returns the libavdevice version.
//
// C-Function: avdevice_version
func Version() uint {
	return uint(C.avdevice_version())
}

// Configuration returns the libavdevice build-time configuration.
//
// C-Function: avdevice_configuration
func Configuration() string {
	return C.GoString(C.avdevice_configuration())
}

// License returns the libavdevice license.
//
// C-Function: avdevice_license
func License() string {
	return C.GoString(C.avdevice_license())
}

// RegisterAll initializes libavdevice and register all the input and output devices.
//
// C-Function: avdevice_register_all
func RegisterAll() {
	C.avdevice_register_all()
}

// AppToDevControlMessage sends control message from application to device.
//
// C-Function: avdevice_app_to_dev_control_message
func AppToDevControlMessage(s *FormatContext, m AppToDevMessageType, data []interface{}) error {
	return avutil.CodeToError(int(C.avdevice_app_to_dev_control_message((*C.struct_AVFormatContext)(s), (C.enum_AVAppToDevMessageType)(m), unsafe.Pointer(&data[0]), C.size_t(len(data)))))
}

// DevToAppControlMessage sends control message from device to application.
//
// C-Function: avdevice_dev_to_app_control_message
func DevToAppControlMessage(fcxt *FormatContext, m DevToAppMessageType, data []interface{}) error {
	return avutil.CodeToError(int(C.avdevice_dev_to_app_control_message((*C.struct_AVFormatContext)(fcxt), (C.enum_AVDevToAppMessageType)(m), unsafe.Pointer(&data[0]), C.size_t(len(data)))))
}

// CapabilitiesCreate initializes capabilities probing API based on AvOption API.
//
// CapabilitiesFree() must be called when query capabilities API is not used anymore.
//
// C-Function: avdevice_capabilities_create
func CapabilitiesCreate(c **DeviceCapabilitiesQuery, s *FormatContext, d **Dictionary) error {
	return avutil.CodeToError(int(C.avdevice_capabilities_create((**C.struct_AVDeviceCapabilitiesQuery)(unsafe.Pointer(c)), (*C.struct_AVFormatContext)(s), (**C.struct_AVDictionary)(unsafe.Pointer(d)))))
}

// CapabilitiesFree frees resources created by CapabilitiesCreate()
//
// C-Function: avdevice_capabilities_free
func CapabilitiesFree(c **DeviceCapabilitiesQuery, s *FormatContext) {
	C.avdevice_capabilities_free((**C.struct_AVDeviceCapabilitiesQuery)(unsafe.Pointer(c)), (*C.struct_AVFormatContext)(s))
}

// ListDevices lists devices.
//
// C-Function: avdevice_list_devices
func ListDevices(s *FormatContext) (*DeviceInfoList, error) {
	var dl *DeviceInfoList
	err := avutil.CodeToError(int(C.avdevice_list_devices((*C.struct_AVFormatContext)(s), (**C.struct_AVDeviceInfoList)(unsafe.Pointer(&dl)))))
	return dl, err
}

// Free is a convenient function to free result of ListDevices().
//
// C-Function: avdevice_free_list_devices
func (d *DeviceInfoList) Free() {
	C.avdevice_free_list_devices((**C.struct_AVDeviceInfoList)(unsafe.Pointer(&d)))
}

package netsdk

// #cgo CFLAGS: -I .
// #cgo LDFLAGS: -L . -ldhnetsdk
// #include <stdio.h>
// #include <stdlib.h>
// #include "dhnetsdk.h"
import (
	"C"
)

type RealPlay struct {
	client *Client
	handle int
}

type WindowHandle uint64

type RTType int

const (
	DH_RType_Realplay RTType = iota
	DH_RType_Multiplay
	DH_RType_Realplay_0
	DH_RType_Realplay_1
	DH_RType_Realplay_2
	DH_RType_Realplay_3
	DH_RType_Multiplay_1
	DH_RType_Multiplay_4
	DH_RType_Multiplay_8
	DH_RType_Multiplay_9
	DH_RType_Multiplay_16
	DH_RType_Multiplay_6
	DH_RType_Multiplay_12
	DH_RType_Multiplay_25
	DH_RType_Multiplay_36
)

func (client *Client) OpenRealPlay(channalId int, hWnd WindowHandle, rtType RTType) (*RealPlay, error) {
	var realplay = &RealPlay{client: client}
	nret := C.CLIENT_RealPlayEx(
		C.LLONG(client.LoginID), C.int(channalId), nil, C.DH_RealPlayType(rtType))

	if nret == 0 {
		return nil, ErrLast()
	}

	realplay.handle = int(nret)

	return realplay, nil
}

type CaptureFormat int

const (
	NET_CAPTURE_BMP     = iota // BMP
	NET_CAPTURE_JPEG           // 100% quality JPEG
	NET_CAPTURE_JPEG_70        // 70% quality JPEG
	NET_CAPTURE_JPEG_50        // 50% quality JPEG
	NET_CAPTURE_JPEG_30        // 30% quality JPEG
	NET_CAPTURE_JPEG_10        // 10% quality JPEG
)

// CapturePictureEx
func (real *RealPlay) CapturePictureEx(filename string, format CaptureFormat) error {

	if C.CLIENT_CapturePictureEx(C.LLONG(real.handle), C.CString(filename), C.NET_CAPTURE_FORMATS(format)) == 0 {
		return ErrLast()
	}

	return nil
}

func (real *RealPlay) Close() error {

	if C.CLIENT_StopRealPlayEx(C.LLONG(real.handle)) == 0 {
		return ErrLast()
	}

	return nil
}

package netsdk

// #cgo CFLAGS: -I .
// #cgo LDFLAGS: -L . -ldhnetsdk

// #include <stdio.h>
// #include <stdlib.h>
// #include "dhnetsdk.h"
// extern void export_fSearchDevicesCBEx2(long lSearchHandle, DEVICE_NET_INFO_EX2 *pDevNetInfo, long pUserData);
import "C"

import (
	"log"
	"unsafe"

	"github.com/mattn/go-pointer"
)

type (
	SearchDeviceFunc func(search *Search, deviceinfoEx *DEVICE_NET_INFO_EX)
	SearchVisitor    struct {
		Search *Search

		Callback SearchDeviceFunc
	}
	Search struct {
		LocalIP      string
		searchHandle int64
		visitorp     unsafe.Pointer
		inparam      NET_IN_STARTSERACH_DEVICE
	}
)

func (search *Search) Start(fn SearchDeviceFunc, searchType EM_SEND_SEARCH_TYPE) error {
	var (
		outparam NET_OUT_STARTSERACH_DEVICE
		inparam  NET_IN_STARTSERACH_DEVICE

		v = SearchVisitor{
			Search:   search,
			Callback: fn,
		}
	)

	p := pointer.Save(v)
	search.visitorp = p

	// handle := StartSearchDevicesEx(&inparam, &outparam)
	copy(inparam.ST_szLocalIp[:], search.LocalIP)
	log.Println("Enter StartSearchDeviceEx...")

	var pstInBuf *C.struct_tagNET_IN_STARTSERACH_DEVICE = (*C.struct_tagNET_IN_STARTSERACH_DEVICE)(unsafe.Pointer(&inparam))
	(*pstInBuf).dwSize = (C.uint)(unsafe.Sizeof(C.struct_tagNET_IN_STARTSERACH_DEVICE{}))

	var pstOutBuf *C.struct_tagNET_OUT_STARTSERACH_DEVICE = (*C.struct_tagNET_OUT_STARTSERACH_DEVICE)(unsafe.Pointer(&outparam))
	(*pstOutBuf).dwSize = (C.uint)(unsafe.Sizeof(C.struct_tagNET_OUT_STARTSERACH_DEVICE{}))

	// pstInBuf.pUserData = unsafe.Pointer(uintptr(*((*int)(unsafe.Pointer(&inparam.ST_pUserData)))))
	pstInBuf.pUserData = p
	pstInBuf.cbSearchDevices = C.fSearchDevicesCBEx(C.export_fSearchDevicesCBEx2)

	lSearchHandle := C.CLIENT_StartSearchDevicesEx(pstInBuf, pstOutBuf)
	if lSearchHandle != 0 {
		log.Println("CLIENT_StartSearchDevicesEx success")
	}
	search.searchHandle = int64(lSearchHandle)

	log.Println("lSearchHandle=", lSearchHandle)
	log.Println("dwSize is :", pstInBuf.dwSize)
	GetLastError()

	return nil
}

func (search *Search) Invoke(lSearchHandle LLONG, pDevNetInfo *DEVICE_NET_INFO_EX2) {
	p := pointer.Restore(search.visitorp)
	if visitor, ok := p.(SearchVisitor); ok {
		visitor.Callback(search, &pDevNetInfo.ST_stuDevInfo)
	}
}

func (search *Search) Stop() bool {
	if search.searchHandle == 0 {
		return false
	}

	if search.visitorp != nil {
		pointer.Unref(search.visitorp)
	}
	return StopSearchDevices(search.searchHandle)
}

//export export_fSearchDevicesCBEx2
func export_fSearchDevicesCBEx2(lSearchHandle C.long, pDevNetInfo *C.struct_tagDEVICE_NET_INFO_EX2, pUserData C.long) {
	var user_data = unsafe.Pointer(uintptr(pUserData))
	if v, ok := pointer.Restore(user_data).(SearchVisitor); ok {
		if v.Callback != nil {
			devinfo := (*DEVICE_NET_INFO_EX2)(unsafe.Pointer(pDevNetInfo))
			v.Callback(v.Search, &devinfo.ST_stuDevInfo)
		}
	}
}

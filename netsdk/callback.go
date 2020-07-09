package netsdk

// #cgo CFLAGS: -I .
// #cgo LDFLAGS: -L . -ldhnetsdk
// #include <stdio.h>
// #include <stdlib.h>
// #include "dhnetsdk.h"
import "C"
import (
	"log"
	"reflect"
	"unsafe"

	"github.com/mattn/go-pointer"
)

//export export_fDisconnect2
func export_fDisconnect2(userdata C.long, ip *C.char, port C.LONG) {
	if v, ok := pointer.Restore(unsafe.Pointer(uintptr(userdata))).(DisconnectVisitor); ok {
		// defer pointer.Unref(user_data)
		if v.Callback != nil {
			v.Callback(C.GoString(ip), int(port))
		}
	}
}

// 声明一个函数类型
type CallBack_fDisConnect func(lLoginID LLONG, pchDVRIP string, nDVRPort int, dwUser LLONG)
type CallBack_fAnalyzerDataCallBack func(lAnalyzerHandle LLONG, dwAlarmType DWORD, pAlarmInfo LLONG,
	pBuffer LLONG, dwBufSize DWORD, dwUser LLONG, nSequence int32, reserved LLONG)

//export goReConnect
func goReConnect(user_data unsafe.Pointer, ip *C.char, port C.int) {
	if v, ok := pointer.Restore(user_data).(ReconnectVisitor); ok {
		// defer pointer.Unref(user_data)
		if v.Callback != nil {
			v.Callback(C.GoString(ip), int(port))
		}
	}

}

//export goDvrMessage
func goDvrMessage(user_data unsafe.Pointer, cmd C.int, buf *C.char, l C.int, ip *C.char, port C.int) {
	if v, ok := pointer.Restore(user_data).(DrvMessageVisitor); ok {
		// defer pointer.Unref(user_data)
		if v.Callback != nil {
			v.Callback(DhAlarmType(cmd), C.GoBytes(unsafe.Pointer(buf), l), C.GoString(ip), int(port))
		}
	}
}

//export export_fDisConnect
func export_fDisConnect(lLoginID C.long, pchDVRIP *C.char, nDVRPort C.int, dwUser C.long) {
	log.Println(pchDVRIP, " disconnect...")
	log.Println(" ------------------------------------------------------------------")
	log.Println(pchDVRIP, " disconnect...")
	var InnerUser LLONG = LLONG(dwUser)
	var rrr *CallBack_fDisConnect = (*CallBack_fDisConnect)((unsafe.Pointer(&InnerUser)))

	(*rrr)(LLONG(lLoginID), C.GoString(pchDVRIP), int(nDVRPort), LLONG(dwUser))
}

type IF_fServiceCallBack interface {
	Invoke(lHandle int, pIp string, wPort uint16, lCommand int, pParam uintptr, dwParamLen int) int
}

//export export_fServiceCallBack
func export_fServiceCallBack(lHandle C.long, pIp *C.char, wPort C.ushort, lCommand C.int, pParam C.long, dwParamLen C.int, dwUserData C.long) C.int {
	// log.Println(int(lHandle), C.GoString(pIp), uint16(wPort), int(lCommand), uintptr(pParam), int(dwParamLen), int(dwUserData))

	id := ObjectId(dwUserData)
	interf := id.Get()
	obj, ok := interf.(IF_fServiceCallBack)

	var ret int
	if ok {
		ret = obj.Invoke(int(lHandle), C.GoString(pIp), uint16(wPort), int(lCommand), uintptr(pParam), int(dwParamLen))
	} else {
		log.Panicf("%T is not IF_fServiceCallBack", interf)
	}
	return C.int(ret)
}

type IF_fMessCallBack interface {
	Invoke(lCommand int, lLoginID int, pBuf unsafe.Pointer, dwBufLen int, pchDVRIP string, nDVRPort int) bool
}

//export export_fMessCallBack
func export_fMessCallBack(lCommand C.int, lLoginID C.long, pBuf *C.char, dwBufLen C.int, pchDVRIP *C.char, nDVRPort C.int, dwUser C.long) C.int {
	//log.Println(int(lCommand), int(lLoginID), int(dwBufLen), C.GoString(pchDVRIP), int(nDVRPort), int(dwUser))
	id := ObjectId(dwUser)
	interf := id.Get()
	obj, ok := interf.(IF_fMessCallBack)
	if ok {
		ret := obj.Invoke(int(lCommand), int(lLoginID), unsafe.Pointer(pBuf), int(dwBufLen), C.GoString(pchDVRIP), int(nDVRPort))
		if ret {
			return 1
		}
	} else {
		log.Panicf("%T is not IF_fMessCallBack", interf)
	}
	return 0
}

// declare fun log
type CallBack_fSDKLogCallBack func(szLogBuffer string, nLogSize uint32)

//export export_fSDKLogCallBack
func export_fSDKLogCallBack(szLogBuffer *C.char, nLogSize C.uint, dwUser C.long) int32 {
	log.Println("Enter fSDKLogCallBack...")
	var cb *CallBack_fSDKLogCallBack = (*CallBack_fSDKLogCallBack)(unsafe.Pointer(&dwUser))
	log.Println("*cb is ", *cb)
	if nil != *cb {
		(*cb)(C.GoString(szLogBuffer), uint32(nLogSize))
	}
	return 0

}

//export export_fAnalyzerDataCallBack
func export_fAnalyzerDataCallBack(lAnalyzerHandle C.long, dwAlarmType C.uint, pAlarmInfo C.long, pBuffer C.long, dwBufSize C.uint, dwUser C.long, nSequence C.int, reserved C.long) int32 {
	log.Println("Enter fAnalyzerDataCallBack...")
	var InnerUser LLONG = LLONG(dwUser)
	var cb *CallBack_fAnalyzerDataCallBack = (*CallBack_fAnalyzerDataCallBack)((unsafe.Pointer(&InnerUser)))

	log.Printf("cb %v", cb)
	(*cb)(LLONG(lAnalyzerHandle), DWORD(dwAlarmType), LLONG(pAlarmInfo), LLONG(pBuffer),
		DWORD(dwBufSize), LLONG(dwUser), int32(nSequence), LLONG(reserved))

	return 1
}

//export export_fAnalyzerDataCallBack2
func export_fAnalyzerDataCallBack2(lAnalyzerHandle C.long, dwAlarmType C.uint, pAlarmInfo C.long, pBuffer C.long, dwBufSize C.DWORD, dwUser C.long, nSequence C.int, reserved C.long) int32 {
	log.Println("Enter fAnalyzerDataCallBack...")

	p := pointer.Restore(unsafe.Pointer(uintptr(dwUser)))

	visitor, ok := p.(*PictureVisitor)
	if !ok {
		log.Printf("visitor nil")
		return 0

	}

	switch EVENT_IVS(dwAlarmType) {
	case EVENT_IVS_TRAFFIC_PARKING:
		alarmType := (*DEV_EVENT_TRAFFIC_PARKING_INFO)(unsafe.Pointer(uintptr(pAlarmInfo)))
		var buf []byte
		data := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
		data.Data = uintptr(unsafe.Pointer(uintptr(pBuffer)))
		data.Len = int(dwBufSize)
		data.Cap = int(dwBufSize)

		visitor.Callback(visitor.Client, EventIvs(dwAlarmType), alarmType, buf, int(nSequence))
	}

	return 1
}

type IF_fSearchDevicesCBEx interface {
	Invoke(lSearchHandle LLONG, pDevNetInfo *DEVICE_NET_INFO_EX2)
}

//export export_fSearchDevicesCBEx
func export_fSearchDevicesCBEx(lSearchHandle C.long, pDevNetInfo *C.struct_tagDEVICE_NET_INFO_EX2, pUserData C.long) {
	//	log.Println("Enter export_fSearchDevicesCBEx...")
	var cb *IF_fSearchDevicesCBEx = (*IF_fSearchDevicesCBEx)(unsafe.Pointer(&pUserData))
	(*cb).Invoke(LLONG(lSearchHandle), (*DEVICE_NET_INFO_EX2)(unsafe.Pointer(pDevNetInfo)))
}

type IF_fSnapRev interface {
	Invoke(lLoginID int, pBuf []byte, EncodeType uint, CmdSerial uint)
}

//export export_fSnapRev
func export_fSnapRev(lLoginID C.long, pBuf *C.uchar, RevLen C.uint, EncodeType C.uint, CmdSerial C.uint, dwUser C.long) {
	id := ObjectId(dwUser)
	interf := id.Get()
	obj, ok := interf.(IF_fSnapRev)
	if ok {
		var buf []byte
		data := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
		data.Data = uintptr(unsafe.Pointer(pBuf))
		data.Len = int(RevLen)
		data.Cap = int(RevLen)
		obj.Invoke(int(lLoginID), buf, uint(EncodeType), uint(CmdSerial))
	} else {
		log.Panicf("%T is not IF_fSnapRev", interf)
	}
}

// Corresponding to SearchDevicesByIPs API
type IF_fSearchDevicesCB interface {
	Invoke(pDevNetInfo *DEVICE_NET_INFO_EX)
}

//export export_fSearchDevicesCB
func export_fSearchDevicesCB(pDevNetInfo *C.struct_tagDEVICE_NET_INFO_EX, pUserData C.long) {
	//	log.Println("Enter export_fSearchDevicesCB...")
	var cb *IF_fSearchDevicesCB = (*IF_fSearchDevicesCB)(unsafe.Pointer(&pUserData))
	(*cb).Invoke((*DEVICE_NET_INFO_EX)(unsafe.Pointer(pDevNetInfo)))

}

type IF_fTimeDownLoadPosCallBack interface {
	Invoke(lPlayHandle int, dwTotalSize uint, dwDownLoadSize int, index int, recordfileinfo NET_RECORDFILE_INFO)
}

//export export_fTimeDownLoadPosCallBack
func export_fTimeDownLoadPosCallBack(lPlayHandle C.long, dwTotalSize C.uint, dwDownLoadSize C.int, index C.int, recordfileinfo C.NET_RECORDFILE_INFO, dwUser C.long) {
	id := ObjectId(dwUser)
	interf := id.Get()
	obj, ok := interf.(IF_fTimeDownLoadPosCallBack)
	if ok {
		obj.Invoke(int(lPlayHandle), uint(dwTotalSize), int(dwDownLoadSize), int(index), *(*NET_RECORDFILE_INFO)(unsafe.Pointer(&recordfileinfo)))
	} else {
		log.Panicf("%T is not IF_fTimeDownLoadPosCallBack", interf)
	}
}

type IF_fDataCallBack interface {
	Invoke(lRealHandle int, dwDataType uint, pBuffer []byte) int
}

//export export_fDataCallBack
func export_fDataCallBack(lRealHandle C.long, dwDataType C.uint, pBuffer *C.uchar, dwBufSize C.uint, dwUser C.long) C.int {
	id := ObjectId(dwUser)
	interf := id.Get()

	/*switch t := obj.(type) {
	case IF_fDataCallBack:
		return C.int(t.Invoke(int(lRealHandle), uint(dwDataType), buf))
	}*/
	obj, ok := interf.(IF_fDataCallBack)
	if ok {
		var buf []byte
		data := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
		data.Data = uintptr(unsafe.Pointer(pBuffer))
		data.Len = int(dwBufSize)
		data.Cap = int(dwBufSize)
		return C.int(obj.Invoke(int(lRealHandle), uint(dwDataType), buf))
	} else {
		log.Panicf("%T is not IF_fDataCallBack", interf)
	}
	return C.int(dwBufSize)
}

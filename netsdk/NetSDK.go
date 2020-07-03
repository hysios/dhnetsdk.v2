// NetSDK project NetSDK.go

package netsdk

/*
#cgo CFLAGS: -I .
#cgo LDFLAGS: -L . -ldhnetsdk

#include <stdio.h>
#include <stdlib.h>
#include "dhnetsdk.h"

extern void export_fDisConnect(long lLoginID, char * pchDVRIP, int nDVRPort, long dwUser);
extern int export_fServiceCallBack(long lHandle, char * pIp, unsigned short wPort, int lCommand, long pParam, int dwParamLen, long dwUserData);
extern int export_fMessCallBack(int lCommand, long lLoginID, char * pBuf, int dwBufLen, char * pchDVRIP, int nDVRPort, long dwUser);
extern int export_fSDKLogCallBack(char * szLogBuffer, uint nLogSize, long dwUser);
extern int export_fAnalyzerDataCallBack(long lAnalyzerHandle, unsigned int dwAlarmType, long pAlarmInfo, long pBuffer,unsigned int dwBufSize, long dwUser, int nSequence,long reserved);
extern int export_fAnalyzerDataCallBack2(long lAnalyzerHandle, unsigned int dwAlarmType, long pAlarmInfo, long pBuffer,unsigned int dwBufSize, long dwUser, int nSequence,long reserved);
extern void export_fSearchDevicesCBEx(long lSearchHandle, DEVICE_NET_INFO_EX2 *pDevNetInfo, long pUserData);
extern void export_fSnapRev(long lLoginID, unsigned char *pBuf, unsigned int RevLen, unsigned int EncodeType, unsigned int CmdSerial, long dwUser);
extern void export_fSearchDevicesCB(DEVICE_NET_INFO_EX *pDevNetInfo, long pUserData);
extern void export_fTimeDownLoadPosCallBack(long lPlayHandle, unsigned int dwTotalSize, int dwDownLoadSize, int index, NET_RECORDFILE_INFO recordfileinfo, long dwUser);
extern int export_fDataCallBack(long lRealHandle, unsigned int dwDataType, unsigned char * pBuffer, unsigned int dwBufSize, long dwUser);
*/
import "C"

import (
	"log"
	"reflect"
	"unsafe"

	"github.com/mattn/go-pointer"
)

// 声明一个函数类型
type CallBack_fDisConnect func(lLoginID LLONG, pchDVRIP string, nDVRPort int, dwUser LLONG)
type CallBack_fAnalyzerDataCallBack func(lAnalyzerHandle LLONG, dwAlarmType DWORD, pAlarmInfo LLONG,
	pBuffer LLONG, dwBufSize DWORD, dwUser LLONG, nSequence int32, reserved LLONG)

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

type PictureVisitor struct {
	Client   *Client
	Callback PictureExFunc
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

// //export export_fSearchDevicesCBEx2
// func export_fSearchDevicesCBEx2(lSearchHandle C.long, pDevNetInfo *C.struct_tagDEVICE_NET_INFO_EX2, pUserData C.long) {
// 	var user_data = unsafe.Pointer(uintptr(pUserData))
// 	if v, ok := pointer.Restore(user_data).(SearchVisitor); ok {
// 		if v.Callback != nil {
// 			devinfo := (*DEVICE_NET_INFO_EX2)(unsafe.Pointer(pDevNetInfo))
// 			v.Callback(v.Search, &devinfo.ST_stuDevInfo)
// 		}
// 	}
// }

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

func InitEx(cb CallBack_fDisConnect, stuInfo *NETSDK_INIT_PARAM) bool {
	if nil == stuInfo {
		log.Println("stuInfo is nil")
		return false
	}

	var lpInitParam = C.struct_tagNETSDK_INIT_PARAM{}
	lpInitParam.nThreadNum = C.int(stuInfo.ST_nThreadNum)

	var userdata *LLONG = (*LLONG)(unsafe.Pointer(&cb))
	log.Println("userData=", *userdata)

	ret := C.CLIENT_InitEx(C.fDisConnect(C.export_fDisConnect), C.long(*userdata), (*C.struct_tagNETSDK_INIT_PARAM)(unsafe.Pointer(&lpInitParam)))
	if ret != 0 {
		log.Println(" call CLIENT_InitEx success")
		return true
	}
	return false
}

// SDK clean up
func Cleanup() {
	log.Println("Enter Cleanup....")
	C.CLIENT_Cleanup()
}

// SDK GetLastError Return the function execution failure code
func GetLastError() int {
	log.Println("Enter GetLastError....")
	return int(C.CLIENT_GetLastError())
}

// Login Ex API
func LoginEx2(pchDVRIP string, wDVRPort int, UserName string, Password string) int {
	log.Println("Enter NET_LoginEx2")
	pCapParam := unsafe.Pointer(uintptr(0))
	lpDeviceInfo := C.struct_tagNET_DEVICEINFO_Ex{}
	error := C.int(0)

	ipTmp := C.CString(pchDVRIP)
	UserNameTmp := C.CString(UserName)
	PasswdTmp := C.CString(Password)
	loginhandle := C.CLIENT_LoginEx2(
		ipTmp,
		C.ushort(wDVRPort),
		UserNameTmp,
		PasswdTmp,
		C.EM_LOGIN_SPEC_CAP_TCP,
		pCapParam,
		&lpDeviceInfo,
		&error)
	C.free(unsafe.Pointer(ipTmp))
	C.free(unsafe.Pointer(UserNameTmp))
	C.free(unsafe.Pointer(PasswdTmp))
	if loginhandle != 0 {
		log.Println("CLIENT_LoginEx2 success")
		log.Println("Serial Num:", C.GoString((*C.char)(unsafe.Pointer(&lpDeviceInfo.sSerialNumber[0]))))
	}
	return int(loginhandle)
}

// Log out API
func Logout(loginhandle int) bool {
	log.Println("Enter Logout")
	if 0 == C.CLIENT_Logout(C.long(loginhandle)) {
		log.Println("CLIENT_Logout fail")
		return false
	}
	return true

}

// High Level Login API
func LoginWithHighLevelSecurity(stInParam *NET_IN_LOGIN_WITH_HIGHLEVEL_SECURITY, stOutParam *NET_OUT_LOGIN_WITH_HIGHLEVEL_SECURITY) int {
	log.Println("Enter LoginWithHighLevelSecurity......")

	var lpInParmaInfo *C.struct_tagNET_IN_LOGIN_WITH_HIGHLEVEL_SECURITY = (*C.struct_tagNET_IN_LOGIN_WITH_HIGHLEVEL_SECURITY)(unsafe.Pointer(stInParam))
	(*lpInParmaInfo).dwSize = (C.uint)(unsafe.Sizeof(C.struct_tagNET_IN_LOGIN_WITH_HIGHLEVEL_SECURITY{}))
	var lPOutParmaInfo *C.struct_tagNET_OUT_LOGIN_WITH_HIGHLEVEL_SECURITY = (*C.struct_tagNET_OUT_LOGIN_WITH_HIGHLEVEL_SECURITY)(unsafe.Pointer(stOutParam))
	(*lPOutParmaInfo).dwSize = (C.uint)(unsafe.Sizeof(C.struct_tagNET_OUT_LOGIN_WITH_HIGHLEVEL_SECURITY{}))
	loginhandle := C.CLIENT_LoginWithHighLevelSecurity(lpInParmaInfo, lPOutParmaInfo)
	log.Println("nPort:", lpInParmaInfo.nPort)
	log.Println("szIP:", C.GoString((*C.char)(unsafe.Pointer(&lpInParmaInfo.szIP[0]))))
	log.Println("szUserName:", C.GoString((*C.char)(unsafe.Pointer(&lpInParmaInfo.szUserName[0]))))
	if loginhandle != 0 {
		log.Println("CLIENT_LoginWithHighLevelSecurity success")
	}

	log.Println("loginhandle=", loginhandle)
	return int(loginhandle)
}

// 主动注册功能,启动服务；nTimeout参数已无效(默认为设备断线后SDK内部登出)
func ListenServer(ip string, port uint16, nTimeout int, dwUserData ObjectId /*IF_fServiceCallBack*/) int {
	if "" == ip || dwUserData.IsNil() {
		log.Println("ip is empty or dwUserData is nil")
		return 0
	}
	ipTmp := C.CString(ip)
	ret := C.CLIENT_ListenServer(ipTmp, C.ushort(port), C.int(nTimeout), C.fServiceCallBack(C.export_fServiceCallBack), C.long(dwUserData))
	C.free(unsafe.Pointer(ipTmp))
	return int(ret)
}

// 停止服务
func StopListenServer(lServerHandle int) bool {
	if 0 == C.CLIENT_StopListenServer(C.long(lServerHandle)) {
		return false
	}
	return true
}

// 向设备订阅报警--扩展
func StartListenEx(lLoginID int) bool {
	log.Println(unsafe.Sizeof(ALARM_MOTIONDETECT_INFO{}), unsafe.Sizeof(C.struct_tagALARM_MOTIONDETECT_INFO{}))
	if 0 == C.CLIENT_StartListenEx(C.long(lLoginID)) {
		return false
	}
	return true
}

// 停止订阅报警
func StopListen(lLoginID int) bool {
	if 0 == C.CLIENT_StopListen(C.long(lLoginID)) {
		return false
	}
	return true
}

// 设置报警回调函数
func SetDVRMessCallBack(dwUser ObjectId /*IF_fMessCallBack*/) {
	if dwUser.IsNil() {
		log.Println("dwUser is nil")
		return
	}
	C.CLIENT_SetDVRMessCallBack(C.fMessCallBack(C.export_fMessCallBack), C.long(dwUser))
}

func (t EM_LOGIN_SPAC_CAP_TYPE) toString() string {

	switch t {

	case EM_LOGIN_SPEC_CAP_TCP:

		return "EM_LOGIN_SPEC_CAP_TCP"
	}
	return ""
}

func RealLoadPictureEx(lLoginID LLONG, nChannelID int32, dwAlarmType DWORD,
	bNeedPicFile BYTE, cbAnalyzerData CallBack_fAnalyzerDataCallBack) LLONG {
	var userdata *LLONG = (*LLONG)(unsafe.Pointer(&cbAnalyzerData))

	lAnalyzerHandle := C.CLIENT_RealLoadPictureEx(
		C.long(lLoginID),
		C.int(nChannelID),
		C.uint(dwAlarmType),
		C.int(bNeedPicFile),
		C.fAnalyzerDataCallBack(C.export_fAnalyzerDataCallBack),
		C.long(*userdata),
		unsafe.Pointer(uintptr(0)))

	if lAnalyzerHandle != 0 {
		log.Println("CLIENT_RealLoadPictureEx success")
	}

	log.Println("lAnalyzerHandle=", lAnalyzerHandle)
	return LLONG(lAnalyzerHandle)
}

// stop RealLoadPicture
func StopLoadPic(lAnalyzerHandle LLONG) bool {
	log.Println("Enter StopLoadPic")
	if 0 == C.CLIENT_StopLoadPic(C.long(lAnalyzerHandle)) {
		log.Println("CLIENT_StopLoadPic fail")
		return false
	}
	return true
}

// Open log function, user malloc memory of pstLogPrintInfo
func LogOpen(pstLogPrintInfo *LOG_SET_PRINT_INFO) bool {
	if nil == pstLogPrintInfo {
		log.Println("pstLogPrintInfo is nil")
		return false
	}

	var lpInParmaInfo *C.struct_tagLogSetPrintInfo = (*C.struct_tagLogSetPrintInfo)(unsafe.Pointer(pstLogPrintInfo))
	lpInParmaInfo.dwSize = C.uint(unsafe.Sizeof(C.struct_tagLogSetPrintInfo{}))

	Blog := C.CLIENT_LogOpen(lpInParmaInfo)
	log.Println("bSetPrintStrategy is :", C.int(lpInParmaInfo.bSetPrintStrategy))
	log.Println("nPrintStrategy is :", lpInParmaInfo.nPrintStrategy)
	if Blog != 0 {
		log.Println("CLIENT_LogOpen success")
		return true
	}
	return false

}

// Log close API
func LogClose() bool {
	log.Println("Enter LogClose....")
	Blog := C.CLIENT_LogClose()
	if Blog != 0 {
		log.Println("CLIENT_LogClose success")
		return true
	}
	return false
}

// asynchronism search device
func StartSearchDevicesEx(pInBuf *NET_IN_STARTSERACH_DEVICE, pOutBuf *NET_OUT_STARTSERACH_DEVICE) LLONG {
	log.Println("Enter StartSearchDeviceEx...")

	var pstInBuf *C.struct_tagNET_IN_STARTSERACH_DEVICE = (*C.struct_tagNET_IN_STARTSERACH_DEVICE)(unsafe.Pointer(pInBuf))
	(*pstInBuf).dwSize = (C.uint)(unsafe.Sizeof(C.struct_tagNET_IN_STARTSERACH_DEVICE{}))

	var pstOutBuf *C.struct_tagNET_OUT_STARTSERACH_DEVICE = (*C.struct_tagNET_OUT_STARTSERACH_DEVICE)(unsafe.Pointer(pOutBuf))
	(*pstOutBuf).dwSize = (C.uint)(unsafe.Sizeof(C.struct_tagNET_OUT_STARTSERACH_DEVICE{}))

	pstInBuf.pUserData = unsafe.Pointer(uintptr(*((*int)(unsafe.Pointer(&pInBuf.ST_pUserData)))))

	pstInBuf.cbSearchDevices = C.fSearchDevicesCBEx(C.export_fSearchDevicesCBEx)

	lSearchHandle := C.CLIENT_StartSearchDevicesEx(pstInBuf, pstOutBuf)
	if lSearchHandle != 0 {
		log.Println("CLIENT_StartSearchDevicesEx success")
	}
	log.Println("lSearchHandle=", lSearchHandle)
	log.Println("dwSize is :", pstInBuf.dwSize)
	GetLastError()

	return LLONG(lSearchHandle)

}

// Stop asynchronism search IPC, NVS and etc in LAN
func StopSearchDevices(lSearchHandle LLONG) bool {
	log.Println("Enter StopSearchDevices...")
	ret := C.CLIENT_StopSearchDevices(C.long(lSearchHandle))
	if 0 == ret {
		return false
	}
	return true
}

// Snapshot request--extensive
func SnapPictureEx(lLoginID int, par *SNAP_PARAMS, reserved unsafe.Pointer) bool {
	if nil == par {
		log.Println("par is nil")
		return false
	}

	// log.Println("SnapPictureEx size", unsafe.Sizeof(SNAP_PARAMS{}), unsafe.Sizeof(C.struct__snap_param{}))
	parLocal := (*C.struct__snap_param)(unsafe.Pointer(par))
	ret := C.CLIENT_SnapPictureEx(C.long(lLoginID), parLocal, (*C.int)(reserved))
	if 0 == ret {
		return false
	}
	return true
}

// Set snapshot callback function
func SetSnapRevCallBack(dwUser ObjectId /*IF_fSnapRev*/) {
	if dwUser.IsNil() {
		log.Println("dwUser is nil")
		return
	}
	C.CLIENT_SetSnapRevCallBack(C.fSnapRev(C.export_fSnapRev), C.long(dwUser))
}

func DownloadByTimeEx(lLoginID int, nChannelId int, nRecordFileType EM_QUERY_RECORD_TYPE, tmStart *NET_TIME, tmEnd *NET_TIME, sSavedFileName string, dwUserData ObjectId /*IF_fTimeDownLoadPosCallBack*/, dwDataUser ObjectId /*IF_fDataCallBack*/, pReserved unsafe.Pointer) int {
	if dwUserData.IsNil() || dwDataUser.IsNil() {
		log.Println("dwUserData or dwDataUser is nil")
		return 0
	}

	tmStartTmp := (*C.NET_TIME)(unsafe.Pointer(tmStart))
	tmEndTmp := (*C.struct_tagNET_TIME)(unsafe.Pointer(tmEnd))
	fileNameTmp := C.CString(sSavedFileName)
	ret := C.CLIENT_DownloadByTimeEx(C.long(lLoginID), C.int(nChannelId), C.int(nRecordFileType), tmStartTmp, tmEndTmp, fileNameTmp,
		C.fTimeDownLoadPosCallBack(C.export_fTimeDownLoadPosCallBack), C.long(dwUserData), C.fDataCallBack(C.export_fDataCallBack), C.long(dwDataUser), pReserved)
	C.free(unsafe.Pointer(fileNameTmp))
	return int(ret)
}

// Init account
func InitDevAccount(stInitAccountIn *NET_IN_INIT_DEVICE_ACCOUNT, stInitAccountOut *NET_OUT_INIT_DEVICE_ACCOUNT, dwWaitTime DWORD, szLocalIp string) bool {
	log.Println("Enter InitDevAccount......")
	var pInitAccountIn *C.struct_tagNET_IN_INIT_DEVICE_ACCOUNT = (*C.struct_tagNET_IN_INIT_DEVICE_ACCOUNT)(unsafe.Pointer(stInitAccountIn))
	(*pInitAccountIn).dwSize = (C.uint)(unsafe.Sizeof(C.struct_tagNET_IN_INIT_DEVICE_ACCOUNT{}))

	var pInitAccountOut *C.struct_tagNET_OUT_INIT_DEVICE_ACCOUNT = (*C.struct_tagNET_OUT_INIT_DEVICE_ACCOUNT)(unsafe.Pointer(stInitAccountOut))
	(*pInitAccountOut).dwSize = (C.uint)(unsafe.Sizeof(C.struct_tagNET_OUT_INIT_DEVICE_ACCOUNT{}))

	szLocalIpTmp := C.CString(szLocalIp)
	ret := C.CLIENT_InitDevAccount(pInitAccountIn, pInitAccountOut, C.uint(dwWaitTime), szLocalIpTmp)
	C.free(unsafe.Pointer(szLocalIpTmp))
	if 0 == ret {
		return false
	}
	return true
}

// search device ip cross VLAN
// search command will be send once when interface called
func SearchDevicesByIPs(stIpSearchInfo *DEVICE_IP_SEARCH_INFO, dwUserData IF_fSearchDevicesCB, szLocalIp string, dwWaitTime DWORD) BOOL {
	log.Println("Enter SearchDevicesByIPs......")

	var pIpSearchInfo *C.struct_tagDEVICE_IP_SEARCH_INFO = (*C.struct_tagDEVICE_IP_SEARCH_INFO)(unsafe.Pointer(stIpSearchInfo))
	(*pIpSearchInfo).dwSize = (C.uint)(unsafe.Sizeof(C.struct_tagDEVICE_IP_SEARCH_INFO{}))

	dwUserLocal := *(*int)(unsafe.Pointer(&dwUserData))

	var lSearchHandle C.int
	if len(szLocalIp) == 0 {
		lSearchHandle = C.CLIENT_SearchDevicesByIPs(pIpSearchInfo, C.fSearchDevicesCB(C.export_fSearchDevicesCB), C.long(dwUserLocal), nil, C.uint(dwWaitTime))
	} else {
		szLocalIpTmp := C.CString(szLocalIp)
		lSearchHandle = C.CLIENT_SearchDevicesByIPs(pIpSearchInfo, C.fSearchDevicesCB(C.export_fSearchDevicesCB), C.long(dwUserLocal), szLocalIpTmp, C.uint(dwWaitTime))
		C.free(unsafe.Pointer(szLocalIpTmp))
	}

	if lSearchHandle != 0 {
		log.Println("CLIENT_SearchDevicesByIPs success")
	}
	log.Println("lSearchHandle=", lSearchHandle)
	GetLastError()

	return BOOL(lSearchHandle)
}

func StopDownload(lFileHandle int) bool {
	if C.CLIENT_StopDownload(C.long(lFileHandle)) == 0 {
		return false
	}
	return true
}

func SetDeviceMode(lLoginID int, emType EM_USEDEV_MODE, pValue unsafe.Pointer) bool {
	if C.CLIENT_SetDeviceMode(C.long(lLoginID), C.EM_USEDEV_MODE(C.int(emType)), pValue) == 0 {
		return false
	}
	return true
}

// Get NMP Device info
func GetNMPDevInfo(lLoginID int, stInParam *NET_IN_GET_NMPDEV_INFO, stOutParam *NET_OUT_GET_NMPDEV_INFO, nWaitTime int32) bool {
	log.Println("Enter GetNMPDevInfo......")
	if nil == stInParam {
		log.Println("stInParam is nil")
		return false
	}
	if nil == stOutParam {
		log.Println("stOutParam is nil")
		return false
	}
	lpInParmaInfo := (*C.struct_tagNET_IN_GET_NMPDEV_INFO)(unsafe.Pointer(stInParam))
	(*lpInParmaInfo).dwSize = (C.uint)(unsafe.Sizeof(C.struct_tagNET_IN_GET_NMPDEV_INFO{}))

	lPOutParmaInfo := (*C.struct_tagNET_OUT_GET_NMPDEV_INFO)(unsafe.Pointer(stOutParam))
	(*lPOutParmaInfo).dwSize = (C.uint)(unsafe.Sizeof(C.struct_tagNET_OUT_GET_NMPDEV_INFO{}))

	ret := C.CLIENT_GetNMPDevInfo(C.long(lLoginID), lpInParmaInfo, lPOutParmaInfo, C.int(nWaitTime))
	if 0 == ret {
		return false
	}
	return true

}

// Set log in network environment
func SetNetworkParam(stNetParam *NET_PARAM) {
	log.Println("Enter SetNetworkParam......")
	if nil == stNetParam {
		log.Println("stNetParam is nil")
	}

	pNetParam := (*C.struct_tagNET_PARAM)(unsafe.Pointer(stNetParam))

	C.CLIENT_SetNetworkParam(pNetParam)
}

// Set optimize mode,user malloc and free pParam's memory,please refer to corresponding structure of emType
func SetOptimizeMode(emType EM_OPTIMIZE_TYPE, pParam unsafe.Pointer) bool {
	if C.CLIENT_SetOptimizeMode(C.EM_OPTIMIZE_TYPE(C.int(emType)), pParam) == 0 {
		return false
	}
	return true
}

func StrZ(addr uintptr) string {
	return C.GoString((*C.char)(unsafe.Pointer(addr)))
}

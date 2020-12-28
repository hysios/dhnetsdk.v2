package netsdk

// #cgo CFLAGS: -I .
// #cgo LDFLAGS: -L . -ldhnetsdk
// #include <stdio.h>
// #include <stdlib.h>
// #include "dhnetsdk.h"
//
// extern void export_fDisConnect(long lLoginID, char * pchDVRIP, int nDVRPort, long dwUser);
// extern int export_fServiceCallBack(long lHandle, char * pIp, unsigned short wPort, int lCommand, long pParam, int dwParamLen, long dwUserData);
// extern int export_fMessCallBack(int lCommand, long lLoginID, char * pBuf, int dwBufLen, char * pchDVRIP, int nDVRPort, long dwUser);
// extern int export_fSDKLogCallBack(char * szLogBuffer, uint nLogSize, long dwUser);
// extern int export_fAnalyzerDataCallBack(long lAnalyzerHandle, unsigned int dwAlarmType, long pAlarmInfo, long pBuffer,unsigned int dwBufSize, long dwUser, int nSequence,long reserved);
// extern int export_fAnalyzerDataCallBack2(long lAnalyzerHandle, unsigned int dwAlarmType, long pAlarmInfo, long pBuffer,unsigned int dwBufSize, long dwUser, int nSequence,long reserved);
// extern void export_fSearchDevicesCBEx(long lSearchHandle, DEVICE_NET_INFO_EX2 *pDevNetInfo, long pUserData);
// extern void export_fSnapRev(long lLoginID, unsigned char *pBuf, unsigned int RevLen, unsigned int EncodeType, unsigned int CmdSerial, long dwUser);
// extern void export_fSearchDevicesCB(DEVICE_NET_INFO_EX *pDevNetInfo, long pUserData);
// extern void export_fTimeDownLoadPosCallBack(long lPlayHandle, unsigned int dwTotalSize, int dwDownLoadSize, int index, NET_RECORDFILE_INFO recordfileinfo, long dwUser);
// extern void export_fTimeDownLoadPosCallBack2(long lPlayHandle, unsigned int dwTotalSize, int dwDownLoadSize, int index, NET_RECORDFILE_INFO recordfileinfo, long dwUser);
// extern int export_fDataCallBack(long lRealHandle, unsigned int dwDataType, unsigned char * pBuffer, unsigned int dwBufSize, long dwUser);
// extern int export_fDataCallBack2(long lRealHandle, unsigned int dwDataType, unsigned char * pBuffer, unsigned int dwBufSize, long dwUser);
// extern void export_fDisconnect2(long dwUser, char *pchDVRIP, LONG nDVRPort);
//
// void CALLBACK cDisConnectFunc(LLONG lLoginID, char *pchDVRIP, LONG nDVRPort, LDWORD dwUser) {
//     if (0 != dwUser)
//     {
//         export_fDisconnect2(dwUser, pchDVRIP, nDVRPort);
//     }
// }
//
// void CALLBACK cReConnectFunc(LLONG lLoginID, char *pchDVRIP, LONG nDVRPort, LDWORD dwUser)
// {
//     if (0 != dwUser)
//     {
//         goReConnect(dwUser, pchDVRIP, nDVRPort);
//     }
// }
//
// BOOL CALLBACK cMessCallBack(LONG lCommand, LLONG lLoginID, char *pBuf, DWORD dwBufLen, char *pchDVRIP, LONG nDVRPort, LDWORD dwUser)
// {
//     if(0 == dwUser)
//     {
//         return FALSE;
//     }
// 	   return goDvrMessage(dwUser, lCommand, pBuf, dwBufLen, pchDVRIP, nDVRPort);
// }
//
import "C"

import (
	"errors"
	"log"
	"time"
	"unsafe"

	"github.com/mattn/go-pointer"
)

type DhAlarmType int

const (
	// 报警类型,对应CLIENT_StartListen接口
	DhCommAlarm         DhAlarmType = C.DH_COMM_ALARM          // 常规报警(包括外部报警,视频丢失,动态检测)
	DhShelterAlarm      DhAlarmType = C.DH_SHELTER_ALARM       // 视频遮挡报警
	DhDiskFullAlarm     DhAlarmType = C.DH_DISK_FULL_ALARM     // 硬盘满报警
	DhDiskErrorAlarm    DhAlarmType = C.DH_DISK_ERROR_ALARM    // 硬盘故障报警
	DhSoundDetectAlarm  DhAlarmType = C.DH_SOUND_DETECT_ALARM  // 音频检测报警
	DhAlarmDecoderAlarm DhAlarmType = C.DH_ALARM_DECODER_ALARM // 报警解码器报警
)

type (
	ReconnectFunc    func(client *Client, ip string, port int)
	DisconnectFunc   func(ip string, port int)
	DVRMessageFunc   func(client *Client, cmd DhAlarmType, buf []byte, ip string, port int) bool
	PictureExFunc    func(client *Client, AlarmType EventIvs, alarmInfo interface{}, frame []byte, seq int) int
	DownloadPosFunc  func(userdata interface{}, total int, download int, index int, info NET_RECORDFILE_INFO)
	DataCallbackFunc func(userdata interface{}, typ int, buf []byte)
)

type (
	DisconnectVisitor struct {
		Callback DisconnectFunc
	}

	ReconnectVisitor struct {
		Client   *Client
		Callback ReconnectFunc
	}

	DrvMessageVisitor struct {
		Client   *Client
		Callback DVRMessageFunc
	}

	TimeDownloadPosVisitor struct {
		UserData interface{}
		Step     int
		Callback DownloadPosFunc
	}

	DataCallbackVisitor struct {
		UserData interface{}
		Callback DataCallbackFunc
	}
)

type PictureVisitor struct {
	Client   *Client
	Callback PictureExFunc
}

// func InitEx(cb CallBack_fDisConnect, stuInfo *NETSDK_INIT_PARAM) bool {
// 	if nil == stuInfo {
// 		log.Println("stuInfo is nil")
// 		return false
// 	}

// 	var lpInitParam = C.struct_tagNETSDK_INIT_PARAM{}
// 	lpInitParam.nThreadNum = C.int(stuInfo.ST_nThreadNum)

// 	var userdata *LLONG = (*LLONG)(unsafe.Pointer(&cb))
// 	log.Println("userData=", *userdata)

// 	ret := C.CLIENT_InitEx(C.fDisConnect(C.export_fDisConnect), C.long(*userdata), (*C.struct_tagNETSDK_INIT_PARAM)(unsafe.Pointer(&lpInitParam)))
// 	if ret != 0 {
// 		log.Println(" call CLIENT_InitEx success")
// 		return true
// 	}
// 	return false
// }

func InitEx(callback DisconnectFunc) error {
	var (
		v = DisconnectVisitor{
			Callback: callback,
		}
	)

	p := pointer.Save(v)
	ret := C.CLIENT_Init(C.fDisConnect(C.cDisConnectFunc), (C.long)(uintptr(p)))
	if ret > 0 {
		return Err(GetLastError())
	}
	return nil
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
func LoginEx2(pchDVRIP string, wDVRPort int, UserName string, Password string, mode EM_LOGIN_SPAC_CAP_TYPE) int {
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
		C.EM_LOGIN_SPAC_CAP_TYPE(mode),
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

func DownloadByTimeEx2(lLoginID int, nChannelId int, nRecordFileType EM_QUERY_RECORD_TYPE, tmStart *NET_TIME, tmEnd *NET_TIME, sSavedFileName string, dwUserData ObjectId /*IF_fTimeDownLoadPosCallBack*/, dwDataUser ObjectId /*IF_fDataCallBack*/, pReserved unsafe.Pointer) int {
	if dwUserData.IsNil() || dwDataUser.IsNil() {
		log.Println("dwUserData or dwDataUser is nil")
		return 0
	}

	tmStartTmp := (*C.NET_TIME)(unsafe.Pointer(tmStart))
	tmEndTmp := (*C.struct_tagNET_TIME)(unsafe.Pointer(tmEnd))
	fileNameTmp := C.CString(sSavedFileName)
	ret := C.CLIENT_DownloadByTimeEx(C.long(lLoginID), C.int(nChannelId), C.int(nRecordFileType), tmStartTmp, tmEndTmp, fileNameTmp,
		C.fTimeDownLoadPosCallBack(C.export_fTimeDownLoadPosCallBack2), C.long(dwUserData), C.fDataCallBack(C.export_fDataCallBack), C.long(dwDataUser), pReserved)
	C.free(unsafe.Pointer(fileNameTmp))
	return int(ret)
}

type download struct {
}

type userdata struct {
}

func DownloadByTimeEx(lLoginID int, nChannelId int, nRecordFileType EM_QUERY_RECORD_TYPE, startTime time.Time, duration time.Duration, sSavedFileName string, userdata interface{}, callback DownloadPosFunc) (uint64, error) {
	var (
		downloadVisitor = &TimeDownloadPosVisitor{
			UserData: userdata,
			Callback: callback,
		}
		callbackVisitor = &DataCallbackVisitor{
			UserData: userdata,
			Callback: func(suerdata interface{}, typ int, buf []byte) {},
		}
	)

	downinfo := pointer.Save(downloadVisitor)
	datacb := pointer.Save(callbackVisitor)

	tmStart := Time2nt(startTime)
	tmEnd := Time2nt(startTime.Add(duration))
	tmStartTmp := (*C.NET_TIME)(unsafe.Pointer(&tmStart))
	tmEndTmp := (*C.NET_TIME)(unsafe.Pointer(&tmEnd))
	fileNameTmp := C.CString(sSavedFileName)
	ret := C.CLIENT_DownloadByTimeEx(
		C.long(lLoginID),
		C.int(nChannelId),
		C.int(nRecordFileType),
		tmStartTmp,
		tmEndTmp,
		fileNameTmp,
		C.fTimeDownLoadPosCallBack(C.export_fTimeDownLoadPosCallBack2),
		C.long(uintptr(downinfo)),
		C.fDataCallBack(C.export_fDataCallBack2),
		C.long(uintptr(datacb)),
		nil,
	)
	C.free(unsafe.Pointer(fileNameTmp))
	if ret == 0 {
		return 0, Err(GetLastError())
	}
	return uint64(ret), nil
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

func SetAutoReconnect(callback ReconnectFunc) error {
	var v = ReconnectVisitor{
		Client:   nil,
		Callback: callback,
	}

	p := pointer.Save(v)
	C.CLIENT_SetAutoReconnect(C.fHaveReConnect(C.cReConnectFunc), (C.long)(uintptr(p)))
	return nil
}

func SetConnectTime(timeout time.Duration, times int) error {
	C.CLIENT_SetConnectTime(C.int(timeout/time.Millisecond), C.int(times))
	return nil
}

func (client *Client) SetAutoReconnect(callback ReconnectFunc) error {
	var v = ReconnectVisitor{
		Client:   client,
		Callback: callback,
	}
	if client.reconnectVisitorp != nil {
		pointer.Unref(client.reconnectVisitorp)
	}
	p := pointer.Save(v)
	client.reconnectVisitorp = p
	// defer pointer.Unref(p)
	C.CLIENT_SetAutoReconnect(C.fHaveReConnect(C.cReConnectFunc), (C.long)(uintptr(p)))
	return nil
}

func (client *Client) SetDVRMessCallBack(callback DVRMessageFunc) error {
	var v = DrvMessageVisitor{
		Client:   client,
		Callback: callback,
	}

	if client.messageVisitorp != nil {
		pointer.Unref(client.messageVisitorp)
	}
	p := pointer.Save(&v)
	client.messageVisitorp = p

	C.CLIENT_SetDVRMessCallBack(C.fMessCallBack(C.cMessCallBack), (C.long)(uintptr(p)))
	return nil
}

func Reboot(lloginID int) error {
	C.CLIENT_ControlDevice(C.long(lloginID), C.CtrlType(0), nil, 5000)
	return nil
}

func ResponseDevReg(sn string, ip string, port int) error {
	ret := C.CLIENT_ResponseDevReg(C.CString(sn), C.CString(ip), C.ushort(port), C.int(1))
	if ret > 0 {
		log.Printf("成功")
		return nil
	}

	return errors.New("response error")
	// panic("nonimplement")
}

package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"time"
	"unsafe"

	"github.com/hysios/dhnetsdk.v2/netsdk"
	"github.com/yudai/pp"
)

func Test_fDisConnect(lLoginID netsdk.LLONG, pchDVRIP string, nDVRPort int, dwUser netsdk.LLONG) {
	fmt.Println("DisConnet")
}

func TestfAnalyzerDataCallBack(lAnalyzerHandle netsdk.LLONG, dwAlarmType netsdk.DWORD, pAlarmInfo netsdk.LLONG,
	pBuffer netsdk.LLONG, dwBufSize netsdk.DWORD, dwUser netsdk.LLONG, nSequence int32, reserved netsdk.LLONG) {

	switch dwAlarmType {
	case netsdk.EVENT_IVS_TRAFFICJUNCTION:
		fmt.Println("交通路口事件---------------------begin---------------------------------------")
		alarmInfo := (*netsdk.DEV_EVENT_TRAFFICJUNCTION_INFO)(unsafe.Pointer(uintptr(pAlarmInfo)))
		fmt.Println("名称---------------", string(alarmInfo.ST_szName[:]))
		fmt.Println("EventID-----------", alarmInfo.ST_nEventID)
		fmt.Println("车道---------------", alarmInfo.ST_nLane)
		fmt.Println("车牌类型------------", string(alarmInfo.ST_stTrafficCar.ST_szPlateType[:]))
		fmt.Println("车牌颜色------------", string(alarmInfo.ST_stTrafficCar.ST_szPlateColor[:]))
		fmt.Println("车牌---------------", string(alarmInfo.ST_stTrafficCar.ST_szPlateNumber[:]))
		fmt.Println("ObjectID----------", alarmInfo.ST_stuObject.ST_nObjectID)
		fmt.Println("GroupID-----------", alarmInfo.ST_stuFileInfo.ST_nGroupId)

		globalfilename := fmt.Sprintf("%04d%02d%02d-%02d%02d%02d%04d-%d_%s",
			alarmInfo.ST_UTC.ST_dwYear,
			alarmInfo.ST_UTC.ST_dwMonth,
			alarmInfo.ST_UTC.ST_dwDay,
			alarmInfo.ST_UTC.ST_dwHour,
			alarmInfo.ST_UTC.ST_dwMinute,
			alarmInfo.ST_UTC.ST_dwSecond,
			alarmInfo.ST_UTC.ST_dwMillisecond,
			alarmInfo.ST_stuObject.ST_nObjectID,
			"big.jpg")

		// 当上报的图有车牌时，显示车牌小图
		if alarmInfo.ST_stuObject.ST_bPicEnble == 1 {
			//m_ctlPlatePic.Load(pBuffer + pTrafficJunctonInfo->stuObject.stPicInfo.dwOffSet, pTrafficJunctonInfo->stuObject.stPicInfo.dwFileLenth);
			file, err := os.OpenFile(globalfilename, os.O_CREATE|os.O_WRONLY, 0666)
			if err != nil {
				fmt.Println("Error OpenFile ")
			}
			defer file.Close()

			var buf []byte
			data := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
			data.Data = uintptr(unsafe.Pointer(uintptr(pBuffer)))
			data.Len = int(alarmInfo.ST_stuObject.ST_stPicInfo.ST_dwOffSet)
			data.Cap = int(alarmInfo.ST_stuObject.ST_stPicInfo.ST_dwOffSet)
			n, err := file.Write(buf)
			_ = n

			platefilename := fmt.Sprintf("%04d%02d%02d-%02d%02d%02d%04d-%d_%s",
				alarmInfo.ST_UTC.ST_dwYear,
				alarmInfo.ST_UTC.ST_dwMonth,
				alarmInfo.ST_UTC.ST_dwDay,
				alarmInfo.ST_UTC.ST_dwHour,
				alarmInfo.ST_UTC.ST_dwMinute,
				alarmInfo.ST_UTC.ST_dwSecond,
				alarmInfo.ST_UTC.ST_dwMillisecond,
				alarmInfo.ST_stuObject.ST_nObjectID,
				"plate.jpg")
			Platefile, err2 := os.OpenFile(platefilename, os.O_CREATE|os.O_WRONLY, 0666)
			if err2 != nil {
				fmt.Println("Error plate jpg OpenFile ", err2)
			}
			defer Platefile.Close()
			var PlateBuf []byte

			var dwOffSet netsdk.DWORD = alarmInfo.ST_stuObject.ST_stPicInfo.ST_dwOffSet
			dwFileLenth := alarmInfo.ST_stuObject.ST_stPicInfo.ST_dwFileLenth
			Platedata := (*reflect.SliceHeader)(unsafe.Pointer(&PlateBuf))

			//xx := unsafe.Pointer(uintptr(pBuffer + dwOffSet))
			Platedata.Data = uintptr(pBuffer + netsdk.LLONG(dwOffSet))
			Platedata.Len = int(dwFileLenth)
			Platedata.Cap = int(dwFileLenth)
			n2, err2 := Platefile.Write(PlateBuf)
			_ = n2
			fmt.Println("platefilename-----", platefilename)
		} else {
			// 保存图片
			file, err := os.OpenFile(globalfilename, os.O_CREATE|os.O_WRONLY, 0666)
			if err != nil {
				fmt.Println("Error OpenFile ")
			}
			defer file.Close()
			var buf []byte
			data := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
			data.Data = uintptr(unsafe.Pointer(uintptr(pBuffer)))
			data.Len = int(dwBufSize)
			data.Cap = int(dwBufSize)
			n, err := file.Write(buf)
			_ = n
		}
		fmt.Println("globalfilename----", globalfilename)
		break
	case netsdk.EVENT_ALARM_MOTIONDETECT:
		fmt.Println("移动侦测事件。。。。。。")
		break
	case netsdk.EVENT_IVS_TRAFFIC_PARKING:
		alarmInfo := (*netsdk.DEV_EVENT_TRAFFIC_PARKING_INFO)(unsafe.Pointer(uintptr(pAlarmInfo)))
		fmt.Println("名称---------------", string(alarmInfo.ST_szName[:]))
		fmt.Println("EventID-----------", alarmInfo.ST_nEventID)
		fmt.Println("车道---------------", alarmInfo.ST_nLane)
		fmt.Println("车牌类型------------", string(alarmInfo.ST_stTrafficCar.ST_szPlateType[:]))
		fmt.Println("车牌颜色------------", string(alarmInfo.ST_stTrafficCar.ST_szPlateColor[:]))
		fmt.Println("车牌---------------", string(alarmInfo.ST_stTrafficCar.ST_szPlateNumber[:]))
		fmt.Println("ObjectID----------", alarmInfo.ST_stuObject.ST_nObjectID)
		fmt.Println("GroupID-----------", alarmInfo.ST_stuFileInfo.ST_nGroupId)
		pp.Printf("alarminfo %v", alarmInfo)
		log.Printf("car %#v", alarmInfo.ST_stTrafficCar)
	default:
		fmt.Println("dwAlarmType:", dwAlarmType)
	}

}

func main() {
	initParam := netsdk.NETSDK_INIT_PARAM{}
	bRet := netsdk.InitEx(Test_fDisConnect, &initParam)
	if false == bRet {
		fmt.Println("Init NetSDK failed")
		return
	}
	fmt.Println("Init NetSDK success")
	defer netsdk.Cleanup()

	var (
		ip             = "192.168.1.190"
		port     int32 = 37777
		username       = "admin"
		passwd         = "admin123"
	)

	fmt.Println("===>Please input device ip, port, username, password, e.g. 192.168.1.1 37777 abc 123456")
	fmt.Scanln(&ip, &port, &username, &passwd)

	stLoginIn := netsdk.NET_IN_LOGIN_WITH_HIGHLEVEL_SECURITY{}
	copy(stLoginIn.ST_szIP[:], []byte(ip))
	stLoginIn.ST_nPort = port
	copy(stLoginIn.ST_szUserName[:], []byte(username))
	copy(stLoginIn.ST_szPassword[:], []byte(passwd))

	stLoginOut := netsdk.NET_OUT_LOGIN_WITH_HIGHLEVEL_SECURITY{}
	lhandle := netsdk.LoginWithHighLevelSecurity(&stLoginIn, &stLoginOut)
	if 0 == lhandle {
		fmt.Printf("LoginWithHighLevelSecurity failed, 0x%x\n", netsdk.GetLastError())
		return
	}
	fmt.Println("LoginWithHighLevelSecurity success")
	defer netsdk.Logout(lhandle)

	SerialNum := stLoginOut.ST_stDeviceInfo.ST_sSerialNumber[:]
	fmt.Println("Serial Num:", string(SerialNum))

	var nChannelID int32 = 0
	fmt.Println("===>Please input channelID(begin with 0):, e.g. 0")
	// fmt.Scanln(&nChannelID)

	var nSleepTime int32 = 100000
	fmt.Println("*****************************************************************")
	fmt.Println("===>Please input Sleep Times(wait time to recv event):, e.g. 10")
	fmt.Println("*****************************************************************")
	fmt.Scanln(&nSleepTime)

	var dwAlarmType netsdk.DWORD = netsdk.DWORD(netsdk.EVENT_IVS_ALL)
	var bNeedPicFile netsdk.BYTE = 1
	lAnalyzerHandle := netsdk.RealLoadPictureEx(netsdk.LLONG(lhandle), nChannelID, dwAlarmType, bNeedPicFile, TestfAnalyzerDataCallBack)

	if lAnalyzerHandle == 0 {
		return
	}
	fmt.Println("RealLoadPictureEx success")
	defer netsdk.StopLoadPic(lAnalyzerHandle)

	time.Sleep(time.Duration(nSleepTime) * time.Second)

}

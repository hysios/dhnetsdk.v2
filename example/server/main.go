package main

import "C"

import (
	"flag"
	"fmt"
	"log"
	"sync"
	"unsafe"

	"github.com/hysios/dhnetsdk.v2/netsdk"
)

// implement interface IF_fServiceCallBack
type ST_fServiceCallBack struct {
	dwUserData int
}

func (p *ST_fServiceCallBack) Invoke(lHandle int, pIp string, wPort uint16, lCommand int, pParam uintptr, dwParamLen int) int {
	if netsdk.NET_DVR_SERIAL_RETURN == lCommand {
		sn := C.GoString((*C.char)(unsafe.Pointer(pParam)))
		fmt.Println("Device register to server:", lHandle, pIp, wPort, sn)
		// if deviceId == sn {
		// 	go commonAlarm(pIp, int(wPort), username, passwd, sn)
		// }
		// client, err := netsdk.Login(fmt.Sprintf("%s:%d", pIp, wPort), "admin", "admin123",
		// 	netsdk.LoginMode(netsdk.EM_LOGIN_SPEC_CAP_SERVER_CONN),
		// 	netsdk.LoginActive(sn))
		// if err != nil {
		// 	log.Printf("error %s\n", err)
		// }
		// log.Printf("client % #v\n", pretty.Formatter(client))
		// log.Printf("client serialNumber %s\n", string(client.DeviceInfo.ST_sSerialNumber[:]))
	}

	return 0
}

var (
	ip   = flag.String("ip", "", "bind listen ip")
	port = flag.Int("port", 7000, "bind listen port")
)

func main() {
	flag.Parse()
	if err := netsdk.InitEx(func(ip string, port int) {
		log.Printf("disconnect ip %s:%d\n", ip, port)
	}); err != nil {
		log.Fatalf("init netsdk error %s", err)
	}
	fmt.Println("Init NetSDK success")
	defer netsdk.Cleanup()

	var (
		cbObj        = &ST_fServiceCallBack{12345}
		cbObjId      = netsdk.NewObjectId(cbObj) // need manual free
		serverHandle = netsdk.ListenServer(*ip, uint16(*port), 0, cbObjId)
	)
	// var serverHandle = netsdk.ListenServer(ip, port, 0, cbObjId)
	if 0 == serverHandle {
		fmt.Printf("ListenServer failed, %s\n", netsdk.Err(netsdk.GetLastError()))
		return
	}
	fmt.Println("ListenServer success")
	defer func() {
		if netsdk.StopListenServer(serverHandle) {
			fmt.Println("StopListenServer success")
		} else {
			fmt.Printf("StopListenServer failed, %s\n", netsdk.Err(netsdk.GetLastError()))
		}
	}()

	// fmt.Println("===>Please input event switch and time, unit S, e.g. true 60")
	// fmt.Scanln(&listening, &eventTime)

	// fmt.Println("===>Please input device id, username, password, e.g. ABCDEFG abc abc123456")
	// fmt.Scanln(&deviceId, &username, &passwd)
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}

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
	switch lCommand {
	case netsdk.NET_DVR_SERIAL_RETURN:
		sn := C.GoString((*C.char)(unsafe.Pointer(pParam)))
		log.Println("Device register to server:", lHandle, pIp, wPort, sn)
		// addr := fmt.Sprintf("%s:%d", "192.168.1.108", 37777)
		// client, err := netsdk.Login(addr, "admin", "admin123")
		// if err != nil {
		// 	log.Printf("login error %s", err)
		// }
		err := netsdk.ResponseDevReg(sn, pIp, int(wPort))
		if err != nil {
			log.Printf("response dev reg error %s", err)
			return 0
		}

		log.Printf("注册成功")
		var (
			inParam  netsdk.NET_IN_LOGIN_WITH_HIGHLEVEL_SECURITY
			outParam netsdk.NET_OUT_LOGIN_WITH_HIGHLEVEL_SECURITY
		)
		inParam.ST_dwSize = uint32(unsafe.Sizeof(inParam))
		copy(inParam.ST_szIP[:], []byte(pIp))
		inParam.ST_nPort = int32(wPort)
		copy(inParam.ST_szUserName[:], []byte("admin"))
		copy(inParam.ST_szPassword[:], []byte("admin123"))
		inParam.ST_emSpecCap = netsdk.EM_LOGIN_SPEC_CAP_SERVER_CONN
		inParam.ST_pCapParam = pParam
		loginId := netsdk.LoginWithHighLevelSecurity(&inParam, &outParam)
		log.Printf("loginId %d", loginId)
		// addr := fmt.Sprintf("%s:%d", pIp, wPort)
		// client, err := netsdk.Login(addr, "admin", "admin123")
		// if err != nil {
		// 	log.Printf("login error %s", err)
		// }
		// // client := netsdk.Client{LoginID: lHandle}

		// log.Printf("client %#v", client)
		// client.Reboot()
		// addr := fmt.Sprintf("%s:%d", pIp, )
		// netsdk.Login(pIp, int(wPort))
	case netsdk.NET_DEV_AUTOREGISTER_RETURN:
		sn := C.GoString((*C.char)(unsafe.Pointer(pParam)))
		log.Println("Device autoregister to server:", lHandle, pIp, wPort, sn)
	case netsdk.NET_DVR_DISCONNECT:
		sn := C.GoString((*C.char)(unsafe.Pointer(pParam)))
		log.Println("Device disconnect to server:", lHandle, pIp, wPort, sn)
	default:
		log.Printf("other command %d", lCommand)
	}

	return 1
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

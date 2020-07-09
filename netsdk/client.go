package netsdk

// #cgo CFLAGS: -I .
// #cgo LDFLAGS: -L . -ldhnetsdk

// #include <stdio.h>
// #include <stdlib.h>
// #include "dhnetsdk.h"
// extern int export_fAnalyzerDataCallBack2(long lAnalyzerHandle, unsigned int dwAlarmType, long pAlarmInfo, long pBuffer,unsigned int dwBufSize, long dwUser, int nSequence,long reserved);
import "C"

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unsafe"

	"github.com/mattn/go-pointer"
)

type Client struct {
	LoginID int

	realloadHandle int64
	subscribP      unsafe.Pointer
	DeviceInfo     NET_DEVICEINFO_Ex
}

func Init(cb DisconnectFunc) error {
	initParam := NETSDK_INIT_PARAM{}
	bRet := InitEx(func(lLoginID LLONG, pchDVRIP string, nDVRPort int, dwUser LLONG) {
		if cb != nil {
			cb(pchDVRIP, nDVRPort)
		}
	}, &initParam)
	if false == bRet {
		return fmt.Errorf("Init NetSDK failed")
	}
	return nil
}

func NewClient() *Client {
	return &Client{}
}

func Login(addr string, user, pass string, opts ...LogOptionFunc) (ncli *Client, err error) {
	var (
		port     int
		inParam  NET_IN_LOGIN_WITH_HIGHLEVEL_SECURITY
		outParam NET_OUT_LOGIN_WITH_HIGHLEVEL_SECURITY
	)

	for _, opt := range opts {
		opt(&inParam)
	}

	addrs := strings.SplitN(addr, ":", 2)
	if len(addrs) == 2 {
		addr = addrs[0]
		if port, err = strconv.Atoi(addrs[1]); err != nil {
			return nil, err
		}
	} else {
		return nil, ErrInvalidAddress
	}
	copy(inParam.ST_szIP[:], []byte(addr))
	inParam.ST_nPort = int32(port)
	copy(inParam.ST_szUserName[:], []byte(user))
	copy(inParam.ST_szPassword[:], []byte(pass))

	id := LoginWithHighLevelSecurity(&inParam, &outParam)

	ncli = &Client{
		LoginID:    id,
		DeviceInfo: outParam.ST_stDeviceInfo,
	}
	return ncli, nil
}

type LogOptionFunc func(inparam *NET_IN_LOGIN_WITH_HIGHLEVEL_SECURITY)

// stLoginIn := netsdk.NET_IN_LOGIN_WITH_HIGHLEVEL_SECURITY{}
// stLoginIn.ST_dwSize = (uint32)(unsafe.Sizeof(netsdk.NET_IN_LOGIN_WITH_HIGHLEVEL_SECURITY{}))
// copy(stLoginIn.ST_szIP[:], []byte(ip)[:])
// stLoginIn.ST_nPort = int32(port)
// copy(stLoginIn.ST_szUserName[:], []byte(user)[:])
// copy(stLoginIn.ST_szPassword[:], []byte(pswd)[:])
// stLoginIn.ST_emSpecCap = netsdk.EM_LOGIN_SPEC_CAP_SERVER_CONN
// snCS := C.CString(sn)
// stLoginIn.ST_pCapParam = uintptr(unsafe.Pointer(snCS))

// stLoginOut := netsdk.NET_OUT_LOGIN_WITH_HIGHLEVEL_SECURITY{}
// stLoginOut.ST_dwSize = (uint32)(unsafe.Sizeof(netsdk.NET_OUT_LOGIN_WITH_HIGHLEVEL_SECURITY{}))
// lhandle := netsdk.LoginWithHighLevelSecurity(&stLoginIn, &stLoginOut)
// C.free(unsafe.Pointer(snCS))
// if lhandle == 0 {
// 	fmt.Printf("LoginWithHighLevelSecurity failed, 0x%x\n", netsdk.GetLastError())
// 	return
// }

func LoginMode(mode EM_LOGIN_SPAC_CAP_TYPE) LogOptionFunc {
	return func(inparam *NET_IN_LOGIN_WITH_HIGHLEVEL_SECURITY) {
		inparam.ST_emSpecCap = mode
	}
}

func LoginActive(sn string) LogOptionFunc {
	return func(inparam *NET_IN_LOGIN_WITH_HIGHLEVEL_SECURITY) {
		inparam.ST_pCapParam = uintptr(unsafe.Pointer(C.CString(sn)))
	}
}

func (client *Client) Logout() bool {
	return Logout(client.LoginID)
}

func (client *Client) StartListen() bool {
	return StartListenEx(client.LoginID)
}

func (client *Client) StopListen() bool {
	return StopListen(client.LoginID)
}

func (client *Client) RealLoadPictureEx(channel int, evt EventIvs, callback PictureExFunc) error {
	var visior = PictureVisitor{
		Client:   client,
		Callback: callback,
	}
	// var userdata *LLONG = (*LLONG)(unsafe.Pointer(&visior))
	p := pointer.Save(&visior)

	lAnalyzerHandle := C.CLIENT_RealLoadPictureEx(
		C.long(client.LoginID),
		C.int(channel),
		C.uint(evt),
		C.int(1),
		C.fAnalyzerDataCallBack(C.export_fAnalyzerDataCallBack2),
		C.long(uintptr(p)),
		unsafe.Pointer(uintptr(0)))

	if lAnalyzerHandle != 0 {
		log.Println("CLIENT_RealLoadPictureEx success")
	} else {
		return errors.New("can't realloadPicture")
	}

	client.subscribP = p
	log.Println("lAnalyzerHandle=", lAnalyzerHandle)
	client.realloadHandle = LLONG(lAnalyzerHandle)

	return nil
}

func (client *Client) StopLoadPic() bool {
	if client.realloadHandle == 0 {
		return false
	}

	pointer.Unref(client.subscribP)
	return StopLoadPic(client.realloadHandle)
}

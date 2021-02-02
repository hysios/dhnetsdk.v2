package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/hysios/dhnetsdk.v2/netsdk"
	"github.com/yudai/pp"
)

// func AnalyzerDataCallBack()

var (
	addr string
	user string
	pass string
)

func init() {
	flag.StringVar(&addr, "addr", "", "大华设备地址")
	flag.StringVar(&user, "user", "", "大华设备登陆帐号")
	flag.StringVar(&pass, "pass", "", "大华设备登陆密码")
	flag.Set("addr", "192.168.1.108:37777")
	flag.Set("user", "admin")
	flag.Set("pass", "admin123")
}

func main() {
	flag.Parse()

	err := netsdk.InitEx(func(ip string, port int) {
		log.Printf("disconnect %s %d", ip, port)
	})
	if err != nil {
		log.Fatalf("init netsdk error %s\n", err)
	}
	fmt.Println("Init NetSDK success")
	defer netsdk.Cleanup()
	// var search netsdk.Search
	// searchAllClients(&search)
	netsdk.SetAutoReconnect(func(client *netsdk.Client, ip string, port int) {
		log.Printf("reconnect ip %s %d", ip, port)
	})
	// netsdk.SetConnectTime(10*time.Second, -1)
	client, err := netsdk.Login(addr, user, pass)
	if err != nil {
		log.Fatalf("login failed %s", err)
	}
	defer client.Logout()

	pp.Printf("client %s", client.DeviceInfo)

	err = client.RealLoadPictureEx(0, netsdk.EVENT_IVS_ALL,
		func(client *netsdk.Client, alarmType netsdk.EventIvs, alarmInfo interface{}, frame []byte, seq int) int {
			switch alarmType {
			case netsdk.EVENT_IVS_TRAFFIC_PARKING:
				if info, ok := alarmInfo.(*netsdk.DEV_EVENT_TRAFFIC_PARKING_INFO); ok {

					fmt.Printf("%s, %d\n", netsdk.Str(info.ST_stuObject.ST_szObjectType[:]), info.ST_stuObject.ST_nObjectID)
					fmt.Printf("%s, %s\n", netsdk.Str(info.ST_stuVehicle.ST_szObjectType[:]), netsdk.Str(info.ST_stuVehicle.ST_szObjectSubType[:]))
					fmt.Printf("PlateNumber %s\n", netsdk.Str(info.ST_stTrafficCar.ST_szPlateNumber[:]))
					fmt.Printf("frame size %d\n", len(frame))
					// fmt.Printf("info %#v\n", info)
					globalfilename := fmt.Sprintf("images/%04d%02d%02d-%02d%02d%02d%04d-%d_%s",
						info.ST_UTC.ST_dwYear,
						info.ST_UTC.ST_dwMonth,
						info.ST_UTC.ST_dwDay,
						info.ST_UTC.ST_dwHour,
						info.ST_UTC.ST_dwMinute,
						info.ST_UTC.ST_dwSecond,
						info.ST_UTC.ST_dwMillisecond,
						info.ST_stuObject.ST_nObjectID,
						"big.jpg")
					file, err := os.OpenFile(globalfilename, os.O_CREATE|os.O_WRONLY, 0666)
					if err != nil {
						fmt.Println("Error OpenFile ")
					}
					defer file.Close()

					n, err := file.Write(frame)
					_ = n
					jsonfile := strings.TrimSuffix(globalfilename, "jpg") + ".json"
					saveJson(jsonfile, info)
					fmt.Println("globalfilename-----", globalfilename)
				}

			case netsdk.EVENT_IVS_TRAFFICJUNCTION:
				if info, ok := alarmInfo.(*netsdk.DEV_EVENT_TRAFFICJUNCTION_INFO); ok {
					globalfilename := fmt.Sprintf("images/%04d%02d%02d-%02d%02d%02d%04d-%d_%s",
						info.ST_UTC.ST_dwYear,
						info.ST_UTC.ST_dwMonth,
						info.ST_UTC.ST_dwDay,
						info.ST_UTC.ST_dwHour,
						info.ST_UTC.ST_dwMinute,
						info.ST_UTC.ST_dwSecond,
						info.ST_UTC.ST_dwMillisecond,
						info.ST_stuObject.ST_nObjectID,
						"big.jpg")
					file, err := os.OpenFile(globalfilename, os.O_CREATE|os.O_WRONLY, 0666)
					if err != nil {
						fmt.Println("Error OpenFile ")
					}
					defer file.Close()

					n, err := file.Write(frame)
					_ = n
					jsonfile := strings.TrimSuffix(globalfilename, "jpg") + ".json"

					saveJson(jsonfile, info)
					fmt.Printf("%s\n", netsdk.Str(info.ST_szName[:]))
					fmt.Printf("ObjectType %s\n", netsdk.Str(info.ST_stuObject.ST_szObjectType[:]))
					fmt.Printf("PlateNumber %s\n", netsdk.Str(info.ST_stuObject.ST_szText[:]))
					fmt.Printf("frame size %d\n", len(frame))

					// pp.Printf("info %s", info)
				}

			case netsdk.EVENT_TRAFFICSNAPSHOT:
				if info, ok := alarmInfo.(*netsdk.DEV_EVENT_TRAFFICSNAPSHOT_INFO); ok {

					// fmt.Printf("%s, %d\n", netsdk.Str(info.ST_stuObject.ST_szObjectType[:]), info.ST_stuObject.ST_nObjectID)
					// fmt.Printf("%s, %s\n", netsdk.Str(info.ST_stuVehicle.ST_szObjectType[:]), netsdk.Str(info.ST_stuVehicle.ST_szObjectSubType[:]))
					// fmt.Printf("PlateNumber %s\n", netsdk.Str(info.ST_stTrafficCar.ST_szPlateNumber[:]))
					// fmt.Printf("frame size %d\n", len(frame))
					// // fmt.Printf("info %#v\n", info)

					// 	info.ST_UTC.ST_dwYear,
					// 	info.ST_UTC.ST_dwMonth,
					// 	info.ST_UTC.ST_dwDay,
					// 	info.ST_UTC.ST_dwHour,
					// 	info.ST_UTC.ST_dwMinute,
					// 	info.ST_UTC.ST_dwSecond,
					// 	info.ST_UTC.ST_dwMillisecond,
					// 	info.ST_stuObject.ST_nObjectID,
					// 	"big.jpg")
					// file, err := os.OpenFile(globalfilename, os.O_CREATE|os.O_WRONLY, 0666)
					// if err != nil {
					// 	fmt.Println("Error OpenFile ")
					// }
					// defer file.Close()

					// n, err := file.Write(frame)
					// _ = n

					pp.Printf("info %s", info)
					// fmt.Println("globalfilename-----", globalfilename)
				}
			}
			return 0
		},
	)
	log.Printf("error %s", err)

	for {
		time.Sleep(1 * time.Second)
	}
}

func saveJson(filename string, info interface{}) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 06666)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	return enc.Encode(info)
}

func searchAllClients(search *netsdk.Search) error {

	ifaces, err := net.Interfaces()
	if err != nil {
		return err
	}
	// handle err
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			return err
		}
		// handle err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
				log.Printf("ipv4: %s\n", ip)
			case *net.IPAddr:
				ip = v.IP
				log.Printf("ipv6: %s\n", ip)
			}

			search.LocalIP = ip.String()
			search.Start(func(search *netsdk.Search, deviceinfo *netsdk.DEVICE_NET_INFO_EX) {
				// log.Printf("deviceinfo\n%#v\n", deviceinfo)
				log.Printf("SerialNo: %s\n", netsdk.Str(deviceinfo.ST_szSerialNo[:]))
				log.Printf("Version: %s\n", netsdk.Str(deviceinfo.ST_szDevSoftVersion[:]))
				log.Printf("Ip: %s\n", netsdk.Str(deviceinfo.ST_szIp[:]))
				log.Printf("Port: %d\n", deviceinfo.ST_nPort)
				log.Printf("HTTP Port: %d\n", deviceinfo.ST_nHttpPort)
				log.Printf("Device Type: %s\n", netsdk.Str(deviceinfo.ST_szDeviceType[:]))
			}, netsdk.EM_SEND_SEARCH_TYPE_MULTICAST_AND_BROADCAST)
			// process IP address

			time.Sleep(5 * time.Second)
			search.Stop()
		}
	}

	return err
}

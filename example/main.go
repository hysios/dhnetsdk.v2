package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/hysios/dhnetsdk.v2/netsdk"
	"github.com/yudai/pp"
)

// func AnalyzerDataCallBack()

func main() {
	err := netsdk.Init(nil)
	if err != nil {
		log.Fatalf("init netsdk error %s\n", err)
	}
	fmt.Println("Init NetSDK success")
	defer netsdk.Cleanup()
	var search netsdk.Search
	searchAllClients(&search)

	client, err := netsdk.Login("192.168.1.190:37777", "admin", "admin123")
	if err != nil {
		log.Fatalf("login failed %s", err)
	}

	pp.Printf("client %s", client.DeviceInfo)

	client.RealLoadPictureEx(0, netsdk.EVENT_IVS_ALL,
		func(client *netsdk.Client, alarmType netsdk.EventIvs, alarmInfo interface{}, frame []byte, seq int) int {
			switch alarmType {
			case netsdk.EVENT_IVS_TRAFFIC_PARKING:
				if info, ok := alarmInfo.(*netsdk.DEV_EVENT_TRAFFIC_PARKING_INFO); ok {

					fmt.Printf("%s, %d\n", netsdk.Str(info.ST_stuObject.ST_szObjectType[:]), info.ST_stuObject.ST_nObjectID)
				}
			}
			return 0
		},
	)

	for {
		time.Sleep(1 * time.Second)
	}
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

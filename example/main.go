package main

import (
	"fmt"
	"log"
	"time"

	"github.com/hysios/dhnetsdk/netsdk"
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

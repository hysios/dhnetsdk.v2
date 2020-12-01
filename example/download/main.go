package main

import (
	"flag"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/araddon/dateparse"
	"github.com/hysios/dhnetsdk.v2/netsdk"
)

var (
	addr      string
	user      string
	pass      string
	startTime string
	duration  time.Duration
	output    string
)

var loc *time.Location

func init() {
	flag.StringVar(&addr, "addr", "", "大华设备地址")
	flag.StringVar(&user, "user", "", "大华设备登陆帐号")
	flag.StringVar(&pass, "pass", "", "大华设备登陆密码")
	flag.StringVar(&startTime, "start-time", "", "开始时间")
	flag.DurationVar(&duration, "duration", 35*time.Second, "持续时间")
	flag.StringVar(&output, "output", "test.dav", "保存文件")

	flag.Set("addr", "192.168.1.190:37777")
	flag.Set("user", "admin")
	flag.Set("pass", "admin123")

	_loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Fatalf("load location error %s", err)
	}
	loc = _loc
}

func main() {
	flag.Parse()
	var stTime time.Time = time.Now().Add(-40 * time.Second)

	err := netsdk.InitEx(nil)
	if err != nil {
		log.Fatalf("init netsdk error %s\n", err)
	}
	fmt.Println("Init NetSDK success")
	defer netsdk.Cleanup()

	client, err := netsdk.Login(addr, user, pass)
	if err != nil {
		log.Fatalf("login failed %s", err)
	}

	if len(startTime) > 0 {
		stTime, err = dateparse.ParseAny(startTime)
		if err != nil {
			log.Fatalf("parse date error %s", err)
		}
	}
	log.Printf("start time %s", stTime)
	if err != nil {
		log.Fatalf("parse any date error %s", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	play, err := client.DownloadByTimeEx(0, netsdk.EM_RECORD_TYPE_ALL, stTime, duration, output, nil)
	if err != nil {
		log.Fatalf("download by time error %s", err)
	}
	log.Printf("playback %#v", play)
	<-play.End()
	play.Stop()
}

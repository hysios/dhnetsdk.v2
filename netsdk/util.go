package netsdk

import (
	"bytes"
	"log"
	"sync"
	"time"
)

type ObjectId int32

var refs struct {
	sync.Mutex
	objs map[ObjectId]interface{}
	next ObjectId
}

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	refs.Lock()
	defer refs.Unlock()

	refs.objs = make(map[ObjectId]interface{})
	refs.next = 1000
}

func NewObjectId(obj interface{}) ObjectId {
	refs.Lock()
	defer refs.Unlock()

	id := refs.next
	refs.next++

	refs.objs[id] = obj
	return id
}

func (id ObjectId) IsNil() bool {
	return id == 0
}

func (id ObjectId) Get() interface{} {
	refs.Lock()
	defer refs.Unlock()

	return refs.objs[id]
}

func (id *ObjectId) Free() interface{} {
	refs.Lock()
	defer refs.Unlock()

	obj := refs.objs[*id]
	delete(refs.objs, *id)
	*id = 0

	return obj
}

func Str(b []byte) string {
	n := bytes.IndexByte(b, 0x00)
	if n >= 0 {
		return string(b[:n])
	}
	return string(b)
}

func NT2time(nt NET_TIME) time.Time {
	return time.Date(
		int(nt.ST_dwYear),
		time.Month(nt.ST_dwMonth),
		int(nt.ST_dwDay),
		int(nt.ST_dwHour),
		int(nt.ST_dwMinute),
		int(nt.ST_dwSecond),
		0,
		time.UTC)
}

func NTex2time(nt NET_TIME_EX) time.Time {
	return time.Date(
		int(nt.ST_dwYear),
		time.Month(nt.ST_dwMonth),
		int(nt.ST_dwDay),
		int(nt.ST_dwHour),
		int(nt.ST_dwMinute),
		int(nt.ST_dwSecond),
		int(nt.ST_dwMillisecond*1e6),
		time.Local)
}

func Time2nt(t time.Time) NET_TIME {
	return NET_TIME{
		ST_dwYear:   uint32(t.Year()),   // 年
		ST_dwMonth:  uint32(t.Month()),  // 月
		ST_dwDay:    uint32(t.Day()),    // 日
		ST_dwHour:   uint32(t.Hour()),   // 时
		ST_dwMinute: uint32(t.Minute()), // 分
		ST_dwSecond: uint32(t.Second()), // 秒
	}
}

func Time2ntex(t time.Time) NET_TIME_EX {
	nano := t.UnixNano() % 1e6

	return NET_TIME_EX{
		ST_dwYear:        uint32(t.Year()),   // 年
		ST_dwMonth:       uint32(t.Month()),  // 月
		ST_dwDay:         uint32(t.Day()),    // 日
		ST_dwHour:        uint32(t.Hour()),   // 时
		ST_dwMinute:      uint32(t.Minute()), // 分
		ST_dwSecond:      uint32(t.Second()), // 秒
		ST_dwMillisecond: uint32(nano),
		ST_dwUTC:         0,
	}
}

func NTdate2time(nt NET_TIME_EX) time.Time {
	t := time.Now()
	return time.Date(
		t.Year(),
		time.Month(nt.ST_dwYear),
		int(nt.ST_dwMonth),
		int(nt.ST_dwDay),
		int(nt.ST_dwHour),
		int(nt.ST_dwMinute),
		int(nt.ST_dwSecond),
		time.Local,
	)
}

func Time2ntdate(t time.Time) NET_TIME_EX {
	return NET_TIME_EX{
		ST_dwYear:   uint32(t.Month()),  // 年
		ST_dwMonth:  uint32(t.Day()),    // 月
		ST_dwDay:    uint32(t.Hour()),   // 日
		ST_dwHour:   uint32(t.Minute()), // 时
		ST_dwMinute: uint32(t.Second()), // 分
		ST_dwSecond: uint32(t.Second()), // 秒
		ST_dwUTC:    0,
	}
}

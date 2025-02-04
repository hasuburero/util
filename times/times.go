package times

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	Layout1 = "2006-01-02 15:04:05.000000"
)

type TimeStamp struct {
	Timer     time.Time
	Timestamp string // formatted with time.StampMilli
	Hour      int    // hour part of timestamp
	Min       int    // min part of timestamp
	Sec       int    // sec part of timestamp
	Ms        int    // ms part of timestamp
	Millis    int    // total milli seconds
}

func (self *TimeStamp) Getms() {
	self.Millis = self.Ms + 1000*self.Sec + 60000*self.Min + 3600000*self.Hour
}

func (self *TimeStamp) GetTime() {
	self.Timer = time.Now()
	self.Timestamp = self.Timer.Format(time.StampMilli)
	buf := strings.Split(self.Timestamp, " ")
	for i, ctx := range buf {
		if ctx == "" {
			buf = append(buf[:i], buf[i+1:]...)
		}
	}
	tbuf := buf[2]
	buf = strings.Split(tbuf, ":")
	self.Hour, _ = strconv.Atoi(string(buf[0]))
	self.Min, _ = strconv.Atoi(string(buf[1]))
	tbuf = buf[2]
	buf = strings.Split(tbuf, ".")
	self.Sec, _ = strconv.Atoi(string(buf[0]))
	self.Ms, _ = strconv.Atoi(string(buf[1]))
}

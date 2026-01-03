package zlog

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

type fileWriter struct {
	logRoot             string
	f                   *os.File
	w                   *bufio.Writer
	mutex               sync.Mutex
	lastUpdateHourInDay int64
}

func (self *fileWriter) Flush() error {
	var err error
	self.mutex.Lock()
	if self.w != nil {
		err = self.w.Flush()
	}
	self.mutex.Unlock()
	return err
}

func (self *fileWriter) Write(d []byte) (int, error) {
	self.mutex.Lock()
	t := time.Now()
	err := self.prepare(&t)
	var n int
	if err == nil {
		n, err = self.w.Write(d)
	}
	self.mutex.Unlock()
	return n, err
}

func (self *fileWriter) prepare(now *time.Time) error {
	hoursInDay := now.Unix() / 3600
	// $time_$group_$id.log跟stat$time_$id.log
	if self.f == nil || hoursInDay != self.lastUpdateHourInDay { // 15:04:05
		fileName := fmt.Sprintf("%s/%s.log", self.logRoot, now.Format("2006010215"))
		f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600|0644)
		if err != nil {
			return err
		}
		if self.f != nil {
			_ = self.w.Flush()
			self.f.Close() //nolint
		}
		self.f = f
		self.w = bufio.NewWriter(f)
		self.lastUpdateHourInDay = hoursInDay
	}

	// if len(prefix) > 0 {
	// 	_, err := self.w.WriteString(prefix)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	self.formatHeader(now)
	// }
	return nil
}

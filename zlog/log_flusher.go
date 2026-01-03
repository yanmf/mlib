package zlog

import (
	"sync"
	"time"
)

var flusherQueue = &struct {
	sync.RWMutex
	flushers []*fileWriter
}{}

func addFlusher(lf *fileWriter) {
	flusherQueue.Lock()
	flusherQueue.flushers = append(flusherQueue.flushers, lf)
	flusherQueue.Unlock()
}

func flushAllLog() {
	flusherQueue.RLock()
	for _, v := range flusherQueue.flushers {
		_ = v.Flush()
	}
	flusherQueue.RUnlock()
}

func startFlusher() {
	go func() {
		for {
			time.Sleep(time.Second)
			flushAllLog()
		}
	}()
}

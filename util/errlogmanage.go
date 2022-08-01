package util

import (
	"strconv"
	"sync"
)

var logIndexErrOnce sync.Once
var logIndexErrInstance *LogIndexErr

type LogIndexErr struct {
	Index string
}

func InitLogIndex() *LogIndexErr {
	logIndexErrOnce.Do(func() {
		logIndexErrInstance = &LogIndexErr{}
	})
	return logIndexErrInstance
}

func (l *LogIndexErr) NewLogIndex(Index string) {
	temp, _ := strconv.Atoi(Index)
	l.Index = strconv.Itoa(temp + 1)
}

func (l *LogIndexErr) NowLogIndex() string {
	return l.Index
}

func (l *LogIndexErr) AddLogIndex() {
	temp, _ := strconv.Atoi(l.Index)
	l.Index = strconv.Itoa(temp + 1)
}

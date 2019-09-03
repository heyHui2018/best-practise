package base

import (
	"github.com/heyHui2018/log"
)

func LogInit() {
	// log.SetOutputByName(GetConfig().Log.Path)
	log.SetLevelByString(GetConfig().Log.Level)
	log.SetRotateByDay()
	log.SetCallDepth(5)
}

type LogP struct {
	TraceId string
}

func (l *LogP) Infof(format string, v ...interface{}) {
	log.Infof(l.TraceId+" "+format, v)
}

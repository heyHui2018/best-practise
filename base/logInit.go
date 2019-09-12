package base

import (
	"github.com/heyHui2018/log"
)

func LogInit() {
	// log.SetOutputByName(GetConfig().Log.Path)
	log.SetLevelByString(GetConfig().Log.Level)
	log.SetRotateByDay()
	// log.SetCallDepth(5)
}

type TLog struct {
	TraceId string
}

func (t *TLog) Infof(format string, v ...interface{}) {
	log.Infof(t.TraceId+" "+format, v)
}

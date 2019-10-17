package base

import (
	"github.com/heyHui2018/log"
)

func LogInit() {
	// log.SetOutputByName(GetConfig().Log.Path)
	log.SetLevelByString(GetConfig().Log.Level)
	log.SetRotateByDay()
}
package base

import (
	"github.com/ngaut/log"
)

func LogInit() {
	// log.SetOutputByName(GetConfig().Log.Path)
	log.SetLevelByString(GetConfig().Log.Level)
	log.SetRotateByDay()
}

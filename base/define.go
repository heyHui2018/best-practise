package base

const (
	Success      = 200
	SystemError  = 500
	MissingParam = 1001
	ParamError   = 1002

	AirVisualDataLock = "AirVisualDataLock"
)

var CodeText = map[int]string{
	Success:      "Success",
	SystemError:  "System error",
	MissingParam: "Missing param",
	ParamError:   "Param error",
}

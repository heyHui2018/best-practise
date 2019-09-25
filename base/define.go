package base

const (
	Success      = 200
	SystemError  = 500
	MissingParam = 1001
	ParamError   = 1002

	AirVisualDataLock = "AirVisualDataLock"

	SockAddr           = "/var/run//docker.sock"                          // docker官网文档定义的套接字默认值
	ImagesSock         = "GET /images/json HTTP/1.0\r\n\r\n"              // docker对外的镜像操作api
	ContainerSock      = "GET /containers/json?all=true HTTP/1.0\r\n\r\n" // docker对外的容器查询api
	StartContainerSock = "POST /containers/%s/start HTTP/1.0\r\n\r\n"     // docker对外的容器启动api
)

var CodeText = map[int]string{
	Success:      "Success",
	SystemError:  "System error",
	MissingParam: "Missing param",
	ParamError:   "Param error",
}

package docker

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/heyHui2018/best-practise/base"
	"github.com/heyHui2018/best-practise/models/docker"
	"github.com/heyHui2018/log"
	"io/ioutil"
	"net"
	"strings"
	"time"
)

func Start() {
	tLog := new(log.TLog)
	interval := base.GetConfig().Docker.Interval

	// 循环监控容器是否正常,有问题重启
	for {
		// 轮询docker
		err := listenDocker(tLog, base.GetConfig().Docker.WhiteList)
		if err != nil {
			log.Warnf(err.Error())
		}
		time.Sleep(time.Duration(interval) * time.Second)
	}
}

func listenDocker(log *log.TLog, whiteList []string) error {
	// 获取容器信息列表
	containers, err := readContainer()
	if err != nil {
		return err
	}
	for k := range whiteList {
		for _, container := range containers {
			for _, cname := range container.Names {
				if cname[1:] == whiteList[k] {
					log.Infof("id=%s, name=%s, state=%s", container.Id[:12], container.Names, container.Status)
					if strings.Contains(container.Status, "Exited") {
						err = startContainer(container.Id, log)
						if err != nil {
							log.Warnf("listenDocker,start container error,err = %v,container = %+v", err, container)
						} else {
							log.Infof("listenDocker,start container successfully,container = %+v", container)
						}
					}
				}
			}
		}
	}
	return nil
}

// 获取 unix sock 连接
func connectDocker() (*net.UnixConn, error) {
	addr := net.UnixAddr{base.SockAddr, "unix"} // SockAddr被设定为docker的/var/run/docker套接字路径值,建立与docker的daemon的连接
	return net.DialUnix("unix", nil, &addr)
}

// 启动容器
func startContainer(id string, log *log.TLog) error {
	conn, err := connectDocker()
	if err != nil {
		return err
	}
	code, err := conn.Write([]byte(fmt.Sprintf(base.StartContainerSock, id)))
	if err != nil {
		return err
	}
	log.Infof("startContainer 完成,code = %v", code)
	return nil
}

// 获取容器列表
func readContainer() ([]docker.Container, error) {
	conn, err := connectDocker() // 建立一个unix连接
	if err != nil {
		return nil, err
	}
	_, err = conn.Write([]byte(base.ContainerSock)) // 建立的tcp连接不能复用，每次操作都建立连接
	if err != nil {
		return nil, err
	}
	result, err := ioutil.ReadAll(conn)
	if err != nil {
		return nil, err
	}
	body := getBody(result)
	var containers []docker.Container
	err = json.Unmarshal(body, &containers)
	if err != nil {
		return nil, err
	}
	if len(containers) == 0 {
		return nil, errors.New("len(containers) == 0")
	}
	return containers, nil
}

// 获取镜像列表
func readImage(conn *net.UnixConn) ([]docker.Image, error) {
	_, err := conn.Write([]byte(base.ImagesSock))
	if err != nil {
		return nil, err
	}

	result, err := ioutil.ReadAll(conn)
	if err != nil {
		return nil, err
	}

	body := getBody(result)
	var images []docker.Image
	err = json.Unmarshal(body, &images)
	if err != nil {
		return nil, err
	}
	return images, nil
}

// 从http响应中提取body
func getBody(result []byte) []byte {
	for i := 0; i <= len(result)-4; i++ {
		if result[i] == 13 && result[i+1] == 10 && result[i+2] == 13 && result[i+3] == 10 {
			return result[i+4:]
		}
	}
}

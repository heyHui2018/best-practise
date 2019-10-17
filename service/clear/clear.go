package clear

import (
	"context"
	"flag"
	"github.com/heyHui2018/best-practise/service/nsq"
	"github.com/heyHui2018/best-practise/service/rabbitMQ"
	"github.com/heyHui2018/log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"time"
)

/*
http服务优雅重启详见：https://github.com/heyHui2018/graceful
*/

var Graceful = flag.Bool("graceful", false, "listen on fd open 3 (internal use only)")

func Clear() {
	// 关闭mq
	rabbitMQ.MQCloseSign = true
	rabbitMQ.ConsumeChannel.Close()
	rabbitMQ.ConsumeConn.Close()
	rabbitMQ.ConsumeWait.Wait()
	log.Info("MQ已正常关闭")
	// 关闭nsq
	nsq.NsqCloseSign = true
	nsq.NsqStopChan <- ""
	if nsq.Producer != nil {
		nsq.Producer.Stop()
	}
	nsq.NsqWait.Wait()
	log.Info("NSQ已正常关闭")
	// ...
}

func Stop(server *http.Server) {
	log.Info("开始停止程序...")
	// api优雅退出
	cxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := server.Shutdown(cxt)
	if err != nil {
		log.Warnf("clear,server.Shutdown error,err = %v", err)
	}
	log.Info("资源清理成功，退出")
	os.Exit(0)
}

func Restart(server *http.Server, listener net.Listener) {
	log.Info("开始重启...")
	l, ok := listener.(*net.TCPListener)
	if !ok {
		log.Warnf("reload,listener's type is not *net.TCPListener,type = %v", reflect.TypeOf(listener).Name())
		return
	}
	file, err := l.File()
	if err != nil {
		log.Warnf("reload,l.File error,err = %v", err)
		return
	}
	args := os.Args
	if !*Graceful {
		args = append(args, "-graceful")
	}
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.ExtraFiles = []*os.File{file}
	err = cmd.Start()
	if err == nil {
		Stop(server)
	}
}

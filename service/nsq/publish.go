package nsq

import (
	"github.com/heyHui2018/log"
)

func Publish(t log.TLog, topic, msg string) {
	if msg == "" {
		t.Warnf("Publish,msg为空")
		return
	}
	if Producer == nil {
		PublishStart()
	}
	err := Producer.Publish(topic, []byte(msg))
	if err != nil {
		t.Warnf("Publish error,err = %v", err)
		Producer = nil
		return
	}
	t.Infof("Publish successfully")
}
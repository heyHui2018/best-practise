package nsq

import (
	"github.com/heyHui2018/log"
	"github.com/heyHui2018/utils"
	"github.com/nsqio/go-nsq"
	"time"
)

type Consumer struct {
	c *nsq.Consumer
}

func (this *Consumer) HandleMessage(msg *nsq.Message) error {
	t := new(log.TLog)
	t.TraceId = time.Now().Format("20060102150405") + utils.GetRandomString()
	t.Infof("HandleMessage,msg = %v", string(msg.Body))
	// do something with msg.Body
	return nil
}

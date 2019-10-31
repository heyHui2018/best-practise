package nsq

import (
	"time"

	"github.com/heyHui2018/log"
	"github.com/heyHui2018/utils"
	"github.com/nsqio/go-nsq"
)

type ConsumerS struct {
	c *nsq.Consumer
}

func (this *ConsumerS) HandleMessage(msg *nsq.Message) error {
	NsqWait.Add(1)
	defer NsqWait.Done()
	t := new(log.TLog)
	t.TraceId = time.Now().Format("20060102150405") + utils.GetRandomString()
	t.Infof("HandleMessage,msg = %v", string(msg.Body))
	// do something with msg.Body
	return nil
}

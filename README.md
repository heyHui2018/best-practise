### 最佳实践

***
有任何意见或建议可加qq：962691478，欢迎交流

***
包含如下框架/组件:
* 1.http框架
    * [x] gin github.com/gin-gonic/gin
* 2.rpc框架
    * [x] grpc google.golang.org/grpc
* 3.配置读取
    * [x] toml github.com/BurntSushi/toml
* 4.数据存储
    * [x] mysql     github.com/go-sql-driver/mysql github.com/go-xorm/xorm
    * [x] redis     github.com/garyburd/redigo/redis
    * [x] influxDB  github.com/influxdata/influxdb/client/v2
* 5.消息中间件
    * [x] rabbitMQ  github.com/streadway/amqp
    * [x] kafka     github.com/Shopify/sarama github.com/bsm/sarama-cluster
    * [x] nsq       github.com/nsqio/go-nsq
* 6.定时任务
    * [x] cron github.com/robfig/cron
* 7.服务发现
    * [x] etcd go.etcd.io/etcd/clientv3
* 8.分布式锁
    * [x] redis github.com/garyburd/redigo/redis
    * [ ] zookeeper
* 9.提醒
    * [x] 邮件 net/smtp
    * [ ] 短信
* 10.数据源
    * [x] weather API   api.airvisual.com
    * [ ] leetCode      github.com/heyHui2018/leetCode
    * [ ] excel
* 11.others
    * [x] JWT       github.com/appleboy/gin-jwt
    * [x] qrCode    github.com/skip2/go-qrcode
    * [x] restart   gracefully github.com/heyHui2018/graceful
    * [x] docker容器监控/重启

***
* [x] 功能：延迟消息 rabbitMQ死信队列
* [x] 功能：接口注册登记,定时发送邮件
* [x] 功能：docker容器监控/重启
* [ ] 功能：根据传入的题号或关键词,查询题目相关数据

***
### TodoList
* 优化grpc router
* rpc负载均衡(各节点部署一个进程监控etcd上的服务ip及端口,随后将这些服务写入iptables上做snat转发,通过iptables的snat修改数据包目的ip实现负载均衡)
* 微信提醒
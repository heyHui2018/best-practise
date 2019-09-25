### 最佳实践
***
有任何意见或建议可加qq：962691478，欢迎交流
***
包含如下框架/组件:
* [x] gin--框架--github.com/gin-gonic/gin
* [x] toml--配置--github.com/BurntSushi/toml
* [x] mysql--数据库--github.com/go-sql-driver/mysql--github.com/go-xorm/xorm
* [x] redis--缓存/分布式锁--github.com/garyburd/redigo/redis
* [ ] zookeeper--分布式锁
* [x] rabbitMQ--消息中间件--github.com/streadway/amqp
    * [x] 死信队列--延迟消息--github.com/streadway/amqp
* [x] 天气API--数据源--api.airvisual.com
* [x] cron--定时任务--github.com/robfig/cron
* [x] 优雅退出
* [x] JWT--github.com/appleboy/gin-jwt
* [x] grpc--微服务调用方式--google.golang.org/grpc
* [x] etcd--服务发现--"go.etcd.io/etcd/clientv3"
* [ ] 短信--提醒
* [x] 邮件--提醒--net/smtp
* [ ] 二维码
* [ ] nsq
* [x] kafka--日志收集--"github.com/Shopify/sarama"--"github.com/bsm/sarama-cluster"
* [x] influxDB--数据收集--github.com/influxdata/influxdb/client/v2
* [ ] grafana--数据展示
* [x] leetCode--数据源--github.com/heyHui2018/leetCode
* [ ] 功能：根据传入的题号或关键词,查询题目相关数据
* [x] 功能：接口注册登记,定时发送邮件
* [x] 功能：docker容器监控/重启
***
### TodoList
* 优化grpc router
* 增加限流逻辑ratelimit
* rpc负载均衡(各节点部署一个进程监控etcd上的服务ip及端口,随后将这些服务写入iptables上做snat转发,通过iptables的snat修改数据包目的ip实现负载均衡)

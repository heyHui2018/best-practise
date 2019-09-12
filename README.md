### 最佳实践
***
包含如下框架/组件:
* [x] gin--框架--github.com/gin-gonic/gin
* [x] toml--配置--github.com/BurntSushi/toml
* [x] mysql--数据库--github.com/go-sql-driver/mysql--github.com/go-xorm/xorm
* [x] redis--缓存/分布式锁--github.com/garyburd/redigo/redis
* [x] rabbitMQ--消息中间件--github.com/streadway/amqp
    * [x] 死信队列
* [x] 天气API--数据源--api.airvisual.com
* [x] cron--定时任务--github.com/robfig/cron
* [x] 优雅退出
* [x] JWT
* [x] grpc--微服务调用方式--google.golang.org/grpc
* [x] etcd--服务发现
* [ ] 短信--提醒
* [x] 邮件--提醒--net/smtp
* [ ] 二维码
* [ ] nsq
* [x] kafka--日志收集
* [x] influxDB--数据收集--github.com/influxdata/influxdb/client/v2
* [ ] grafana--数据展示
* [x] leetCode--数据源--github.com/heyHui2018/leetCode
* [ ] 功能：根据传入的题号或关键词,查询题目相关数据
* [x] 功能：接口注册登记,定时发送邮件
***
### TodoList
* 优化grpc router

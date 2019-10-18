### NSQ

#### 1.nsqlookupd
通过tcp(端口4160)管理nsqd服务,通过tcp(端口4161)管理nsqadmin服务,为客户端提供nsqd地址查询服务
一套nsq服务仅一个nsqlookupd服务,集群中可多台但之间无联系.nsqlookupd崩溃不会影响nsqd服务

#### 2.nsqadmin
一个队topic和channel统一管理的操作界面及监控数据展示.默认访问地址：http://127.0.0.1:4171/

#### 3.nsqd
负责收发message及维护队列,默认监听tcp端口4150及http端口4151及另一可选https端口
* 同一topic同一channel的消费者负载均衡(非轮询)
* 当channel存在时,即使channel无消费者,生产者的msg也会被缓存到队列
* 队列中的msg至少会被消费一次
* 限定内存占用,当channel的在内存中的msg数量超出时,msg会被缓存到磁盘
* topic及channel在创建后会被保存,需定时清理无效topic及channel

#### 4.连接方式
* 消费者直连nsqd,此法使得nsqd服务无法动态扩展.当检测到连接断开后,每隔x秒会自动重连
* 消费者通过http查询nsqlookupd获取nsqd地址,再连接nsqd,此法客户端会轮询nsqlookupd
* 生产者必须直连nsqd且断开后需手动重连
* 消费者可以同时接收不同nsqd节点的同名topic数据

#### 

### RabbitMQ

#### 1.channel
频繁建立、销毁TCP连接开销太大

#### 2.消息未确认
RabbitMQ在接收到消费者的ack后会把消息从队列上删除.

若消费者在发送ack之前和RabbitMQ断开了连接(或者从队列上取消了订阅),RabbitMQ会认为该消息未被分发,进而会分发给下一个订阅者.
RabbitMQ没有用到超时机制,仅通过Consumer的连接中断来确认该Message没有被正确处理,即RabbitMQ给了Consumer足够长的时间来做数据处理,
只要连接存在,RabbitMQ会一直等待Consumer的确认消息.

#### 3.消息丢失
* 生产者发给RabbitMQ但MQ未收到.开启confirm模式,参考第6条.(此方法的性能比事务机制更高,因事务机制是同步的,而发送确认是异步)
* RabbitMQ未持久化就挂了.持久化,参考第4条.
* RabbitMQ转发给了消费者但消费者挂了.开启手动ack,MQ会一直等待消费者的ack,参考第2条.

#### 4.持久化
durable属性决定了RabbitMQ是否需要在崩溃或者重启之后重新创建队列(或者交换器),有三个要点：
* 把投递模式选项设置为2(持久)
* 交换器持久化
* 队列持久化

特点：
持久性消息从服务器重启中恢复的方式是写入磁盘的持久化日志文件中,当发布一条持久性消息到持久化的交换器上时,消息提交到日志文件后才会响应;
持久性消息如果路由到了非持久化的队列当中,会自动从持久性日志中移除,即无法从服务器重启中恢复;
一旦被正确消费(经过确认后),RabbitMQ会在持久化日志中将这条消息标记为等待垃圾收集.在消费之前,如果重启,服务器会重建交换器和队列以及绑定,
重播持久性日志文件中的消息到合适的队列或者交换器上,这取决于宕机时消息处在哪个环节上

#### 5.死信
* 死信
    * 消息被拒绝(reject或nack)且requeue=false
    * 消息TTL过期
    * 队列达到最大长度
* 死信交换机

可与任何一个普通队列绑定,然后在业务队列出现死信时将消息发到死信队列

* 死信队列：

与死信交换机绑定用于存放死信的队列

* 使用

定义普通队列时指定参数：
x-dead-letter-exchange: 用来设置死信后发送的交换机
x-dead-letter-routing-key：用来设置死信的routingKey

#### 6.confirm模式
* 原理

所有在该信道上面发布的消息都会被指派一个唯一的ID(从1开始),一旦消息被投递到所有匹配的队列之后,broker会发送一个确认给生产者(包含消息的唯一ID),
这就使得生产者知道消息已经正确到达目的队列了,如果消息和队列是可持久化的,那么确认消息会在消息写入磁盘之后发出.confirm模式最大的好处在于异步,
一旦发布一条消息,生产者应用程序就可以在等待信道返回确认的同时继续发送下一条消息,如果RabbitMQ因为自身内部错误导致消息丢失,就会发送一条nack消息,
生产者应用程序可以在回调方法中处理该nack消息.
在channel 被设置成 confirm 模式之后,所有被 publish 的后续消息都将被ack或nack一次,但是无法对快慢做任何保证,且同一条消息不会既被ack又被nack.

* 开启方式

生产者通过调用channel的confirmSelect方法将channel设置为confirm模式,如果没有设置no-wait标志的话,broker会返回confirm.
select-ok表示同意发送者将当前channel信道设置为confirm模式(从目前RabbitMQ最新版本3.6来看,如果调用了channel.confirmSelect方法,
默认情况下是直接将no-wait设置成false的,也就是默认情况下broker是必须回传confirm.select-ok的).

#### 7.Qos
在消费RabbitMQ时,若使用"github.com/streadway/amqp"的Consume函数,要注意,此函数会自动新建channel,将获取的消息存入此channel,即内存中,
不受调用者的channel缓冲区的限制.若不想 Consume函数无限制地消费,可在Consume之前调用Qos函数：Channel.Qos(1, 0, true),
并将 Consume 函数的自动确认参数设为false,在range Consume获得的数据的for循环中,调用ACK函数手动确认,并将参数设为false(当参数为true时,为批量确认,
会将此之前消费的消息都确认掉).
设置Qos有利于消息的公平分发,不会产生某一消费者很忙其余很闲的情况,具体逻辑为：RabbitMQ会轮询(默认情况)向消费者分发消息,当设置了Qos后,
若此消费者未返回确认消息,MQ就不会向其发送新的消息,而会发送给下一个消费者.

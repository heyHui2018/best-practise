### ETCD

#### 1.安装
* 普通安装
    * 源码下载解压(https://github.com/etcd-io/etcd/releases)
    * 将可执行文件复制进$GOPATH/bin(mv etcd* /$GOPATH/bin)
    * 配置环境变量并生效(设置使用api3的命令 echo 'export ETCDCTL_API=3' >> /etc/profile)(source /etc/profile)
    * 启动(nohup ./etcd &)(若想通过机器IP访问则需要在启动时指定IP nohup ./etcd --listen-client-urls http://0.0.0.0:2379 
    --advertise-client-urls http://0.0.0.0:2379 --listen-peer-urls http://0.0.0.0:2380 > /tmp/etcd.log 2>&1 &)
* docker安装
    * 查找镜像(docker search etcd)
    * 拉取镜像(此处选择v3镜像 docker pull xieyanze/etcd3)
    * 运行节点0(docker run -d -p 2380:2380 -p 2379:2379 --name etcd xieyanze/etcd3 -name etcd 
    -advertise-client-urls http://172.16.16.114:2379 -listen-client-urls http://0.0.0.0:2379 -initial-advertise-peer-urls 
    http://172.16.16.114:2380 -listen-peer-urls http://0.0.0.0:2380  -initial-cluster-token etcd-cluster-1 -initial-cluster 
    "etcd=http://172.16.16.114:2380" -initial-cluster-state new 当启动报错时,重启docker即可 /etc/init.d/docker restart)
    * 验证(curl -L http://172.16.16.114:2479/v2/members 或 etcdctl --endpoints="http://172.16.16.114:2379" get foo)
    
#### 2.API
* 键值相关
    * Put 存
    * Get 取
    * Delete 删
    * Compact 压缩
    * Do 执行(Put/Get/Delete也基于Do)
    * Txn 事务,仅支持If/Then/Else/Commit
* Watch
    * Watch 监听key的变化
    * Close 关闭
* 租约相关
    * Grant 分配
    * Revoke 释放
    * TimeToLive 获取剩余TTL时间
    * Leases 获取所有租约
    * KeepAlive 续约
    * KeepAliveOnce 仅续约一次
    * Close 关闭
* 集群相关
    * MemberList 获取集群所有成员
    * MemberAdd 添加成员
    * MemberRemove 移除成员
    * MemberUpdate 更新成员
* 锁
    * Lock 获取锁
    * Unlock 释放锁
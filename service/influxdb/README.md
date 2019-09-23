### InfluxDB

#### 1.docker安装
* 查找镜像(docker search influxdb)
* 拉取镜像(docker pull influxdb)
* 运行(docker run -idt --name influxdb -p 8086:8086 influxdb)
* 创建database(1.进入容器 docker exec -it 容器ID /bin/bash 2.进入/usr/bin/ 3.启动influxdb的客户端 ./influx 4.create database test)
    
#### 2.关键词
```
name: census
time                     butterflies     honeybees     location   scientist
2015-08-18T00:00:00Z      12                23           1         langstroth
2015-08-18T00:00:00Z      1                 30           1         perpetua
2015-08-18T00:06:00Z      11                28           1         langstroth
2015-08-18T00:06:00Z      3                 28           1         perpetua
2015-08-18T05:54:00Z      2                 11           2         langstroth
2015-08-18T06:00:00Z      1                 10           2         langstroth
2015-08-18T06:06:00Z      8                 23           2         perpetua
2015-08-18T06:12:00Z      7                 22           2         perpetua
```
* timestamp 顾名思义
* field key 表中的butterflies和honeybees,必须存在,无索引,作为查询条件时,会扫描所有值,性能不及tag
* field value 表中的butterflies和honeybees下的值,类型可为string/float/integer/boolean
* tag key 表中的location和scientist,非必须存在,但建议使用,有索引
* tag value 表中的location和scientist下的值,类型仅能为string
* measurement 类似db中的表
* retention policy 存储策略
* point 类似于db中的数据行,由time/field/tags组成
* database 顾名思义
* series 表示数据可在图表上展示成多少条线
* RP Retention Policies 保留策略,某个库可有多个保留策略,但保留策略必须独一无二
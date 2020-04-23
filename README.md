# go-api-consumer

## start

### 1. run docker kafka
1.设置HOST文件：（xxx为你本机的ip，不能写成127.0.0.1）
    xxx test-kafka

2.运行以下命令：
```
$ docker-compose -f example/docker-compose.yml up -d
```

### 2. 运行go-api-consumer
```
go run .
```

### 3.download shell kafka
下载地址： http://kafka.apache.org/quickstart

## 3.1 send message
打开shell，并cd到 ...../kafka_2.12-2.3.0/bin/windows
```bash
./kafka-console-producer.bat --broker-list localhost:9092 --topic fruit
> {"authToken":"","createdAt":"2019-07-30T23:51:30.5899707Z","payload":{"name":"apple"},"requestId":"","status":"FruitCreated"}
```

## 4.accept message

Open mysql, you will see the new data in your fruit database

参考：
https://www.cnblogs.com/mignet/p/window_kafka.html
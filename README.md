# go-api-consumer

## prepare

install kafka
```
https://www.cnblogs.com/mignet/p/window_kafka.html
https://www.apache.org/dyn/closer.cgi?path=/kafka/2.3.0/kafka_2.12-2.3.0.tgz
```
install mysql
- start local mysql(port is 3308)
- create database handerly named fruit

## send message

```bash
./kafka-console-producer.bat --broker-list localhost:9092 --topic fruit
> {"authToken":"","createdAt":"2019-07-30T23:51:30.5899707Z","payload":{"name":"apple"},"requestId":"","status":"FruitCreated"}
```

## accept message

Open mysql, you will see the new data in your fruit database
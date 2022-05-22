# WEEK-9

1. 总结几种 socket 粘包的解包方式：fix length/delimiter based/length field based frame decoder。尝试举例其应用。

| 方式 | 说明 | 应用 |
|:---|:---|:---|
|fix length|按照固定长度接收数据，内容如果小于定长需要填充，造成浪费|适用于数据长度固定的场景, 比如指令发送|
|delimiter based|使用特殊符号做为数据包边界，选用的特殊符号不能在 Body 中出现|HTTP 中用`/r/n`划分请求行、请求头|
|length field based frame decoder|在数据包固定位置添加包长度信息，适用于变长的消息传递|HTTP 中 Body 部分由 header 中的 content-length 指明数据长度|

2. 实现一个从 socket connection 中解码出 goim 协议的解码器。

代码:
1. 客户端: [client.go](./client.go)
2. 服务端: [server.go](./server.go)
3. main: [main.go](./main.go)

说明:

    客户端使用 goim 协议写入数据;

    服务端使用自定义解码器进行解析;

    


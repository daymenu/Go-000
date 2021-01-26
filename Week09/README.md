# 作业

## 代码介绍

server 端的代码可以同时为n个客户端服务，每个客户端会启动4个 goroutine

- handle goruntine 一个连接的封装
- read goruntine 负责读取客户端发送的数据
- wirte goruntine 负责返回客户端发送的数据
- pkg.Heart gorutine 负责心跳探活

实现

- 服务端接收到客户端 exit 的君子协定退出
- 客户端没有发出君子协定就退出，心跳向客户端写入，如果写入失败就退出

问题

- 老师们说的心跳是不是我写的这样呢？
- 实现第二点 “如果写入失败就退出” 是有问题的，是不是应该用 errors.Is 或者 其他方式判断下什么类型的错误退出？

用 Go 实现一个 tcp server ，用两个 goroutine 读写 conn，两个 goroutine 通过 chan 可以传递 message，能够正确退出
以上作业，要求提交到 GitHub 上面，Week09 作业地址：
https://github.com/Go-000/Go-000/issues/82

-----end-----

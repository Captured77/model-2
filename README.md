# Model-2 作业

## 要求：

    编写一个HTTP服务器

* 接收客户端的request，并将request中带的header写入response header
* 读取当前系统的环境变量中的VERSION配置，并写入response header
* Server端记录访问日志包括客户端IP，HTTP返回码，输出到server端的标准输出
* 当访问localhost/healthz时，应返回200

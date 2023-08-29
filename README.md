# car-tcp
golang 32960 project
本项目是用golang针对32960定制化开发项目，TCP框架选用zinx，在原来的基础上修改粘包方法及messageId类型完成应对。

**使用到的技术栈：**
golang、gorm、gin框架、zinx框架、dapr、消息队列（redis、rabbit、kafka任选其一）、redis、mysql、clickhouse、http、tcp

**项目架构图如下：**
![TCP服务器](https://github.com/wangcheng0509/car-tcp/assets/34881771/2d145e14-46ff-4ff4-9753-2eb4ce15451f)

**运行说明：**
1、IDE使用VSCode，并安装dapr扩展：
  https://v1-5.docs.dapr.io/zh-hans/getting-started/install-dapr-selfhost/
  https://v1-5.docs.dapr.io/zh-hans/developing-applications/ides/vscode/
2、依次打开：tcoService、tcpConsumer、client-test、openService（对外开放终端命令下发）

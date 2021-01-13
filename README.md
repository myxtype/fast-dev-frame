# fast-dev-frame

Go后端快速开发框架

# 组件

- Middleware：gin
- ORM：gorm
- Redis：go-redis
- 日志：zap
- 配置：viper
- 任务调度：cron.v3
- 接口限流器：tollbooth
- 类型转换：cast

# How to use

你需要克隆或者下载到本地来使用。

本项目是一个快速开发框架，是一个组件的集合，并非一个独立的框架，使用它，你可能需要一些学习成本。

目录介绍：

- admin 管理后台接口目录，存放管理后台的`controller`和路由。
- cmd 可执行文件目录，所有编译后的执行文件都在此处。
- conf 配置文件目录，请将你的配置结构体放到`conf/config.go`中。
- job 定时任务目录，将定时任务的调度放到`job/bootstrap`中。
- model 数据模型目录。
- pkg 组件目录，第三方组件存放的目录，适合将别人的代码Copy过来，修改一下。
- rest 为前台提供接口的目录，放控制器`controller`和路由的地方。
- store 数据库存储目录，mysql的操作放到`store/mysql`中，redis操作放到`store/redisdb`中。
- service 具体业绩逻辑目录。
- worker 任务目录，队列，或者需要一直跑的任务。

你可以添加更多的目录来表示你的服务，例如：`websocket`表示你为用户提供长连接服务，你需要遵循规范，在`cmd`目录下创建与服务名称相同的文件夹，并提供`main.go`。

# pkg目录介绍

- ecode 错误码定义
- exporter csv文件导出助手
- grace 优雅的启动和关闭
- i18n 国际化
- logger 日志
- middleware 中间件
- queue 队列助手
- queworker 队列worker
- worker 管理器

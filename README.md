# fast-dev-frame
Go后端快速开发框架

# 组件

- Middleware：gin
- ORM：gorm
- Redis：go-redis
- 日志：zap
- 配置：viper

# How to use

你需要克隆或者下载到本地来使用。

本项目是一个快速开发框架，是一个组件的集合，并非一个独立的框架，需要一些学习成本。

目录介绍：

- cmd 可执行文件目录，所有编译后的执行文件都在此处。
- conf 配置文件目录，请将你的配置放到`conf/config.go`中。
- job 定时任务目录，创建一个定时任务，调度放到`job/bootstrap`中。
- models 数据模型目录，所有的数据表定义都放到`models.go`中，数据库操作放到`models/mysql`目录中。
- pkg 自定义组件，第三方组件存放的目录，适合将别人的代码Copy过来，修改一下。
- rest 默认的接口目录，放控制器`controller`和路由的地方。
- service 具体业绩逻辑目录。
- worker 任务目录，队列，或者需要一直跑的任务。

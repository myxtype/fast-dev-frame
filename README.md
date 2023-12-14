# fast-dev-frame

Go后端快速开发框架

# 技术

- 数据库：gorm
- 缓存Redis：go-redis
- 定时任务：github.com/robfig/cron
- 配置：github.com/spf13/viper
- 日志：go.uber.org/zap
- http服务：github.com/gin-gonic/gin
- JWT权限认证

开箱即用的gin中间件缓存、一键编译、优化的项目目录结构

# 目录

- app：应用目录，里面存放子应用
- build：存放可执行文件，编译后的可执行文件存放在此
- cmd：程序入口，多个程序多个目录
- conf：配置定义，将需要的配置定义到这里，然后直接使用
- model：数据模型定义，使用gorm一键迁移
- pkg：工具包
- service：业务实现
- store：提供数据访问，目前开箱即用的db和redis目录

# 编译

- make build：编译本机可执行的
- make build-linux：编译linux的
- make all：编译、打包


# conf

默认读取当前执行文件目录下的`config.toml`文件，可使用`-conf <File Path>`来指定`config.toml`配置文件的目录。

示例：

```shell
./main -conf = /root
```

就会读取`/root/config.toml`文件。

# 配置示例

config.toml

```toml
# 配置文件
# 需要哪些配置看conf/config.go
##################################

# 日志配置
[Logger]
Level = "debug"
Target = "console"
Filename = "./log/stream.log"

# Rest服务配置
[RestServer]
Addr = "127.0.0.1:8000"

# Admin服务配置
[AdminServer]
Addr = "127.0.0.1:8001"

# 数据库配置
[DataSource]
Addr = "127.0.0.1:3600"
Database = "frame"
User = "root"
Password = "ZEvmL3rFpRXd6BIX"
MaxIdle = 10
MaxOpen = 100
Migrate = false

# Redis配置
[Redis]
Addr = "127.0.0.1:6379"
Password = ""
DB = 0
```

# 功能

- 配置改动监控（配置修改不用重启程序）
- 高性能读取（配置缓存到内存中，读取超快）
- 配置结构化（将所有配置都定义结构体，代码清晰明了）

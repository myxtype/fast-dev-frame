# conf

默认读取当前执行文件目录下的`config.toml`文件，可使用`-conf <File Path>`来指定`config.toml`配置文件的目录。

示例：

```shell
./main -conf = /root
```

就会读取`/root/config.toml`文件，需要在入口处执行`flag.Parse()`。

# 配置示例

config.toml
```toml
# 配置文件
##########################

# Rest服务
[RestServer]
Addr = "127.0.0.1:8000"
```

# 功能

- 配置改动监控（配置修改不用重启程序）
- 高性能读取（配置缓存到内存中，读取超快）
- 配置结构化（将所有配置都定义结构体，代码清晰明了）

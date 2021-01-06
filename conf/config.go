package conf

// 全局配置
type gbeConfig struct {
	Logger      loggerConfig
	RestServer  serverConfig
	AdminServer serverConfig
	DataSource  dataSourceConfig
	Redis       redisConfig
}

// 日志配置
type loggerConfig struct {
	Level    string // 打印日志的等级 debug/info/warn/error/panic/fatal
	Target   string // 日志打印目标："file" or "console"
	Filename string // 日志文件
}

// 接口服务配置
type serverConfig struct {
	Addr string // 监听地址
}

// 数据库配置
type dataSourceConfig struct {
	Addr        string // 地址
	Database    string // 数据库
	User        string // 用户
	Password    string // 密码
	MaxIdle     int    // 最大空闲数
	MaxOpen     int    // 最大连接数
	Migrate     bool   // 是否执行迁移
	LogDisabled bool   // 是否禁用SQL日志
}

// Redis配置
type redisConfig struct {
	Addr     string // 地址
	Password string // 密码
	DB       int    // 数据库
}

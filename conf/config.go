package conf

// 全局配置
type gbeConfig struct {
	RestServer  serverConfig
	AdminServer serverConfig
	DataSource  dataSourceConfig
	Redis       redisConfig
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

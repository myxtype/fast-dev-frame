package conf

import (
	"flag"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

var (
	configPath string
	config     gbeConfig
	configOnce sync.Once
)

func init() {
	flag.StringVar(&configPath, "conf", "", "Config file path. This path must include config.toml file.")
}

func GetConfig() *gbeConfig {
	configOnce.Do(func() {
		flag.Parse() // 解析参数

		viper.SetConfigName("config") // 配置文件名称
		viper.SetConfigType("toml")   // 配置文件类型
		if configPath == "" {
			if p, err := execPath(); err == nil {
				for _, n := range p {
					viper.AddConfigPath(n)
				}
			} else {
				panic(err)
			}
		} else {
			viper.AddConfigPath(configPath)
		}
		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}
		// 设置配置文件监听
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			if e.Op != fsnotify.Remove {
				if err := viper.Unmarshal(&config); err != nil {
					log.Println("Reload config error", err)
				}
			}
		})
		if err := viper.Unmarshal(&config); err != nil {
			panic(err)
		}
	})
	return &config
}

func execPath() (p []string, err error) {
	p = []string{"./"}
	if _, currentPath, _, ok := runtime.Caller(0); ok {
		p = append(p, filepath.Dir(currentPath))
	}
	if t, err := filepath.Abs(filepath.Dir(os.Args[0])); err == nil {
		p = append(p, t)
	}
	return
}

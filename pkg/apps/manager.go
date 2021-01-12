package apps

import (
	"frame/pkg/logger"
	"sync"
)

type App interface {
	Start(sw *sync.WaitGroup)
}

// 子应用管理工具
// 同一个执行文件中启动多个服务，结束进程时等待服务全部结束
type AppManager struct {
	sw   *sync.WaitGroup
	list []App
}

func NewAppManager() *AppManager {
	return &AppManager{
		sw:   &sync.WaitGroup{},
		list: make([]App, 0, 10),
	}
}

// Add 添加一个子应用
func (am *AppManager) Add(a App) {
	am.list = append(am.list, a)
}

// Start 启动全部子应用
func (am *AppManager) Start() {
	am.sw.Add(len(am.list))
	for _, app := range am.list {
		go func(a App) {
			defer func() {
				err := recover()
				if err != nil {
					logger.Sugar.Error(err)
				}
			}()
			a.Start(am.sw)
		}(app)
	}
}

// Wait 等待子应用执行结束
func (am *AppManager) Wait() {
	am.sw.Wait()
}

package apps

import "sync"

// 函数类子子应用
type FuncApp struct {
	f func()
}

func NewFuncApp(f func()) *FuncApp {
	return &FuncApp{f: f}
}

func (a *FuncApp) Start(sw *sync.WaitGroup) {
	defer sw.Done()
	a.f()
}

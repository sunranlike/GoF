package Option

import (
	"context"
	"fmt"
	"log"
	"sync"
)

type Option interface {
	Initialize()
	Afterinit()
}

// DoSomething 遍历选项,执行选项的初始化工作和收尾工作,优雅
func DoSomething(ctx *context.Context, op ...Option) {
	for _, o := range op {
		o.Initialize()
		defer o.Afterinit()
	}

}

type MyOption struct {
	lock sync.Mutex
	log  log.Logger
}

func (my *MyOption) Initialize() {
	fmt.Println("do some initialization")
	my.lock.Lock()

}
func (my *MyOption) Afterinit() {
	fmt.Println("do some Afterinit")
	my.lock.Unlock()
}

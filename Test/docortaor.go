package Test

import (
	"context"
	"fmt"
	"time"
)

type Myctx struct {
	context context.Context
	next    Handler
}
type Handler func(c *Myctx)

func Record(H Handler, ctx1 *Myctx) Handler {

	return func(ctx *Myctx) {
		timeNow := time.Now()
		H(ctx1)
		fmt.Println(time.Since(timeNow))
	}
}

func Test(ctx *Myctx) {
	fmt.Println("test")
	time.Sleep(4 * time.Second)
}

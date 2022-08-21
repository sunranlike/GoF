package Test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestCtx(t *testing.T) {
	ctx, c := context.WithTimeout(context.Background(), 10*time.Second)
	defer c()
	ctx2 := context.WithValue(ctx, "test", 1)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				continue
			default:
				time.Sleep(1 * time.Second)
				fmt.Println(ctx2.Value("test"))
			}

		}
	}(ctx2)
	time.Sleep(10 * time.Second)

}

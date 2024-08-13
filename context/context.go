package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go work(ctx)
	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(time.Second)
	fmt.Println("继续")
	time.Sleep(100 * time.Second)

}

func work(ctx context.Context) {
	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("取消")
			return
		default:
			fmt.Println(i)
			time.Sleep(time.Second)
		}

	}
}

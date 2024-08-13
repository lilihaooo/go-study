package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	down := doLongRunningTask(context.Background())
	select {
	case <-down:
		fmt.Println("do long running task")
	case <-time.After(5 * time.Second):
		fmt.Println("time out")
	}

}

func doLongRunningTask(ctx context.Context) chan struct{} {
	done := make(chan struct{}, 1)
	time.Sleep(12 * time.Second)
	//done <- struct{}{}
	close(done) // 关闭down后 case <-down 会被执行
	return done
}

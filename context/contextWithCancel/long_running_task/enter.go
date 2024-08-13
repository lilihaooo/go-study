package main

import (
	"context"
	"fmt"
	"time"
)

/*
此demo功能： 当执行一个10s的任务时， 在第5s的时候主动通过context.WithCancel取消任务
*/

func doLongRunningTask(ctx context.Context) chan struct{} {
	done := make(chan struct{})
	go func() {
		select {
		case <-time.After(10 * time.Second): // 假设任务需要10秒完成
			fmt.Println("Long running task completed")
			close(done)
		case <-ctx.Done():
			fmt.Println("Long running task canceled")
		}
	}()
	return done
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		done := doLongRunningTask(ctx)
		select {
		case <-done:
			// 任务完成或取消
			fmt.Println("任务完成")
		case <-time.After(5 * time.Second):
			// 5秒后取消任务
			cancel()
		}
	}()
	// 等待一段时间来观察结果
	time.Sleep(20 * time.Second)
}

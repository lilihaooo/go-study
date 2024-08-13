package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
描述： 在多个go执行一个查找任务时， 如果其中一个goroutine找到了目标， 就关闭其他go
思路： 开启5个go
*/

func fetchData(ctx context.Context, source string) chan struct{} {
	done := make(chan struct{})
	go func() {
		select {
		case <-time.After(2 * time.Second): // 模拟数据检索
			fmt.Printf("Data fetched from %s\n", source)
			close(done)
		case <-ctx.Done():
			fmt.Printf("Data fetch from %s canceled\n", source)
		}
	}()
	return done
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	sources := []string{"SourceA", "SourceB", "SourceC"}

	for _, source := range sources {
		wg.Add(1)
		go func(source string) {
			defer wg.Done()
			done := fetchData(ctx, source)
			select {
			case <-done:
				// 数据检索成功或取消
			case <-time.After(1 * time.Second):
				// 假设我们在这里检测到SourceA已经返回数据，所以取消其他源
				if source == "SourceA" {
					cancel()
					return
				}
			}
		}(source)
	}

	wg.Wait()
	fmt.Println("All data fetches complete or canceled")
}

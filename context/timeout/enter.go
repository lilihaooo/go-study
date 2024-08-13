package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	test()

}

func test() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	time.Sleep(2 * time.Second)

	work(ctx)

}

func work(ctx context.Context) {
	if ctx.Err() != nil {

		fmt.Println(ctx.Err())
		return
	}

	fmt.Println("ok")
}

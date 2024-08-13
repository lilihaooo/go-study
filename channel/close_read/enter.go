package main

import (
	"fmt"
	"sync"
)

/*
结论： 读取一个已经关闭的队列 那么会读取到零值
注意： for 中读取一个已经关闭的队列， 小心死循环
*/

func main() {
	ch := make(chan int, 5)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	producer(ch, wg)
	go consumer(ch, wg)
	wg.Wait()

}

func producer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		select {
		case ch <- i:
		default:
			fmt.Println("channel full")
			close(ch)
			return
		}
	}
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		v := <-ch
		fmt.Println(v)
	}
}

func test() {

}

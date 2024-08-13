package main

import (
	"fmt"
	"time"
)

func main() {
	waitc := make(chan struct{})
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			time.Sleep(time.Second)
			if i == 5 {
				close(waitc)
			}
		}

	}()
	<-waitc
}

package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var num int

var gorpCh = make(chan struct{}, 10)

// var gorpCh chan struct{}
var pathName = "test"
var wg sync.WaitGroup
var numCh = make(chan struct{})

func main() {
	start := time.Now()
	wg.Add(1)
	go func() {
		defer fmt.Println(222)
		defer wg.Done()

		for range numCh {
			num++
		}
	}()

	path := "/Users/lihao"

	search(path)
	wg.Wait()
	defer close(numCh)

	end := time.Since(start)
	fmt.Println("Time: ", end)
	fmt.Println(num)
}

func search(path string) {

	defer wg.Done()
	defer func() { <-gorpCh }()
	gorpCh <- struct{}{}

	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, file := range files {
		numCh <- struct{}{}
		if file.Name() == pathName {
			numCh <- struct{}{}
			//time.Sleep(5000 * time.Millisecond)
		}
		if file.IsDir() {
			wg.Add(1)
			go search(path + "/" + file.Name())
		}
	}

}

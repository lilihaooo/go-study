package main

import (
	"fmt"
	"os"
	"time"
)

var query = "test"
var matches int

var workerCount int = 0
var maxWorkerCount = 100
var searchRequest = make(chan string)
var workerDown = make(chan bool)
var foundMatch = make(chan bool)

func main() {
	start := time.Now()
	workerCount = 1
	go search("/Users/lihao/", true)
	waitForWorker()
	elapsed := time.Since(start)
	fmt.Printf("Search took %s\n", elapsed)
	fmt.Println(matches)
}

func waitForWorker() {
	for {
		select {
		case page := <-searchRequest:
			workerCount += 1
			go search(page, true)
		case <-workerDown:
			workerCount -= 1
			if workerCount == 0 {
				return
			}
		case <-foundMatch:
			matches += 1
		}
	}
}

func search(path string, master bool) {
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		foundMatch <- true
		name := file.Name()
		if name == query {
			foundMatch <- true
		}
		if file.IsDir() {
			if workerCount < maxWorkerCount {
				searchRequest <- path + name + "/"
			} else {
				search(path+name+"/", false)
			}

		}
	}
	if master {
		workerDown <- true
	}

}

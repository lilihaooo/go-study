package main

import "fmt"

var i = 43

func main() {
	defer fmt.Println("world")
	defer fmt.Println("hello")

	defer fmt.Println("333")
	if i > 7 {
		return
	}
	defer func() {
		fmt.Println("world")
		fmt.Println("hello")
	}()

}

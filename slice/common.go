package main

import (
	"fmt"
)

func main() {

	var s []int
	s = nil // 赋值nil依然可以使用
	for i := 0; i < 10; i++ {
		s = append(s, i)
		fmt.Println(len(s))
	}

}

package main

import (
	"fmt"
	"time"
)

type P struct {
	Name string
	any
}

func main() {

	//d := "1h"
	//dr, err := time.ParseDuration(d)
	//
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(dr)

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	var p P
	p.Name = "hah"
	p.any = 3
	fmt.Println(p)

}

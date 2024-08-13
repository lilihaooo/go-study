package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	d := "  1d 23m44s   "
	d = strings.TrimSpace(d)
	//fmt.Println(d)
	work(d)

}

func work(d string) {
	if strings.Contains(d, "d") {
		index := strings.Index(d, "d")
		hour, _ := strconv.Atoi(d[:index])
		fmt.Println(hour)

		dr := time.Hour * 24 * time.Duration(hour)

		ndr, err := time.ParseDuration(d[index+1:])
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(dr + ndr)
	}
}

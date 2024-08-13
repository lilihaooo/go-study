package hah

import (
	"fmt"
	"time"
)

var Name = "lihao"

func Tt() {
	func() {
		for i := 0; i < 20; i++ {
			fmt.Println(Name)
			time.Sleep(1 * time.Second)
		}
	}()
}

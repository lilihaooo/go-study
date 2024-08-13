package main

import (
	"go_study/bianliangfuzhi/hah"
	"time"
)

func main() {

	go hah.Tt()

	go func() {
		time.Sleep(5 * time.Second)
		hah.Name = "syy"
	}()

	select {}

}

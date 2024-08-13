package main

import (
	"fmt"
)

/*
在这个示例中，recoverFromPanic 函数使用 recover 函数来捕获和处理恐慌。
main 函数中的 panic 触发了恐慌，而 recoverFromPanic 函数从中恢复，允许程序继续执行。
*/

func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}

func main() {
	defer recoverFromPanic()

	panic("This is a panic!")

}

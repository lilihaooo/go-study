package main

import (
	"errors"
	"fmt"
)

func main() {
	fmtErrorf()
}

/*
使用 fmt.Errorf 向错误添加上下文
*/
func fmtErrorf() {
	err := errors.New("original error")
	contextErr := fmt.Errorf("additional context: %w", err)
	fmt.Println(contextErr)
}

// 使用 errors 包进行错误包装
func studyErrors() {

}

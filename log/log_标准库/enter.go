package main

import (
	"errors"
	"fmt"
	"log"
)

/*
log提供了三组函数：

1. Print/Printf/Println：正常输出日志；

2. Panic/Panicf/Panicln：输出日志后，以拼装好的字符串为参数调用panic；

3. Fatal/Fatalf/Fatalln：输出日志后，调用os.Exit(1)退出程序。 命名比较容易辨别，带f后缀的有格式化功能，带ln后缀的会在日志后增加一个换行符。

注意，上面的程序中由于调用log.Panicf会panic，所以log.Fatalf并不会调用
*/

/*
选项

• Ldate：输出当地时区的日期，如2020/02/07；

• Ltime：输出当地时区的时间，如11:45:45；

• Lmicroseconds：输出的时间精确到微秒，设置了该选项就不用设置Ltime了。如11:45:45.123123；

• Llongfile：输出长文件名+行号，含包名，如github.com/darjun/go-daily-lib/log/flag/main.go:50；

• Lshortfile：输出短文件名+行号，不含包名，如main.go:50；

• LUTC：如果设置了Ldate或Ltime，将输出 UTC 时间，而非当地时区。
*/

func division(x float32, y float32) (float32, error) {
	if y == 0 {
		return 0, errors.New("can't divide by zero")
	}

	return x / y, nil
}

func main() {
	var x float32 = 11
	var y float32

	res, err := division(x, y)

	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)
	log.SetPrefix("Debug: ")

	if err != nil {
		log.Print(err)
		return
	}

	fmt.Println(res)
}

package main

import "fmt"

func main() {
	var a int = 100
	var b int = 200

	change(&a, &b)
	fmt.Println(a, b)

}

func change(x, y *int) {
	x, y = y, x
}

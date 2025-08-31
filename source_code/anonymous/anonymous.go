package main

import "fmt"

func main() {
	var a func()
	a = func() {
		fmt.Println("a()...")
	}
	a()

	// 求两个数的和
	b := func(a int, b int) int {
		return a + b
	}(1, 2)
	fmt.Println(b)
}

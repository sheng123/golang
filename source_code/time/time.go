package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	str := "hello"
	for i := range 1000 {
		str += strconv.Itoa(i)
	}
}

func main() {
	t := time.Now()
	test()
	cost := time.Since(t)
	fmt.Printf("函数执行耗时: %v\n", cost)
}

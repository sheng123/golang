package main

import (
	"fmt"
	"strings"
)

// 1.返回的匿名函数和 suffix 一起组成了一个闭包
// 2.使用传统的方法也可以实现这个需求，但是需要传入多个参数，
// 使用闭包的话，只需要传入 .jpg 一次就可以完成

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}

		return name
	}
}

func main() {
	f := makeSuffix(".jpg")

	fmt.Println("card =", f("card"))
	fmt.Println("test =", f("test.jpg"))
}

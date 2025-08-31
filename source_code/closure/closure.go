package main

import "fmt"

// 闭包是有函数和相关的引用环境（函数定义时所能访问的外部变量）
// 组合而成的实体。当函数引用了其外部变量时，这些变量的生命周期
// 会与闭包绑定，即使外部函数返回，这些变量也不会被销毁，而是继续被
//闭包持有和操作

// 类似面向对象的类和实例一样，一旦实例化以后，对应的参数就定下来了
// 当实例销毁的时候，才会一起销毁，否则会一直存在

func multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	myCounter := counter()
	fmt.Println(myCounter())
	fmt.Println(myCounter())

	double := multiplier(2)
	triple := multiplier(3)
	fmt.Println(double(5)) // 输出: 10
	fmt.Println(triple(5)) // 输出: 15
}

package main

import (
	"fmt"
	"funcinit/utils"
)

var age = test()

func test() int {
	fmt.Println("test()...")
	return 90
}

func init() {
	fmt.Println("init()...")
}

func main() {
	fmt.Println("main()...age = ", age)
	fmt.Println("main()...utils.Name = ", utils.Name)
	fmt.Println("main()...utils.Age = ", utils.Age)
}

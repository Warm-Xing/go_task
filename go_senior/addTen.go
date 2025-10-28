package main

import "fmt"

// 定义接收整数指针的函数，将指针指向的值增加10
func addTen(num *int) {
	*num += 10 // 通过指针间接访问并修改原始值
}

func main() {
	x := 5
	fmt.Println("修改前的值:", x) // 输出：修改前的值: 5

	addTen(&x) // 传递变量x的地址

	fmt.Println("修改后的值:", x) // 输出：修改后的值: 15
}

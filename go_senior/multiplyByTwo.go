package main

import "fmt"

// multiplyByTwo 接收切片指针，将每个元素乘以2
func multiplyByTwo(slicePtr *[]int) {
	// 解引用切片指针获取切片本身，然后遍历修改元素
	for i := range *slicePtr {
		(*slicePtr)[i] *= 2 // 注意括号：先解引用指针，再访问切片元素
	}
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println("修改前:", nums) // 输出：修改前: [1 2 3 4]

	multiplyByTwo(&nums) // 传递切片的指针

	fmt.Println("修改后:", nums) // 输出：修改后: [2 4 6 8]
}

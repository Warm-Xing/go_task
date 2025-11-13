package main

import "fmt"

func twoSum(nums []int, target int) []int {
	//起一个map函数，用于把数组中的数存为键，循环次数存为值。
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if idx, exists := numMap[complement]; exists {
			//返回加起来等于target数的index值。
			return []int{idx, i}
		}
		numMap[num] = i
	}

	return nil
}

func main() {
	fmt.Println(twoSum([]int{1, 3, 5, 7, 9}, 10))
}

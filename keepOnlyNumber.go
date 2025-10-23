package main

import "fmt"

func keepOnlyNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
			fmt.Println("不等时候去重过程:", nums[:i+1])
		}
	}

	return i + 1
}

func main() {
	fmt.Println(keepOnlyNumber([]int{1, 1, 2, 2, 6}))
}

package main

import "fmt"

func lastPlusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	// 如果所有位都是9，需要在最前面加1
	return append([]int{1}, digits...)
}

func main() {
	fmt.Println(lastPlusOne([]int{1, 2, 3, 4, 6}))
}

package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}
	match := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != match[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("([])"))
	// 可以在这里添加测试代码
	// 示例: fmt.Println(isValid("()[]{}")) // 输出: true
}

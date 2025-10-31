package main

import "fmt"

// Person 结构体
type person struct {
	name string
	age  int
}

// Employee 结构体组合 Person
type employee struct {
	person
	employeeID string
}

// Employee 的方法
func (e employee) PrintInfo() {
	fmt.Println("员工信息:")
	fmt.Println("  姓名:", e.name)
	fmt.Println("  年龄:", e.age)
	fmt.Println("  员工ID:", e.employeeID)
}

func main() {
	emp := employee{
		person: person{
			name: "张三",
			age:  30,
		},
		employeeID: "E001",
	}

	emp.PrintInfo()
}

package main

import (
	"fmt"
	"math"
)

// Shape 接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle 正三角形结构体
type Rectangle struct {
	Width  float64
	Height float64
}

// Rectangle 实现 Shape 接口
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle 结构体
type Circle struct {
	Radius float64
}

// Circle 实现 Shape 接口
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func printShapeInfo(s Shape) {
	fmt.Printf("面积: %.2f, 周长: %.2f\n", s.Area(), s.Perimeter())
}

//	拓展正三角形triangle 实现 Shape 接口
//
// 拓展triangle 结构体
type triangle struct {
	base   float64
	Height float64
}

// 拓展正三角形triangle面积 接口
func (t triangle) Area() float64 {
	return (t.base * t.Height) / 2
}

// 拓展正三角形triangle周长
func (t triangle) Perimeter() float64 {
	return 3 * t.base
}

// 拓展正三角形定义radians函数
func radians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func main() {
	rect := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 4}

	tri := triangle{base: 5, Height: 2.5 * math.Tan(radians(60))}

	fmt.Println("矩形:")
	printShapeInfo(rect)

	fmt.Println("圆形:")
	printShapeInfo(circle)

	fmt.Println("正三角形:")
	printShapeInfo(tri)
}

package main

import (
	"fmt"
	"math"
)

// 这里是一个几何体的基本接口。
type geometry interface {
	area() float64
	perim() float64
}

// 让 rect2 和 circle 实现这个接口
type rect2 struct {
	width, height float64
}
type circle struct {
	radius float64
}

// 在Go中实现一个接口，我们只需要实现接口中的所有方法。这里我们让 rect2 实现了 geometry 接口。
func (r rect2) area() float64 {
	return r.width * r.height
}
func (r rect2) perim() float64 {
	return 2*r.width + 2*r.height
}

// circle 的实现。
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 如果一个函数的参数是接口类型,那么函数就可以用这个接口中的方法,并且可以用在任何实现了这个接口的类上
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
func main() {
	r := rect2{width: 3, height: 4}
	c := circle{radius: 5}
	// 结构体类型 circle 和 rect2 都实现了 geometry接口，所以我们可以使用它们的实例作为 measure 的参数。
	measure(r)
	measure(c)
}

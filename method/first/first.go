package main

import (
	"fmt"
	"math"
	"time"
)

type Point struct {
	x, y float64
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(time.Now())
	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q))
	perim := Path {
		{1, 2},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())

	//函数和方法区别有没有接收器

	funcValue := p.Distance //menthod value, variable.function(p.Distance) 选择器， p 接收器变量
	fmt.Println(funcValue(q))

	funcExpression := Point.Distance //method expression type.function
	fmt.Println(funcExpression(p, q))//第一个参数做接收器，所以有时候函数无参数需传入一个参数
}

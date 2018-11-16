package main

import (
	"fmt"
	"image/color"
	"math"
	"sync"
)

type point struct {
	x, y float64
}
//如果结构体定义了匿名字段，那么他拥有这些字段的所有方法和变量，方法名按级别索引，本结构体，匿名包
//级方法，同级方法冲突会报错
type coloredPoint struct {
	point //内嵌字段，相当于x, y,可以不通过point 调用,如果point是指针类型*point 不能直接访问xy
	col color.RGBA
}

func (px point) Distance(py point) float64 {
	return math.Hypot(py.x - px.x, py.y - px.y)
}

type ColorPoint struct {
	*point
	col color.RGBA
}

func main() {
	var cp coloredPoint
	cp.x = 1
	fmt.Println(cp.point.x)
	cp.y = 2
	fmt.Println(cp.y)

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = coloredPoint{point{1, 1}, red}
	var q = coloredPoint{point{5, 4}, blue}
	fmt.Println(p.Distance(q.point))

	P := ColorPoint{&point{1, 2}, red}
	Q := ColorPoint{&point{5, 4}, blue}
	fmt.Println(*P.point, *Q.point)
	Q.point = P.point
	fmt.Println(*(P.point), *(Q.point))
	p.point = q.point
	fmt.Println(p.point, q.point)
}

//注意下面两个结构的联系
var (
	mu sync.Mutex
	mapping= make(map[string]string)
)

func Lookup(key string) string {
	mu.Lock()
	v := mapping[key]
	mu.Unlock()
	return v
}


var cache = struct {
	sync.Mutex
	mapping map[string]string
} {
	mapping : make(map[string]string),
}


func  lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}
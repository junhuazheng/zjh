//自定义struct案例及其用法示例

package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) Distence(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y) //math.Hypot 对给定的直角三角形的两个直角边，求其斜边长度。返回斜边值。
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}

	fmt.Println(p.Distence(q))

}

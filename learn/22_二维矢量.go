package main

import "math"

type Vec2 struct {
	X float32
	Y float32
}

//加
func (v Vec2) Add(ot Vec2) Vec2 {
	
	return Vec2{
		v.X + ot.X
		v.Y + ot.Y
	}
}

//减
func (v Vec2) Sub(ot Vec2) Vec2 {
	
	return Vec2 {
		v.X - ot.X
		v.Y - ot.Y
	}
}

//乘
func (v Vec2) Scale(s float32) Vec2 {
	return Vec2{v.X * s, v.Y * s}
}

//距离
func (v Vec2) Dis(ot Vec2) float32 {
	dx := v.X - ot.X
	dy := v.Y - ot.Y
	
	return float32(math.Sqrt(float64(dx *dx + dy * dy))) //math.Sqrt是开方函数
}

//插值
func (v Vec2) Nor() Vec2 {
	mag := v.X*v.X + v.Y*v.Y
	if mag > 0 {
		one := 1 / float32(math.Sqrt(float64(mag)))
		return Vec2{v.X * one, v.Y * one}
	}
	
	return Vec2{0, 0}
}
//在计算机中，小对象由于值复制时的速度较快，所以适合使用非指针接收器。
//大对象因为复制性能较低，适合使用指针接收器，在接收器和参数间传递时不进行复制，只是传递指针。

//二维矢量模拟玩家移动
//在游戏中，一般使用二维矢量保存玩家的位置。使用矢量运算可以计算出玩家移动的位置。
//本例子中，收线实现二维矢量对象，接着构造玩家对象，最后使用矢量对象和玩家对象共同模拟玩家移动的过程

//矢量是数学中的概念，二维矢量拥有两个方向的信息，可以同时进行加、减、乘（缩放）、距离、单位化等计算。
//在计算机中，使用拥有X和Y两个分量的Vec2结构体视线数学中二维向量的概念：

package main

import "math"

type Vec2 struct {
	X, Y float32
}

//加
func (v Vec2) Add(other Vec2) Vec2 {

	//使用自身Vec2和通过Add()方法传入Vec2进行相加。相加后，结果以返回值形式返回，不会修改Vec2的成员
	return Vec2{
		v.X + other.X,
		v.Y + other.Y,
	}
}

//减
func (v Vec2) Sub(other Vec2) Vec2 {

	return Vec2{
		v.X - other.X,
		v.Y - other.Y,
	}
}

//乘（缩放或叫矢量乘法，是对矢量的每个分量乘上缩放比，Scale()方法传入一个参数同时成两个分量，表示这个缩放是一个等比缩放。
func (v Vec2) Scale(s float32) Vec2 {

	return Vec2{v.X * s, v.Y * s}
}

//距离
func (v Vec2) DistanceTo(other Vec2) float32 {
	dx := v.X - other.X
	dy := v.Y - other.Y

	//math.Sqrt()是开方函数，参数是float64，使用时需要转换，返回值也是flota64，需要转换回float32
	return float32(math.Sqrt(float64(dx*dx + dy*dy)))
}

//插值（定义矢量单位化）
func (v Vec2) Normalize() Vec2 {
	mag := v.X*v.X + v.Y*v.Y
	if mag > 0 {
		oneOverMag := 1 / float32(math.Sqrt(float64(mag)))
		return Vec2{v.X * oneOverMag, v.Y * oneOverMag}
	}

	return Vec2{0, 0}
}

package main

import (
	"fmt"
	"reflect"
	"testing"
)

//原生结构体的赋值过程：

//声明一个结构体，拥有一个字段
type data struct {
	Hp int
}

func BenchmarkNativeAssign(b *testing.B) { //使用基准化测试的入口

	v := data{Hp: 2} //实例化并赋值

	//停止基准测试的计时器
	b.StopTimer()

	//重置基准测试计时器数据
	b.ResetTimer()

	//重新启动基准测试计时器
	b.StartTimer()

	//以上三步：由于测试的重点必须放在赋值上，因此需要极大程度地降低其他代码的干扰
	//于是在赋值完成后，将基准测试的计时器复位并重新开始

	//根据基准测试数据进行循环测试
	for i := 0; i < b.N; i++ {

		//结构体成员赋值测试
		v.Hp = 3
	}
}

//使用反射访问结构体成员并赋值的过程
func BenchmarkReflectAssign(b *testing.B) {

	v := data{Hp: 2}

	//取v的地址并转为反射值对象。此时值对象里的类型为*data，使用值得Elent()方法取元素，获得data的反射值对象
	vv := reflect.ValueOf(&v).Elem()

	//根据名字取结构体成员的反射对象。
	f := vv.FieldByName("Hp")

	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		//使用反射对象的SetInt()方法，给data结构的Hp字段设置数值为3
		f.SetInt(3)
	}
}

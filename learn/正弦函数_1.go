package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func main() {

	//图片大小
	const size = 300 //声明一个size常量，值为300
	//根据给定大小穿件灰度图
	pic := image.NewGray(image.Rect(0, 0, size, size))
	//使用image包的NewGray（）函数创建一个图片对象，使用区域由image.Rect结构提供。
	//image.Rect描述一个方形的两个定位点(x1, y1)和(x2, y2)。
	//image.Rect(0, 0, size, size)表示使用完整灰度图像素，尺寸为宽300，长300

	//遍历每个像素
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {

			//将每个像素的灰度设为255，也就是白色
			pic.SetGray(x, y, color.Gray{255})
		}
	}

	//从0到最大像素生成x坐标
	for x := 0; x < size; x++ {

		//让sin的值得范围在0~2Pi之间
		s := float64(x) * 2 * math.Pi / size //计算math.Sin的定义域
		//rate := x / size
		//s := rate * 2 * math.Pi

		//sin的幅度为一半的像素。想下偏移一半像素并翻转
		y := size/2 - math.Sin(s)*size/2

		//用黑色绘制sin轨迹
		pic.SetGray(x, int(y), color.Gray{0})
	}

	//创建文件
	file, err := os.Create("sin.png")

	if err != nil {
		log.Fatal(err)
	}
	//使用png格式将数据写入文件
	png.Encode(file, pic) //将image信息写入文件中

	//关闭文件
	file.Close()
}

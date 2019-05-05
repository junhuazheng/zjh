package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

//爬取网页内容
func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err1 = err
		return
	}

	defer resp.Body.Close()

	//读取网页body
	buf := make([]byte, 1024*4)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 { //读取结束，或者，出问题
			fmt.Println("resp.Body.Read err = ", err)
			break
		}
		result += string(buf[:n])
	}
	return
}

func DoWork(start, end int) {
	fmt.Printf("正在爬取 %d 到 %d 的页面\n", start, end)

	//目标范围
	//https://tieba.baidu.com/f?kw=%E4%B8%96%E7%95%8Crpg&ie=utf-8&pn=0 下一页加50
	for i := start; i <= end; i++ {
		url := "https://tieba.baidu.com/f?kw=%E4%B8%96%E7%95%8Crpg&ie=utf-8&pn=0" + strconv.Itoa((i-1)*50)
		fmt.Println("url = ", url)

		//爬，将所有的网站的内容全部趴下来
		result, err := HttpGet(url)
		if err != nil {
			fmt.Println("HttpGet err = ", err)
			continue
		}

		//把内容写入到文件
		fileName := strconv.Itoa(i) + "html"
		f, err1 := os.Create(fileName) //创建文件
		if err1 != nil {
			fmt.Println("os.Create err1 = ", err1)
			continue
		}

		//把内容写入到文件
		f.WriteString(result)

		f.Close()

	}
}

func main() {
	var start, end int
	fmt.Printf("请输入起始页 : ")
	fmt.Scan(&start)
	fmt.Printf("请输入终止页 : ")
	fmt.Scan(&end)

	DoWork(start, end)

}

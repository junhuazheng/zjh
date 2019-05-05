/*
爬虫的4个主要步骤:
1、明确目标(范围及准备在哪个网站去搜索)
2、爬(将所有的网站内容全部爬下来)
3、取(去掉我们不需要的数据)
4、处理数据(按照我们想要的方式存储和使用)

网页规律：
https://www.pengfue.com/xiaohua_1.html 下一页 +1

主页面规律:
<h1 class="dp-b"><a href=" 一个段子url链接 "

段子URL规律:
<h1> 标题 <h1>  只取一个
<div class="content-txt pt10"> 段子内容 <a id="prev" href="
*/
package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func HttpGet(url string) (result string, err error) {

	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}

	defer resp.Body.Close()

	//读取网页内容
	buf := make([]byte, 1024*4)
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 { //读取结束或出现错误
			break
		}
		result += string(buf[:n])
	}
	return
}

func SpiderPage(i int) {
	//明确目标
	//https://www.pengfue.com/index_1.html
	url := "https://www.pengfue.com/xiaohua_" + strconv.Itoa(i) + ".html"
	fmt.Printf("正在爬取第%d网页； %s\n", i, url)

	//开始爬页面内容
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet err = ", err)
		return
	}

	//取，<h1 class="dp-b"><a href=" 一个段子url链接 "
	//解析表达式
	re := regexp.MustCompile(`<h1 class="dp-b"><a href="(?s:(.*?))"`)
	if re == nil {
		fmt.Println("regexp.MustCompile err")
		return
	}
	//取关键信息
	joyUrls := re.FindAllStringSubmatch(result, -1) //-1代表取所有内容

	//取网址
	//第一个返回下标，第二个返回内容
	for _, data := range joyUrls {
		//fmt.Println("url = ", data[1])
		//开始爬取每一个笑话，每一个段子
		title, content, err := SpiderOneJoy(data[1])
		if err != nil {
			fmt.Println("SpiderOneJoy err = ", err)
			continue
		}
		fmt.Println("title = ", title)
		fmt.Println("content = ", content)
	}
}

//开始爬取每一个笑话，每一个段子
func SpiderOneJoy(url string) (title, content string, err error) {
	//开始爬取页面内容
	result, err1 := HttpGet(url)
	if err1 != nil {
		err = err1
		return
	}

	//取关键信息
	//取标题
	re1 := regexp.MustCompile(`<h1>(?s:(.*?))<h1>`)
	if re1 == nil {
		err = fmt.Errorf("%s", "regexp.MustCompile err")
		return
	}
	tmpTitle := re1.FindAllStringSubmatch(result, 1) //最后一个参数为1，只过滤第一个
	for _, data := range tmpTitle {
		title = data[1]
		// title = strings.Replace(title, "\n", "", -1) 替换
		// title = strings.Replace(title, "\r", "", -1)
		// title = strings.Replace(title, " ", "", -1)
		title = strings.Replace(title, "\t", "", -1)
		break
	}

	//取内容
	re2 := regexp.MustCompile(`div class="content-txt pt10">(?s:(.*?))<a id="prev" href="`)
	if re2 == nil {
		err = fmt.Errorf("%s", "regexp.MustCompile err")
		return
	}
	tmpContent := re2.FindAllStringSubmatch(result, -1) //只有一个用1和-1都可以
	for _, data := range tmpContent {
		content = data[1]
		content = strings.Replace(content, "\t", "", -1)
		content = strings.Replace(content, "\t", "", -1)
		content = strings.Replace(content, "\t", "", -1)
		// content = strings.Replace(content, "\t", "", -1)
		break
	}

	return
}

func DoWork(start, end int) {

	for i := start; i <= end; i++ {
		//爬取主页面
		SpiderPage(i)
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

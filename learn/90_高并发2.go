package main

import (
	"sync"
)

/*
上一个例子，一个简单的Web服务就搭建起来了。
接下来我们一步一步理解这个Web服务是怎么运行的，怎么做到高并发的。
我们顺着http.HandleFunc(".", response)方法顺着代码一直往上看。
*/

func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	DefaultServeMux.HanndleFunc(pattern, handler)
}

var DefaultServeMux = &defaultServeMux
var defaultServeMux ServeMux

type ServeMux struct {
	mu sync.RWMutex //读写锁。并发处理需要的锁
	m map[string]muxEntry //路由规则map。一个规则一个muxEntry
	hosts bool //规则中是否带有host信息
}

一个路由规则字符串，对应一个handler处理方法
type muxEntry struct {
	h Handler
	pattern string
}

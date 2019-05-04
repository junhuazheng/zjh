package mux

import (
	"flag"
	"net/http"
)

type muxHandler struct {
	handlers    map[string]http.Handler
	handleFuncs map[string]func(resp http.ResponseWriter, req *http.Request)
}

func NewMuxHandler() *muxHandler {
	return &muxHandler{
		make(map[string]http.Handler),
		make(map[string]func(resp http.ResponseWriter, req *http.Request)),
	}
}

func (handler *muxHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	urlPath := req.URL.Path
	if h1, ok := handler.handlers[urlPath]; ok {
		h1.ServeHTTP(resp, req)
		return
	}

	if fn, ok := handler.handleFuncs[urlPath]; ok {
		fn(resp, req)
		return
	}

	http.NotFound(resp, req)
}

func (hander *muxHandler) Handle(pattern string, h1 http.Handler) {
	hander.handlers[pattern] = h1
}

func (handler *muxHandler) handleFunc(pattern string, fn func(resp http.ResponseWriter, req *http.Request)) {
	handler.handleFuncs[pattern] = fn
}

var port string

func main() {
	flag.StringVar(&port, "port", ":8080", "port to listen")
	flag.Parse()

	router := mux.NewMuxHandler()
	router.Handle("/hello/golang/", &BaseHander{})

	router.HandleFunc("/hello/world", func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte("hello world"))
	})
	http.ListenAndServe(port, router)
}

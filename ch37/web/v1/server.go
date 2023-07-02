package v1

import (
	"net"
	"net/http"
)

type HandleFunc func(ctx Context)
type Server interface {
	http.Handler
	Start(addr string) error

	AddRoute(method string, path string, handleFunc HandleFunc)
}

type HTTPServer struct {
}

func (h *HTTPServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx := &Context{
		Req:  request,
		Resp: writer,
	}

	h.serve(ctx)
}

func (h *HTTPServer) Start(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	// 在这里，可以让用户注册所谓的 after start 回调
	// 比如说往你的 admin 注册一下自己这个实例
	// 在这里执行一些你业务所需的前置条件
	return http.Serve(l, h)
}

func (h *HTTPServer) AddRoute(method string, path string, handleFunc HandleFunc) {
	//TODO implement me
	panic("implement me")
}

var _ Server = &HTTPServer{}

func (h *HTTPServer) serve(ctx *Context) {

}

func (h *HTTPServer) Start1(addr string) error {
	return http.ListenAndServe(addr, h)
}

func (h *HTTPServer) Get(path string, handleFunc HandleFunc) {
	h.AddRoute(http.MethodGet, path, handleFunc)
}

func (h *HTTPServer) Post(path string, handleFunc HandleFunc) {
	h.AddRoute(http.MethodGet, path, handleFunc)
}

func (h *HTTPServer) Options(path string, handleFunc HandleFunc) {
	h.AddRoute(http.MethodGet, path, handleFunc)
}

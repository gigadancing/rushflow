package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// 中间件控制流量
type middleWareHandler struct {
	r *httprouter.Router
	l *ConnLimiter
}

// 实现Handler接口
func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !m.l.GetConn() {
		sendErrorResponse(w, http.StatusTooManyRequests, "Too many requests")
		return
	}
	m.r.ServeHTTP(w, r)
	defer m.l.ReleaseConn()
}

//
func NewMiddleWareHandler(r *httprouter.Router, cc int) http.Handler {
	m := middleWareHandler{}
	m.r = r
	m.l = NewConnLimiter(cc)
	return m
}

//
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/videos/:vid-id", streamHandler)
	router.POST("/upload/:vid-id", uploadHandler)
	router.GET("/testpage", testPageHandler)
	return router
}

func main() {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r, 2)
	http.ListenAndServe(":9000", mh)
}

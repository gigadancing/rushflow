package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

//
type middleWareHandler struct {
	r *httprouter.Router
}

//
func NewMiddleWareHandaler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

//
func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 检查session
	validateUserSession(r)
	m.r.ServeHTTP(w, r)
}

//
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)
	return router
}

func main() {
	r := RegisterHandlers()
	mh := NewMiddleWareHandaler(r)
	http.ListenAndServe(":8000", mh)
}

/* handler -> validation{1.request, 2.user} -> business logic ->response
 * data model
 * error handling
 */
// main -> middleware -> defs(message,err) -> handlers -> dbops -> response

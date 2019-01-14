package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"rushflow/api/session"
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
	router.GET("/user/:username", GetUserInfo)
	router.POST("/user/:username/videos", AddNewVideo)
	router.GET("/user/:username/videos", ListAllVideos)
	router.DELETE("/user/:username/videos/:vid-id", DeleteVideo)
	router.POST("/videos/:vid-id/comments", PostComment)
	router.GET("/videos/:vid-id/comments", ShowComments)
	return router
}

//
func Prepare() {
	session.LoadSessionFromDB()
}

func main() {
	Prepare()
	r := RegisterHandlers()
	mh := NewMiddleWareHandaler(r)
	http.ListenAndServe(":8000", mh)
}

/* handler -> validation{1.request, 2.user} -> business logic ->response
 * data model
 * error handling
 */
// main -> middleware -> defs(message,err) -> handlers -> dbops -> response

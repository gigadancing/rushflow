package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

//
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/videos/:vid-id", streamHandler)
	router.POST("/upload/:vid-id", uploadHandler)
	return router
}

func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":9000", r)
}

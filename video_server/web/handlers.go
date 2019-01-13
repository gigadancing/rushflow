package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
)

//
type HomePage struct {
	Name string
}

//
type UserPage struct {
	Name string
}

//
func homeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		t   *template.Template
		err error
	)
	p := &HomePage{
		Name: "xiaodianying",
	}
	if t, err = template.ParseFiles("./templates/home.html"); err != nil {
		log.Printf("Parsing templates home.html error:%v", err)
		return
	}
	//
	t.Execute(w, p)
	return
}

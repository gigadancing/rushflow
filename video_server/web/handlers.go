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

	// visitor
	// 取cookie
	cname, err1 := r.Cookie("username")
	sid, err2 := r.Cookie("session")

	// cookie中没有登录信息
	if err1 != nil || err2 != nil {
		p := &HomePage{
			Name: "chendada",
		}
		if t, err = template.ParseFiles("./templates/home.html"); err != nil {
			log.Printf("Parsing home.html error:%v", err)
			return
		}
		t.Execute(w, p)
		return
	}

	// user
	// cookie中有登录信息
	if len(cname.Value) != 0 && len(sid.Value) != 0 {
		http.Redirect(w, r, "/userhome", http.StatusFound)
		return
	}

}

//
func userHomeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		p   *UserPage
		t   *template.Template
		err error
	)

	cname, err1 := r.Cookie("username")
	_, err2 := r.Cookie("session")
	//
	if err1 != nil || err2 != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	//
	fname := r.FormValue("username")
	if len(cname.Value) != 0 {
		p = &UserPage{
			Name: cname.Value,
		}
	} else if len(fname) == 0 {
		p = &UserPage{
			Name: fname,
		}
	}

	if t, err = template.ParseFiles("./templates/userhome.html"); err != nil {
		log.Printf("Parse userhome.html error:%v", err)
		return
	}

	t.Execute(w, p)
}

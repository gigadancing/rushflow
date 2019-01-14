package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
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

//
func apiHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method != http.MethodPost {
		re, _ := json.Marshal(ErrorRequestNotRecognized)
		io.WriteString(w, string(re))
		return
	}

	res, _ := ioutil.ReadAll(r.Body)
	apiBody := &ApiBody{}
	if err := json.Unmarshal(res, apiBody); err != nil {
		re, _ := json.Marshal(ErrorRequestBodyParseFailed)
		io.WriteString(w, string(re))
		return
	}

	request(apiBody, w, r)
	defer r.Body.Close()

}

//
func proxyHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	u, _ := url.Parse("http://127.0.0.1:9000/")
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ServeHTTP(w, r)
}

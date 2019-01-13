package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	client *http.Client
)

func init() {
	client = &http.Client{}
}

//
func request(b *ApiBody, w http.ResponseWriter, r *http.Request) {
	var (
		resp *http.Response
		err  error
	)

	switch b.Method {
	case http.MethodGet:
		req, _ := http.NewRequest("GET", b.Url, nil)
		req.Header = r.Header
		if resp, err = client.Do(req); err != nil {
			log.Printf(err.Error())
			return
		}
		normalResponse(w, resp)

	case http.MethodPost:
		req, _ := http.NewRequest("POST", b.Url, bytes.NewBuffer([]byte(b.ReqBody)))
		req.Header = r.Header

		if resp, err = client.Do(req); err != nil {
			log.Printf(err.Error())
			return
		}
		normalResponse(w, resp)

	case http.MethodDelete:
		req, _ := http.NewRequest("Delete", b.Url, nil)
		req.Header = r.Header

		if resp, err = client.Do(req); err != nil {
			log.Printf(err.Error())
			return
		}
		normalResponse(w, resp)

	default:
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Bad api request")
		return
	}
}

//
func normalResponse(w http.ResponseWriter, r *http.Response) {
	var (
		res []byte
		err error
	)

	if res, err = ioutil.ReadAll(r.Body); err != nil {
		re, _ := json.Marshal(ErrorInternalFaults)
		w.WriteHeader(500)
		io.WriteString(w, string(re))
		return
	}

	w.WriteHeader(r.StatusCode)
	io.WriteString(w, string(res))
}

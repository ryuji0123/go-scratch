package main

import (
	"go-scratch/ofmt"
	"go-scratch/onet/ohttp"
	"net/http"
)


type OHelloHandler struct {

}

func (oh *OHelloHandler) ServeHTTP(w ohttp.ResponseWriter, r *http.Request) {
	ofmt.Fprintf(w, "Hello!")
}

func init() {
	oh := OHelloHandler{}
	ohttp.Handle("/", &oh)
	ohttp.Handle("/index", &oh)
	oserver := ohttp.Server{
		Addr: "127.0.0.1:8800",
	}
	oserver.ListenAndServe()
}

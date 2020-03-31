package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct {

}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}


func main() {
	hello:= HelloHandler{}
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.Handle("/", &hello)
	server.ListenAndServe()
}

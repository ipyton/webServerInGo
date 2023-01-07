package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct {
}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//如果自己写必须要实现该方法
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "helloworld", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)     //自动转换，带有f
	http.ListenAndServe(":8080", nil) //default router
	myHandler := MyHandler{}

	http.Handle("/myHandler", &myHandler)
	server := http.Server{
		Addr:        ":8080",
		Handler:     &myHandler,
		ReadTimeout: 20,
	}
	server.ListenAndServe()

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	http.ListenAndServe("8080", mux)

}

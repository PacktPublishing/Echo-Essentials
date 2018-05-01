package main

import "net/http"

func main() {
	http.Handle("/", new(myHandler))
	http.ListenAndServe(":8080", nil)
}

type myHandler struct{}

func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("hello!"))
}

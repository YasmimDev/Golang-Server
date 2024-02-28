package main

import "net/http"

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", HandleFunc)
	http.ListenAndServe(":8080", mux)

}

func HandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))

}

package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	// リクエストを解析
	// コンテンツタイプにより使用するメソッドが異なる
	// multipart/form-data の場合はParseMultipartForm/MultipartFormを用いる

	// r.ParseForm()
	// r.ParseMultipartForm(1024)

	// fmt.Fprintln(w, r.Form)
	// fmt.Fprintln(w, r.PostForm)
	// fmt.Fprintln(w, r.MultipartForm)
	// fmt.Fprintln(w, r.FormValue("hello"))
	// fmt.Fprintln(w, r.PostFormValue("hello"))
	fmt.Fprintln(w, "(1)", r.FormValue("hello"))
	fmt.Fprintln(w, "(2)", r.PostFormValue("hello"))
	fmt.Fprintln(w, "(3)", r.PostForm)
	fmt.Fprintln(w, "(4)", r.MultipartForm)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}

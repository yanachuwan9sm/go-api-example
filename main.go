package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}
	postArticleHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Article...\n")
	}
	articleListHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Article List\n")
	}
	articleDetailHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Article No.1\n")
	}
	postNiceHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Nice...\n")
	}
	postCommentHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Comment...\n")
	}

	// 定義した helloHandlerを使うように登録
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/article", postArticleHandler)
	http.HandleFunc("/article/list", articleListHandler)
	http.HandleFunc("/article/1", articleDetailHandler)
	http.HandleFunc("/article/nice", postNiceHandler)
	http.HandleFunc("/comment", postCommentHandler)

	// サーバー起動時のログを出力
	log.Println("server start at port 8080")
	//  ListenAndServe関数にて、サーバーを起動
	log.Fatal(http.ListenAndServe(":8080", nil))
}

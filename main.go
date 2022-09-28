package main

import (
	"log"
	"net/http"

	"github.com/yanachuwan9sm/myapi-tutorial/handlers"
)

func main() {

	// 定義した helloHandlerを使うように登録
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/article", handlers.PostArticleHandler)
	http.HandleFunc("/article/list", handlers.ArticleListHandler)
	http.HandleFunc("/article/1", handlers.ArticleDetailHandler)
	http.HandleFunc("/article/nice", handlers.PostNiceHandler)
	http.HandleFunc("/comment", handlers.PostCommentHandler)

	// サーバー起動時のログを出力
	log.Println("server start at port 8080")
	//  ListenAndServe関数にて、サーバーを起動
	log.Fatal(http.ListenAndServe(":8080", nil))
}

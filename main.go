package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yanachuwan9sm/myapi-tutorial/handlers"
)

func main() {

	//? 標準パッケージ　net/http のみでルーティングを実装

	// 定義した helloHandlerを使うように登録
	// http.HandleFunc("/hello", handlers.HelloHandler)
	// http.HandleFunc("/article", handlers.PostArticleHandler)
	// http.HandleFunc("/article/list", handlers.ArticleListHandler)
	// http.HandleFunc("/article/1", handlers.ArticleDetailHandler)
	// http.HandleFunc("/article/nice", handlers.PostNiceHandler)
	// http.HandleFunc("/comment", handlers.PostCommentHandler)

	//? gorilla/mux(サードパーティー) パッケージでルーティングを実装

	// ルータの作成
	r := mux.NewRouter()
	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/1", handlers.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)

	// サーバー起動時のログを出力
	log.Println("server start at port 8080")
	//  ListenAndServe関数にて、サーバーを起動
	//* 第二引数は、サーバーの中で使うルータを指定するもの。標準パッケージ net/http では
	//* ルータが渡されず nil だったのは、
	//* Go の HTTP サーバーがデフォルトで持っているルータが自動的に採用されているため。
	// log.Fatal(http.ListenAndServe(":8080", nil))
	log.Fatal(http.ListenAndServe(":8080", r))
}

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/yanachuwan9sm/myapi-tutorial/controllers"
	"github.com/yanachuwan9sm/myapi-tutorial/services"
)

// var (
// 	dbUser     = os.Getenv("DB_USER")
// 	dbPassword = os.Getenv("DB_PASSWORD")
// 	dbDatabase = os.Getenv("DB_NAME")
// 	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
// )

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		// .env読めなかった場合の処理
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}
	dbUser := os.Getenv("USERNAME")
	dbPassword := os.Getenv("USERPASS")
	dbDatabase := os.Getenv("DATABASE")
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	// サーバー全体で使用する sql.DB 型を生成
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("fail to connect DB")
		return
	}
	// sql.DB型をもとに、サーバー全体で使用するサービス構造体MyAppServiceを生成
	ser := services.NewMyAppService(db)

	// MyAppService型をもとに、サーバー全体で使用するコントローラ構造体MyAppControllerを生成
	con := controllers.NewMyAppController(ser)

	// コントローラ型 MyAppController のハンドラメソッドとパスとの関連付けを行う
	r := mux.NewRouter() // ルータの作成
	r.HandleFunc("/article", con.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", con.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", con.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", con.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", con.PostCommentHandler).Methods(http.MethodPost)

	// サーバー起動時のログを出力
	log.Println("server start at port 8080")

	//  ListenAndServe関数にて、サーバーを起動
	//* 第二引数は、サーバーの中で使うルータを指定するもの。標準パッケージ net/http では
	//* ルータが渡されず nil だったのは、
	//* Go の HTTP サーバーがデフォルトで持っているルータが自動的に採用されているため。
	// log.Fatal(http.ListenAndServe(":8080", nil))
	log.Fatal(http.ListenAndServe(":8080", r))
}

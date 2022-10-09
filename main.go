package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/yanachuwan9sm/myapi-tutorial/api"
)

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
	// ser := services.NewMyAppService(db)

	// サーバー全体で使用するコントローラ構造体MyAppControllerを生成
	// インターフェースへの書き換えに伴い、
	// NewMyAppController 関数の引数が MyAppService 型から MyAppServicer インターフェース型に変更
	// しかし、MyAppService 型である変数 ser は、MyAppServicer インターフェースにそのまま代入可能なので、
	// 記述の変更は必要ない。

	// サービス構造体 MyAppService(変数 ser) をもとに、
	// ArticleController(変数 aCon) と CommentController(変数 cCon) を作成
	// aCon := controllers.NewArticleController(ser)
	// cCon := controllers.NewCommentController(ser)

	// 2 つのコントローラ構造体から、gorilla/mux のルータを作成
	r := api.NewRouter(db)

	// サーバー起動時のログを出力
	log.Println("server start at port 8080")

	//  ListenAndServe関数にて、サーバーを起動
	//* 第二引数は、サーバーの中で使うルータを指定するもの。標準パッケージ net/http では
	//* ルータが渡されず nil だったのは、
	//* Go の HTTP サーバーがデフォルトで持っているルータが自動的に採用されているため。
	// log.Fatal(http.ListenAndServe(":8080", nil))
	log.Fatal(http.ListenAndServe(":8080", r))
}

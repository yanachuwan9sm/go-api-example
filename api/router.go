package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yanachuwan9sm/myapi-tutorial/controllers"
	"github.com/yanachuwan9sm/myapi-tutorial/services"
)

// パスーハンドラ関数の対応付けがなされた gorilla/mux ルータを 作成し、戻り値として返却する関数
func NewRouter(db *sql.DB) *mux.Router {

	// sql.DB型をもとに、サーバー全体で使用するサービス構造体MyAppServiceを生成
	ser := services.NewMyAppService(db)

	// サービス構造体 MyAppService(変数 ser) をもとに、
	// ArticleController(変数 aCon) と CommentController(変数 cCon) を作成
	aCon := controllers.NewArticleController(ser)
	cCon := controllers.NewCommentController(ser)

	r := mux.NewRouter()
	r.HandleFunc("/article", aCon.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", aCon.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", aCon.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", aCon.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", cCon.PostCommentHandler).Methods(http.MethodPost)
	return r
}

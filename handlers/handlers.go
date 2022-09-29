package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GET /hello のハンドラ
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	// 第二引数 req に、受け取った HTTP リクエストの情報が入っている

	// req.Method が GET のときのみ正常応答を返す
	// if req.Method == http.MethodGet {
	io.WriteString(w, "Hello, world!\n")
	// } else {
	// Invalid methodというレスポンスを、405番のステータスコードと共に返す
	// 	http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	// }
}

// POST /article のハンドラ
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Article...\n")
}

// GET /article/list のハンドラ
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {

	queryMap := req.URL.Query()

	var page int

	if p, ok := queryMap["page"]; ok && len(p) > 0 {

		// パラメータ page に対応する 1 つ目の値を採用し、それを数値に変換する
		var err error
		page, err = strconv.Atoi(p[0]) // パラメータ page に対応する 1 つ目の値を採用し、それを数値に変換する

		if err != nil {
			// 数字に変換できなかった場合には、リクエスト不正(400)のエラーを返す
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}

	} else {
		// クエリパラメータがURLに存在しない場合
		page = 1
	}

	resString := fmt.Sprintf("Article List (page %d)\n", page)
	io.WriteString(w, resString)
}

// GET /article/{id} のハンドラ
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		// パスパラメータが不明な場合、リクエスト不正のエラー(400)を返却
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
	resString := fmt.Sprintf("Article No . %d\n", articleID)
	io.WriteString(w, resString)
}

// POST /article/nice のハンドラ
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice...\n")

}

// POST /comment のハンドラ
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Article...\n")
}

package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yanachuwan9sm/myapi-tutorial/models"
	"github.com/yanachuwan9sm/myapi-tutorial/services"
)

// コントローラ構造体を定義
type MyAppController struct {
	// フィールドに MyAppService構造体を含める
	service *services.MyAppService
}

// コンストラクタの定義
func NewMyAppController(s *services.MyAppService) *MyAppController {
	return &MyAppController{service: s}
}

// POST /article のハンドラ
func (c *MyAppController) PostArticleHandler(w http.ResponseWriter, req *http.Request) {

	var reqArticle models.Article

	// ストリームから直接リクエストデータを取るようにしたことで、
	// デコード前の「Content-Length ヘッダフィールドの値からバイトスライスを作り、
	// そこにリクエストボディの中身を書き込む」と いう操作がまるまるいらない。
	// (直接デコーダの Decode メソッドを呼び出すだけで済むため)
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}

	article, err := c.service.PostArticleService(reqArticle)

	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)

}

// GET /article/list のハンドラ
func (c *MyAppController) ArticleListHandler(w http.ResponseWriter, req *http.Request) {

	queryMap := req.URL.Query()

	// クエリパラメータ pager を取得
	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		// パラメータ page に対応する 1 つ目の値を採用し、それを数値に変換する
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			// 数字に変換できなかった場合には、リクエスト不正(400)のエラーを返す
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		// クエリパラメータがURLに存在しない場合
		page = 1
	}

	articleList, err := c.service.GetArticleListService(page)

	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(articleList)
}

// GET /article/{id} のハンドラ
func (c *MyAppController) ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {

	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		// パスパラメータが不明な場合、リクエスト不正のエラー(400)を返却
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	article, err := c.service.GetArticleService(articleID)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}

// POST /article/nice のハンドラ
func (c *MyAppController) PostNiceHandler(w http.ResponseWriter, req *http.Request) {

	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	article, err := c.service.PostNiceService(reqArticle)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}

// POST /comment のハンドラ
func (c *MyAppController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {

	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(comment)
}

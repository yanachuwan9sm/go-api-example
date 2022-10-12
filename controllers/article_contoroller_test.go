package controllers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestArticleListHandler(t *testing.T) {
	var tests = []struct {
		name       string
		query      string
		resultCode int
	}{
		{name: "number query", query: "1", resultCode: http.StatusOK},
		{name: "alphabet query", query: "aaa", resultCode: http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// httptest.NewRequest関数でリクエスト作成
			url := fmt.Sprintf("http://localhost:8080/article/list?page=%s", tt.query)
			req := httptest.NewRequest(http.MethodGet, url, nil)

			// httptest.ResponseRecorder構造体を用意
			res := httptest.NewRecorder()

			// ハンドラメソッドを実行
			aCon.ArticleListHandler(res, req)

			if res.Code != tt.resultCode {
				t.Errorf("unexpected StatusCode: want %d but %d\n", tt.resultCode, res.Code)
			}
		})
	}
}

func TestArticleDetailHandler(t *testing.T) {

	// テストケースを用意
	var tests = []struct {
		name       string
		articleID  string
		resultCode int
	}{
		{name: "number pathparam", articleID: "1", resultCode: http.StatusOK},
		{name: "alphabet pathparam", articleID: "aaa", resultCode: http.StatusNotFound},
	}

	// テーブルドリブンに実行
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// httptest.NewRequest関数でリクエスト作成
			url := fmt.Sprintf("http://localhost:8080/article/%s", tt.articleID)
			req := httptest.NewRequest(http.MethodGet, url, nil)

			// httptest.ResponseRecorder構造体を用意
			res := httptest.NewRecorder()

			// テスト対象であるハンドラメソッド ArticleListHandler では
			// gorilla/muxパッケージ内で定義される mux.Vars関数でパスパラメータidを取得
			// -> mux.Vars 関数によるパスパタメータの取得は、gorilla/muxルータ経由で受け取ったリクエストでしかうまく動作しない
			// aCon.ArticleListHandler(res, req)

			// 公式ドキュの通りに書く

			r := mux.NewRouter()
			// テストで使うパスとハンドラの対応関係をルータに登録
			r.HandleFunc("/article/{id:[0-9]+}",
				aCon.ArticleDetailHandler).Methods(http.MethodGet)
			// ルータ r 経由でテスト用リクエストを送信
			r.ServeHTTP(res, req)

			if res.Code != tt.resultCode {
				t.Errorf("unexpected StatusCode: want %d but %d\n", tt.resultCode, res.Code)
			}
		})
	}
}

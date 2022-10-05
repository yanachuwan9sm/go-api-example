package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yanachuwan9sm/myapi-tutorial/models"
	"github.com/yanachuwan9sm/myapi-tutorial/services"
)

// GET /hello のハンドラ
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

// POST /article のハンドラ
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {

	// //? リクエストボディの内容を格納するバイト列を作成

	// // リクエストヘッダの Content-Length フィールドの値を取得
	// length, err := strconv.Atoi(req.Header.Get("Content-Length"))

	// if err != nil {
	// 	// getメソッドの返却値(string)からint型への変換が失敗した場合 400 番エラー (BadRequest) を返却
	// 	http.Error(w, "cannnot get content length\n", http.StatusBadRequest)
	// 	return
	// }
	// // make関数を使ってその長さのバイトスライスを作成
	// reqBodybuffer := make([]byte, length)

	// //? Readメソッドでリクエストボディを読み出し

	// //* _, err := req.Body.Read(reqBodybuffer);
	// // 変数 reqBodybuffer にリクエストボディの内容が入る
	// // 戻り値 err に、読み取り時に起きたエラーの内容が格納される

	// //* Read メソッドから返ってきたエラーの種類によって、正常にボディの読み込みが終わったのかそうでないのかが分かれます。
	// // 戻り値で返ってきたエラーが io.EOF かどうかを判定するために、標準パッケージ errors の中にある errors.Is 関数を利用
	// // エラーが io.EOF だった場合: 正常にボディの中身を最後まで読み取ることができた
	// // エラーが io.EOF 以外だった場合: ボディの中身を読み取る際に異常が発生した

	// if _, err := req.Body.Read(reqBodybuffer); !errors.Is(err, io.EOF) {
	// 	// エラーが io.EOF以外の場合、サーバー内で異常が起きたことを示すエラー(500)を返却
	// 	http.Error(w, "fail to get request body\n", http.StatusBadRequest)
	// 	return
	// }

	// // ボディを Close する
	// defer req.Body.Close()

	var reqArticle models.Article
	// if err := json.Unmarshal(reqBodybuffer, &reqArticle); err != nil {
	// 	http.Error(w, "fail to docode json\n", http.StatusBadRequest)
	// 	return
	// }

	// ストリームから直接リクエストデータを取るようにしたことで、
	// デコード前の「Content-Length ヘッダフィールドの値からバイトスライスを作り、
	// そこにリクエストボディの中身を書き込む」と いう操作がまるまるいらない。
	// (直接デコーダの Decode メソッドを呼び出すだけで済むため)
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {

		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}

	// デコードした構造体を再度エンコード
	// article := reqArticle
	// jsonData, err := json.Marshal(article)
	// if err != nil {
	// 	http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
	// 	return
	// }
	// w.Write(jsonData)

	article, err := services.PostArticleService(reqArticle)

	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)

}

// GET /article/list のハンドラ
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {

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

	articleList, err := services.GetArticleListService(page)

	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	// jsonData, err := json.Marshal(articleList)
	// if err != nil {
	// 	errMsg := fmt.Sprintf("fail to encode json (page %d)\n", page)
	// 	http.Error(w, errMsg, http.StatusInternalServerError)
	// 	return
	// }

	// w.Write(jsonData)
	json.NewEncoder(w).Encode(articleList)
}

// GET /article/{id} のハンドラ
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {

	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		// パスパラメータが不明な場合、リクエスト不正のエラー(400)を返却
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	article, err := services.GetArticleService(articleID)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}

// POST /article/nice のハンドラ
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {

	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	article, err := services.PostNiceService(reqArticle)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}

// POST /comment のハンドラ
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {

	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	comment, err := services.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(comment)
}

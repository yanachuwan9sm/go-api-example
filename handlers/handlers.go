package handlers

import (
	"fmt"
	"io"
	"net/http"
)

// GET /hello のハンドラ
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	// 第二引数 req に、受け取った HTTP リクエストの情報が入っている

	// req.Method が GET のときのみ正常応答を返す
	if req.Method == http.MethodGet {
		io.WriteString(w, "Hello, world!\n")
	} else {
		// Invalid methodというレスポンスを、405番のステータスコードと共に返す
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// POST /article のハンドラ
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Article...\n")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// GET /article/list のハンドラ
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		io.WriteString(w, "Article List\n")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// GET /article/1 のハンドラ
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		articleID := 1
		resString := fmt.Sprintf("Article No . %d\n", articleID)
		io.WriteString(w, resString)
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// POST /article/nice のハンドラ
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Nice...\n")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// POST /comment のハンドラ
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Article...\n")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

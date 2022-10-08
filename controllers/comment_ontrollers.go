package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/yanachuwan9sm/myapi-tutorial/controllers/services"
	"github.com/yanachuwan9sm/myapi-tutorial/models"
)

// コントローラ構造体を定義
type CommentController struct {
	service services.CommentServicer // Comment用のサービスインターフェース
}

// コンストラクタの定義
func NewCommentController(s services.CommentServicer) *CommentController {
	return &CommentController{service: s}
}

// POST /comment のハンドラ
func (c *CommentController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {

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

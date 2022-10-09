package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/yanachuwan9sm/myapi-tutorial/apperrors"
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
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, req, err)
	}

	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(comment)
}

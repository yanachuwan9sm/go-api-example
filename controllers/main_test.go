package controllers_test

import (
	"testing"

	"github.com/yanachuwan9sm/myapi-tutorial/controllers"
	"github.com/yanachuwan9sm/myapi-tutorial/controllers/testdata"
)

// 1. テストに使うリソース (コントローラ構造体) を用意
var aCon *controllers.ArticleController

func TestMain(m *testing.M) {
	ser := testdata.NewServiceMock()
	aCon = controllers.NewArticleController(ser)

	m.Run()
}

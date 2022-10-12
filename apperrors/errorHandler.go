/*
* 独自エラー MyAppError を受け取り、その内容にしたがってユーザーに
* 適切な HTTP レスポンスを返すようなエラーハンドラ
 */

package apperrors

import (
	"encoding/json"
	"errors"
	"net/http"
)

// エラーが発生したときのレスポンス処理をここで一括で行う
func ErrorHandler(w http.ResponseWriter, req *http.Request, err error) {
	var appErr *MyAppError

	// 受け取ったエラーを独自エラー型に変換
	if !errors.As(err, &appErr) {

		// 第三引数で渡されたエラーが MyAppError 型 でない場合
		// (開発者が想定していなかった不明なエラーが起きた)
		appErr = &MyAppError{
			ErrCode: Unknown,
			Message: "internal process failed",
			Err:     err,
		}
	}

	// ユーザーに返却する HTTPレスポンスコードを格納する変数
	var statusCode int

	switch appErr.ErrCode {
	// 指定の記事が存在しない場合
	case NAData:
		statusCode = http.StatusNotFound // 404(NotFound)
	// コメント投稿先として指定された記事がなかった (NoTargetData) 場合
	// リクエストボディの json デコードに失敗した (ReqBodyDecodeFailed) 場合
	// リクエストパラメータの値が不正だった (BadParam) 場合
	case NoTargetData, ReqBodyDecodeFailed, BadParam:
		statusCode = http.StatusBadRequest // 400(BadRequest)
	default:
		statusCode = http.StatusInternalServerError // 500(InternalServerError)
	}

	// レスポンスの作成
	w.WriteHeader(statusCode)         // リクエストヘッダにステータスコードを書込
	json.NewEncoder(w).Encode(appErr) // リクエストボディにエラー内容をjsonエンコードして書込

}

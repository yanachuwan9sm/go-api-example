package apperrors

type MyAppError struct {
	// ErrCode型のErrCodeフィールド
	// (フィールド名を省略した場合、型名がそのままフィールド名になる)
	ErrCode
	// string型のMessageフィールド
	Message string
	// エラーチェーンのための内部エラー
	Err error `json:"-"` // json:"-"タグを付与することで、このフィールドはjsonエンコードされない
}

func (myErr *MyAppError) Error() string {
	return myErr.Err.Error()
}

// errors.Unwrap 関数によって内部のエラーを取り出せるようにするには、
// その独自エラー構造体に Unwrap メソッドが必要
func (myErr *MyAppError) Unwrap() error {
	return myErr.Err
}

// 元となるエラーを受け取り、MyAppError型にラップして返すラップメソッド
func (code ErrCode) Wrap(err error, message string) error {
	// メソッドのレシーバーとなったエラーコードを含んだ MyAppError構造体を作る
	return &MyAppError{ErrCode: code, Message: message, Err: err}
}

package middlewares

import (
	"context"
	"sync"
)

var (
	logNo int = 1
	mu    sync.Mutex
)

// コンテキストにトレース ID を付加して返す関数
func SetTraceID(ctx context.Context, traceID int) context.Context {
	// ctx に、(key: "traceID", value: 変数 traceID の値) をセット
	return context.WithValue(ctx, "traceID", traceID)
}

// コンテキストに含まれているトレース ID を取り出す関数
func GetTraceID(ctx context.Context) int {
	// キー"traceID"に紐づく値をコンテキストから取り出す
	id := ctx.Value("traceID")
	// 変数 id は any 型なので、int 型にアサーションする
	if idInt, ok := id.(int); ok {
		return idInt
	}

	return 0
}

func newTraceID() int {
	var no int

	mu.Lock()
	no = logNo
	logNo += 1
	mu.Unlock()

	return no
}

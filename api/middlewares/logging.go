package middlewares

import (
	"log"
	"net/http"
)

// レスポンスのロギングする機能を持つ ResponseWriter 構造体 (resLoggingWriter 構造体)
// 委譲によって、Header メソッド・Write メソッド・WriteHeader メソッドを持つ
// -> resLoggingWriter構造体自体が ResponseWriterインターフェースを満たす
type resLoggingWriter struct {
	// 元々使用していた http.ResponseWriter を格納するためのフィールド
	http.ResponseWriter
	// ハンドラが使ったレスポンスコードを格納しておくためのフィールド
	code int
}

// コンストラクタを作成
func NewResLoggingWriter(w http.ResponseWriter) *resLoggingWriter {
	return &resLoggingWriter{ResponseWriter: w, code: http.StatusOK}
}

// ハンドラが HTTP レスポンスコードを書き込むときに使うメソッド
// memo :
// resLoggingWriter 構造体は委譲によって既に定義されたWriteHeader メソッド
// と同じメソッドをその構造体にも定義し、メソッドの処理内容を上書き(オーバーライド)
func (rsw *resLoggingWriter) WriteHeader(code int) {

	// resLoggingWriter構造体のcodeフィールドに、使うレスポンスコードを保存する
	rsw.code = code
	// HTTPレスポンスに使うレスポンスコードを指定
	// (=WriteHeaderメソッド本来の機能を呼び出し)
	rsw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		traceID := newTraceID()

		// リクエスト情報をロギング
		log.Printf("[%d]%s %s\n", traceID, req.RequestURI, req.Method)

		// TraceIDを付加したコンテキストを取得
		ctx := SetTraceID(req.Context(), traceID)
		// Request型にコンテキストをセットする
		req = req.WithContext(ctx)

		rlw := NewResLoggingWriter(w)

		// 元のハンドラを実行
		next.ServeHTTP(rlw, req)

		// レスポンス情報をロギング
		log.Printf("[%d]res: %d", traceID, rlw.code)
	})
}

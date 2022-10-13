package services_test

import "testing"

// GetArticleService メソッドの処理時間を計測するためのベンチマークテスト
// ベンチマークの取り方 -> 処理時間を測りたい関数・メソッドを複数回実行し、その平均を求める方法
func BenchmarkGetArticleService(b *testing.B) {

	//前処理
	articleID := 1
	// 前処理+テスト対象の実行時間が計測されてしまうため、タイマーをリセット
	b.ResetTimer()

	// 処理時間を測りたい関数・メソッドを複数回実行し、その平均を求める
	// 繰り返す回数は *testing.B構造体が持つ b.N により、よしなに決定される。
	for i := 0; i < b.N; i++ {
		_, err := aSer.GetArticleService(articleID)
		if err != nil {
			b.Error(err)
			break
		}
	}
}

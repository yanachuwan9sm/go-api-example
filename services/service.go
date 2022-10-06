package services

import "database/sql"

// 1. サービス構造体を定義
type MyAppService struct {
	// 2. フィールドに sql.DB 型を含める
	db *sql.DB
}

// コンストラクタの定義
func NewMyAppService(db *sql.DB) *MyAppService {
	return &MyAppService{db: db}
}

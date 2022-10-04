package repositories_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// テスト全体で共有する sql.DB 型
var testDB *sql.DB

// 全テスト共通の前処理を書く
func setup() error {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
	var err error
	// testDBにOpen関数で得たsql.DB型を代入
	testDB, err = sql.Open("mysql", dbConn)
	if err != nil {
		return err
	}
	return nil
}

// 前テスト共通の後処理を書く
func teardown() {
	// teardown関数の中からtestDBは参照可能
	testDB.Close()
}

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		// TestMain 関数の引数となっている testing.M は Fatal 系メソッドを持たないので、
		// os.Exit(1)でテストを FAIL・終了させる
		os.Exit(1)
	}
	m.Run()
	teardown()
}

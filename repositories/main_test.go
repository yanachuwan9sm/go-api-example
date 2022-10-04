package repositories_test

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// DB接続部分
var (
	dbUser     = "docker"
	dbPassword = "docker"
	dbDatabase = "sampledb"
	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

func connectDB() error {
	var err error
	testDB, err = sql.Open("mysql", dbConn)
	if err != nil {
		return err
	}
	return nil
}

func setupTestData() error {

	// mysql command を用いて setupDB.sql を実行
	cmd := exec.Command("mysql", "-h", "127.0.0.1", "-u", "docker", "sampledb",
		"--password=docker", "-e", "source ./testdata/setupDB.sql")
	// Runメソッドでコマンド実行
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func cleanupDB() error {
	// mysql command を用いて cleanupDB.sql を実行
	cmd := exec.Command("mysql", "-h", "127.0.0.1", "-u", "docker", "sampledb",
		"--password=docker", "-e", "source ./testdata/cleanupDB.sql")
	// Runメソッドでコマンド実行
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

// テスト全体で共有する sql.DB 型
var testDB *sql.DB

// 全テスト共通の前処理を書く
func setup() error {

	if err := connectDB(); err != nil {
		return err
	}

	if err := cleanupDB(); err != nil {
		fmt.Println("cleanup", err)
		return err
	}

	if err := setupTestData(); err != nil {
		fmt.Println("setup", err)
		return err
	}

	return nil

}

// 前テスト共通の後処理を書く
func teardown() {
	// teardown関数の中からtestDBは参照可能
	cleanupDB()
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

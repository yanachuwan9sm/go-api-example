package services_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yanachuwan9sm/myapi-tutorial/services"
)

// レシーバーとして使うための MyAppService 構造体
var aSer *services.MyAppService

func TestMain(m *testing.M) {

	// データベースに必要な変数
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	aSer = services.NewMyAppService(db)
	// 個別のベンチマークテストの実行
	m.Run()

}

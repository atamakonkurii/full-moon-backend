package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq"
)

// TestUser : テーブルデータ
type TestUser struct {
    UserID   int
    Password string
}

// メイン関数
func main() {

    // Db: データベースに接続するためのハンドラ
    var Db *sql.DB
    // Dbの初期化
    Db, err := sql.Open("postgres", "host=db port=5432 user=postgres password=password dbname=full_moon_db sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }

	// Dbの接続確認
	err = Db.Ping()
    if err != nil{
        log.Fatal(err)
    }

    // SQL文の構築
    sql := "SELECT user_id, user_password FROM TEST_USER WHERE user_id=$1;"
	// sql := "SELECT * FROM TEST_USER;"

    // preparedstatement の生成
    pstatement, err := Db.Prepare(sql)
    if err != nil {
        log.Fatal(err)
    }

    // 検索パラメータ（ユーザID）
    queryID := 1
    // 検索結果格納用の TestUser
    var testUser TestUser

    // queryID を埋め込み SQL の実行、検索結果1件の取得
    err = pstatement.QueryRow(queryID).Scan(&testUser.UserID, &testUser.Password)
    if err != nil {
        log.Fatal(err)
    }

    // 検索結果の表示
    fmt.Println(testUser.UserID, testUser.Password)
}

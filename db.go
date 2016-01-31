package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB_URL = os.Getenv("INSTAGRAM_DB")
var DB_PASS = os.Getenv("INSTAGRAM_DB_PASS")

func insertFollows(follows []string) {
	connection_url := fmt.Sprintf("root:%v@tcp(%v:3306)/instagram_statistics", DB_PASS, DB_URL)

	db, err := sql.Open("mysql", connection_url)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	stmtIns, err := db.Prepare("INSERT INTO Follows VALUES (null, ?, ?)")
	if err != nil {
		fmt.Println(err)
	}
	defer stmtIns.Close()

	for i := 0; i < len(follows); i++ {
		_, err = stmtIns.Exec(follows[i], time.Now())
	}
}

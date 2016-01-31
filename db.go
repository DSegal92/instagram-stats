package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB_URL = os.Getenv("INSTAGRAM_DB")
var DB_PASS = os.Getenv("INSTAGRAM_DB_PASS")

func insertRelations(relationship string, users []string) {
	connection_url := fmt.Sprintf("root:%v@tcp(%v:3306)/instagram_statistics", DB_PASS, DB_URL)
	updateTime := time.Now()

	db, err := sql.Open("mysql", connection_url)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	db_name := strings.Title(relationship)
	insert_statement := fmt.Sprintf(`INSERT INTO %v (id, username, most_recently_seen, first_seen)
																	VALUES (null, ?, ?, ?)
																	on duplicate key update
																		most_recently_seen = values(most_recently_seen)`, db_name)
	stmtIns, err := db.Prepare(insert_statement)

	if err != nil {
		fmt.Println(err)
	}
	defer stmtIns.Close()

	for i := 0; i < len(users); i++ {
		_, err = stmtIns.Exec(users[i], updateTime, updateTime)
	}
}

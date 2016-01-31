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

func insertRelations(relationship string, users []string, updateTime time.Time) {
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

func updateStatistics(follows int, followers int, updateTime time.Time) {
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

	insert_statement := fmt.Sprintf(`INSERT INTO Statistics (id, date, follows, followers)
																	VALUES (null, ?, ?, ?)`)
	stmtIns, err := db.Prepare(insert_statement)

	if err != nil {
		fmt.Println(err)
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(updateTime, follows, followers)
}

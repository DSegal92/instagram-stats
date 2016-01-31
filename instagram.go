package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var ACCESS_TOKEN = os.Getenv("INSTAGRAM_ACCESS_KEY")

type Follows struct {
	Pagination struct {
		NextCursor string `json:"next_cursor"`
	}
	Data []struct {
		Username string
	}
}

func collectUsers(relation string) []string {
	cursor := ""
	users := make([]string, 0)

	for {
		url := fmt.Sprintf("https://api.instagram.com/v1/users/self/%v?access_token=%v&cursor=%v", relation, ACCESS_TOKEN, cursor)
		content := getContent(url)

		var record Follows
		err := json.Unmarshal(content, &record)
		if err != nil {
			fmt.Println(err)
		}

		for i := 0; i < len(record.Data); i++ {
			users = append(users, record.Data[i].Username)
		}

		if record.Pagination.NextCursor == "" {
			break
		}
		cursor = record.Pagination.NextCursor
	}

	return users
}

func main() {
	updateTime := time.Now()

	follows := collectUsers("follows")
	insertRelations("follows", follows, updateTime)

	followers := collectUsers("followed-by")
	insertRelations("followers", followers, updateTime)
}

func getContent(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	return body
}

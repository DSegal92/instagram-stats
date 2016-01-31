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

func main() {
	cursor := ""

	follows := make([]string, 0)

	for {
		url := fmt.Sprintf("https://api.instagram.com/v1/users/self/follows?access_token=%v&cursor=%v", ACCESS_TOKEN, cursor)
		content := getContent(url)

		var record Follows
		err := json.Unmarshal(content, &record)
		if err != nil {
			fmt.Println(err)
		}

		for i := 0; i < len(record.Data); i++ {
			follows = append(follows, record.Data[i].Username)
		}

		if record.Pagination.NextCursor == "" {
			break
		}
		cursor = record.Pagination.NextCursor
	}

	insertFollows(follows)
	// for i := 0; i < len(follows); i++ {
	// 	fmt.Println(follows[i])
	// }
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

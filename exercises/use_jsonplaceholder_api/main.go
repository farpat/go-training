package use_jsonplaceholder_api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type Post struct {
	Id     int    `json:"id"`
	UserId int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func UpdateFirstPost() error {
	firstPost, err := RetrieveFirstPost()

	if err != nil {
		return err
	}

	firstPost.Title = firstPost.Title + " (updated)"
	body, err := json.Marshal(firstPost)
	if err != nil {
		return err
	}
	_, err = http.NewRequest(http.MethodPut, "https://jsonplaceholder.typicode.com/posts/"+strconv.Itoa(firstPost.Id), bytes.NewReader(body))
	if err != nil {
		return err
	}

	return nil
}

func RetrieveFirstPost() (Post, error) {
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		return Post{}, err
	}

	defer response.Body.Close()

	var posts []Post
	err = json.NewDecoder(response.Body).Decode(&posts)

	if err != nil {
		return Post{}, err
	}

	if len(posts) == 0 {
		return Post{}, errors.New("no posts found")
	}

	return posts[0], nil
}

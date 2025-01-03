package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type JokeResponse struct {
	Category  []string `json:"category"`
	CreatedAt string   `json:"created_at"`
	IconUrl   string   `json:"icon_url"`
	Id        string   `json:"id"`
	UpdatedAt string   `json:"updated_at"`
	Url       string   `json:"url"`
	Value     string   `json:"value"`
}

func main() {
	url := "https://api.chucknorris.io/jokes/random"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var joke JokeResponse
	err = json.Unmarshal(body, &joke)
	if err != nil {
		panic(err)
	}

	fmt.Println(joke.Value)
}

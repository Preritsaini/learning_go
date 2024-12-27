package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string
	Price    int
	Platform string
	Password string
	Tags     []string
}

func main() {
	//EncodeJson()
	decodeJson()
}
func EncodeJson() {
	lcoCources := []course{
		{"Reactjs bootcamp", 299, "learcode", "abc123", []string{"web dev", "Js"}},
		{"React bootcamp", 299, "learcode", "abc123", []string{"web dev", "Js"}},
		{"js bootcamp", 299, "learcode", "abc123", []string{"web dev", "Js"}},
	}
	finalJson, err := json.MarshalIndent(lcoCources, "", "\t")
	if err != nil {
		panic(err)

	}
	fmt.Printf(string(finalJson))
}
func decodeJson() {
	jsonData := []byte(`
		{
			"userId": 1,
			"id": 1,
			"title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
			"body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"
		}
	`)

	var post Post
	err := json.Unmarshal(jsonData, &post)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Printf("Decoded JSON:\n%+v\n", post)
}

package main

import (
	"encoding/json"
	"fmt"
)

type todo struct {
	ID     int    `json:"id"`
	Title  string `json:"name"`
	Status string `json:"status"`
}

func main() {
	data := []byte(`{
		"id": 2,
		"name": "buy new ipad",
		"status": "Active"
	}`)

	var t todo
	err := json.Unmarshal(data, &t)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%#v\n", t)
}

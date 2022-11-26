package main

import (
	"encoding/json"
	"fmt"
)

// struct todo
type todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

func main() {
	t := todo{
		ID:     0,
		Title:  "buy new ipad",
		Status: "active",
	}

	b, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("% #v\n", string(b))
}

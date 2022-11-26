package main

import (
	"fmt"
	"net/http"
)

// todo
/*
{
		"id": 0,
    "title": "buy new ipad",
    "status": "active"
  }
*/

var id int
var title string
var status string

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		body := []byte(s)
		w.WriteHeader(http.StatusOK)
		w.Write(body)
		return
	}

	if r.Method == "POST" {
		body := []byte(`{"message": "AnuchitO"}`)
		w.WriteHeader(http.StatusCreated)
		w.Write(body)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func main() {
	fmt.Println("start..")
	http.HandleFunc("/hello", handler)
	fmt.Println("server..")
	err := http.ListenAndServe("localhost:2565", nil)
	if err != nil {
		fmt.Println("error..", err)
		return
	}
	fmt.Println("end...")
}

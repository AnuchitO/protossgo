package main

import (
	"encoding/json"
	"fmt"
	"io"
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
type todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var t1 = todo{ID: 1, Title: "buy new ipad", Status: "active"}
var t2 = todo{ID: 2, Title: "hello new ipad", Status: "active"}

var todos = []todo{t1, t2}

func handler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		body, err := json.Marshal(todos)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(body)
		return
	}

	if req.Method == "POST" {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		var t todo
		err = json.Unmarshal(body, &t)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		todos = append(todos, t)
		w.WriteHeader(http.StatusCreated)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func main() {
	fmt.Println("start..")
	http.HandleFunc("/todos", handler)
	fmt.Println("server..")
	err := http.ListenAndServe("localhost:1323", nil)
	if err != nil {
		fmt.Println("error..", err)
		return
	}
	fmt.Println("end...")
}

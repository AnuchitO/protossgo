package todo

type T struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var Todos = []T{
	{ID: 1, Title: "Todo 1", Status: "active"},
}

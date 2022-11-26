package main

import (
	"fmt"

	"github.com/AnuchitO/apiecho/todo"
	"github.com/labstack/echo/v4"
)

func serve() *echo.Echo {
	e := echo.New()
	e.GET("/todos", todo.GetTodos)
	e.POST("/todos", todo.CreateTodos)
	return e
}

func main() {
	e := serve()
	err := e.Start(":1323")
	fmt.Println(err)
}

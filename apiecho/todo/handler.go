package todo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetTodos(c echo.Context) error {
	return c.JSON(http.StatusOK, Todos)
}

func CreateTodos(c echo.Context) error {
	var t T
	if err := c.Bind(&t); err != nil {
		return err
	}
	Todos = append(Todos, t)
	return c.JSON(http.StatusCreated, t)
}

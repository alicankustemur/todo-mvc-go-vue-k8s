package main

import (
	"fmt"
	"github.com/labstack/echo"
)

type ( 
	Todo struct {
		Id        int    `json:"id"`
		Title     string `json:"title"`
		Order     int    `json:"order"`
		Completed bool   `json:"completed"`
		Url       string `json:"url"`
	}
	handler struct {
		db map[string]*Todo
	}
)

func (t *Todo) SetUrl(c echo.Context) {
	scheme := "http"
	if c.IsTLS() {
		scheme = "https"
	}
	t.Url = fmt.Sprintf("%v://%v/todos/%v", scheme, c.Request().Host, t.Id)
}
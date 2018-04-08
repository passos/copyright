package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func rootHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func pingHandler(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

func getAccount(c echo.Context) error {
	// TODO: run select SQL
	account := &Account{

	}
	if err := c.Bind(account); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, account)
}

func createAccount(c echo.Context) error {
	// TODO: run insert SQL
	account := &Account{

	}
	if err := c.Bind(account); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, account)
}

func updateAccount(c echo.Context) error {
	// TODO: run update SQL
	account := &Account{

	}
	if err := c.Bind(account); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, account)
}

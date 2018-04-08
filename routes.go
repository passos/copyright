package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func pingHandler(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

type (
	Account struct {
		ID         int    `json:"id"`
		Email      string `json:"email"`
		Username   string `json:"username"`
		IdentityID string `json:"identity_id"`
	}
)

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

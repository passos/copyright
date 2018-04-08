package main

import (
	"net/http"

	"fmt"

	"github.com/labstack/echo"
	_ "github.com/labstack/gommon/log"
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
	account := &Account{}
	if err := c.Bind(account); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, account)
}

//curl -d "email=yekai@sohu.com&username=yekai&identity_id=123" "http://localhost:8086"
//curl -l -H "Content-type: application/json" -X POST -d '{"email":"yekai@sohu.com","username":"test","identity_id":"123"}'

func createAccount(c echo.Context) error {
	fmt.Println("createAccount is called")
	//log.Debugf("%s", "createAccount is called")
	c.Logger().Debug("run call createAccount")
	// TODO: run insert SQL
	account := &Account{}
	if err := c.Bind(account); err != nil {
		return err
	}
	fmt.Printf("%+v", account)
	_, err := account.AddAccount()
	if err != nil {
		fmt.Println("run add account err", err)
		return err
	}
	return c.JSON(http.StatusCreated, account)
}

func updateAccount(c echo.Context) error {
	// TODO: run QUERY SQL
	id := c.Param("id")
	fmt.Println("id===", id)
	account := &Account{}
	if err := c.Bind(account); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, account)
}

//查看账户信息 url=/account/:id method=get
func queryAccount(c echo.Context) error {
	// TODO: run QUERY SQL
	fmt.Println("queryAccount is called")
	id := c.Param("id")
	fmt.Println("id===", id)
	accinfo := query("select * from account where account_id=" + id)
	fmt.Println(accinfo[0])
	account := accinfo[0] //:= &Account{}
	return c.JSON(http.StatusCreated, account)
}

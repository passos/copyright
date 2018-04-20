package main

import (
	"fmt"

	_ "github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	_ "github.com/labstack/gommon/log"
)

type Vote struct {
	AccountID    int    `json:"account_id"`
	Pass         string `json:"pass"`
	Content_hash string `json:"content_hash"`
	Comment      string `json:"comment"`
}

//投票处理
func Voting(c echo.Context) error {
	// TODO: run insert sql
	var resp Resp
	resp.Errno = RECODE_OK
	resp.ErrMsg = RecodeText(resp.Errno)
	defer ReturnData(c, &resp)
	vote := &Vote{}
	//获得参数
	err := c.Bind(vote)
	if err != nil {
		resp.Errno = RECODE_DATAERR
		resp.ErrMsg = RecodeText(resp.Errno)
		return err
	}
	//判断图片hash是否有效
	if vote.Content_hash == "" {
		resp.Errno = RECODE_DATAERR
		resp.ErrMsg = RecodeText(resp.Errno)
		return err
	}
	sess, _ := session.Get("session", c)
	accid, ok := sess.Values["account_id"].(int)
	fromaddr, ok := sess.Values["address"].(string)
	if !ok {
		fmt.Println("session not exists", ok)
		resp.Errno = RECODE_SESSIONERR
		resp.ErrMsg = RecodeText(resp.Errno)
		return nil
	}
	vote.AccountID = accid
	if vote.AddVote() != nil {
		resp.Errno = RECODE_DBERR
		resp.ErrMsg = RecodeText(resp.Errno)
		return err
	}

	go transfer20(vote.Pass, fromaddr, config.Eth.MgrAddress, int64(100))
	return nil
}

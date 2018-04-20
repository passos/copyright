package main

import (
	"crypto/sha256"
	"io"
	_ "math/big"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"fmt"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	_ "github.com/labstack/gommon/log"
)

//用户图片显示一页最大数量
var PAGE_MAX_PICTURES int = 5

type (
	Account struct {
		ID         int    `json:"id"`
		Email      string `json:"email"`
		Username   string `json:"username"`
		IdentityID string `json:"identity_id"`
		Address    string `json:"address"`
	}
	Content struct {
		ContentID int `json:"content_id"`
		//AccountID    int       `json:"account_id"`
		Title        string    `json:"title"`
		Content      []byte    `json:"content"`
		Content_hash string    `json:"content_hash"`
		Ts           time.Time `json:"ts"`
	}
	Aution struct {
		AccountID    int       `json:"account_id"`
		ContentID    int       `json:"content_id"`
		Content_hash string    `json:"content_hash"`
		Percent      int       `json:"percent"`
		Price        int       `json:"price"`
		SellPrice    int       `json:"sell_price"`
		SellPercent  int       `json:"sell_percent"`
		Ts           time.Time `json:"ts"`
		Status       int       `json:"status"`
		Address      string    `json:"address"`
		Pass         string    `json:"pass"`
		TokenID      int       `json:"tokenid"`
	}
	Trade struct {
		AccountID int `json:"account_id"`
		Amount    int `json:"amount"` //可以加或者减
		Percent   int `json:"percent"`
	}
	Resp struct {
		Errno  string      `json:"errno"`
		ErrMsg string      `json:"errmsg"`
		Data   interface{} `json:"data"`
	}
)

func pingHandler(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

//静态页面处理
func staticHandler(c echo.Context) error {
	urlpath := c.Request().URL.Path
	//fmt.Println()
	if strings.Index(urlpath, "api") > 0 {
		return nil
	}
	if strings.Index(urlpath, "static") > 0 {
		rs := []rune(urlpath)
		urlpath = string(rs[1:])
		http.ServeFile(c.Response(), c.Request(), urlpath)
		return nil
	}
	if strings.Index(urlpath, "html") > 0 {
		urlpath = "/home" + urlpath
	}
	if urlpath == "/" {
		urlpath = "/home/index.html"
	}

	http.ServeFile(c.Response(), c.Request(), "static/pc"+urlpath)
	return c.NoContent(http.StatusOK)
}

//静态页面处理
func staticUserHandler(c echo.Context) error {
	urlpath := c.Request().URL.Path
	//fmt.Println()
	if strings.Index(urlpath, "api") > 0 {
		return nil
	}
	if strings.Index(urlpath, "static") > 0 {
		rs := []rune(urlpath)
		urlpath = string(rs[1:])
		http.ServeFile(c.Response(), c.Request(), urlpath)
		return nil
	}

	http.ServeFile(c.Response(), c.Request(), "static/pc"+urlpath)
	return c.NoContent(http.StatusOK)
}

//resp数据响应
func ReturnData(c echo.Context, resp interface{}) {
	c.JSON(http.StatusOK, resp)
}

//登陆验证,成功会返回account_id
func getAccount(c echo.Context) error {
	// TODO: run select SQL
	var resp Resp
	resp.Errno = RECODE_OK
	resp.ErrMsg = RecodeText(resp.Errno)
	defer ReturnData(c, &resp)
	account := &Account{}
	if err := c.Bind(account); err != nil {
		resp.Errno = RECODE_DATAERR
		resp.ErrMsg = RecodeText(resp.Errno)
		return err
	}
	pass := account.IdentityID
	account.IdentityID = GetMd5(account.IdentityID)
	fmt.Println("user=", account.Username, "pass=", account.IdentityID)
	sql := "select * from account where username='" + account.Username + "' and identity_id='" + account.IdentityID + "'"
	fmt.Println("sql==", sql)
	rows, idx, err := query(sql)
	if idx > 0 {
		account.Email = rows[0]["email"]
		account.ID, err = strconv.Atoi(rows[0]["account_id"])
		account.Address = rows[0]["address"]
	} else {
		fmt.Println("user or password err", err)
		resp.Errno = RECODE_USERERR
		resp.ErrMsg = RecodeText(resp.Errno)
		return err
	}
	//设置session
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   300,
		HttpOnly: true,
	}
	sess.Values["account_id"] = account.ID
	sess.Values["username"] = account.Username
	sess.Values["address"] = account.Address
	sess.Values["pass"] = pass
	sess.Save(c.Request(), c.Response())

	fmt.Println(sess.Values)
	mapAcc := make(map[string]interface{})
	mapAcc["account_id"] = account.ID
	mapAcc["username"] = account.Username
	mapAcc["address"] = account.Address
	resp.Data = mapAcc
	//return c.JSON(http.StatusOK, account)
	return nil
}

// curl -H "Content-type: application/json" -X POST -d '{"email":"yekai@sohu.com","username":"yekai","identity_id":"123"}' "http://localhost:8086/account"
// curl -H "Content-type: application/json" -X POST -d '{"username":"test","identity_id":"123"}' "http://localhost:8086/login"
// curl --form "image=@./etc/wyz.jpg" --form "account_id=1" --form "title=wyz.jpg" http://localhost:8086/content
// curl -X GET "http://localhost:8086/content/e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"

func createAccount(c echo.Context) error {
	var resp Resp
	resp.Errno = RECODE_OK
	resp.ErrMsg = RecodeText(resp.Errno)
	defer ReturnData(c, &resp)
	fmt.Println("createAccount is called")
	//log.Debugf("%s", "createAccount is called")
	c.Logger().Debug("run call createAccount")
	// TODO: run insert SQL
	account := &Account{}

	if err := c.Bind(account); err != nil {
		resp.Errno = RECODE_DATAERR
		resp.ErrMsg = RecodeText(resp.Errno)
		return err
	}
	fmt.Printf("%+v\n", account)
	address, err := GetAccAddress(account.IdentityID, config.Eth.Rpc)
	account.Address = address
	if err != nil {
		fmt.Println("createAccount:get address err", err)
		resp.Errno = RECODE_THIRDERR
		resp.ErrMsg = RecodeText(resp.Errno)
		return err
	}
	//	pass := account.IdentityID //为了之后存入session
	account.IdentityID = GetMd5(account.IdentityID)
	_, err = account.AddAccount()
	if err != nil {
		fmt.Println("run add account err", err)
		resp.Errno = RECODE_DBERR
		resp.ErrMsg = RecodeText(resp.Errno)
		return err
	}

	//设置session
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   300,
		HttpOnly: true,
	}
	sess.Values["account_id"] = account.ID
	sess.Values["username"] = account.Username
	sess.Values["address"] = account.Address
	//	sess.Values["pass"] = pass
	sess.Save(c.Request(), c.Response())
	fmt.Println(sess.Values)
	mapAcc := make(map[string]interface{})
	mapAcc["account_id"] = account.ID
	mapAcc["username"] = account.Username
	mapAcc["address"] = account.Address
	resp.Data = mapAcc
	//调用初始化账户pixc
	go InitAccToken(account.Address)

	return nil //c.JSON(http.StatusCreated, account)
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
	//	sess, _ := session.Get("session", c)
	//	fmt.Println(sess.Values)
	accinfo, _, err := query("select * from account where account_id=" + id)
	if err != nil {
		return err
	}
	fmt.Println(accinfo[0])
	account := accinfo[0] //:= &Account{}
	return c.JSON(http.StatusCreated, account)
}

//上传图片
//curl --form "fileupload=@wyz.jpg" http://localhost:8086/content
func UploadContent(c echo.Context) error {
	// TODO: run Insert SQL + call contract

	var resp Resp
	resp.Errno = RECODE_OK
	resp.ErrMsg = RecodeText(resp.Errno)
	defer ReturnData(c, &resp)

	fmt.Println("UploadContent is called")
	content := &Content{}
	//fmt.Println(c.FormValue("account_id"))
	//val, _ := c.FormParams()
	//fmt.Println(val.Get("account_id"))
	h, err := c.FormFile("fileName")
	if err != nil {
		fmt.Println("Formfile err", err)
		resp.Errno = RECODE_PARAMERR
		resp.ErrMsg = RecodeText(resp.Errno)
		return err
	}

	fmt.Println(h.Filename, h.Size)
	file, err := h.Open()
	content.Content = make([]byte, h.Size)
	num, err := file.Read(content.Content)
	if err != nil {
		fmt.Println("read file err")
		resp.Errno = RECODE_DATAERR
		resp.ErrMsg = RecodeText(resp.Errno)
		return err
	}
	fmt.Println("read file num ===", num, "size==", h.Size)
	defer file.Close()
	file.Seek(0, os.SEEK_SET) //调整到文件头

	haha := sha256.New()

	_, err = io.Copy(haha, file)
	if err != nil {
		fmt.Println("hash-copy file err", err)
		resp.Errno = RECODE_HASHERR
		resp.ErrMsg = RecodeText(resp.Errno)
		return err
	}
	content.Content_hash = fmt.Sprintf("%x", haha.Sum(nil))
	fmt.Printf("gethash....==%x\n", haha.Sum(nil))

	//从session中获取account_id
	sess, _ := session.Get("session", c)
	accid, ok := sess.Values["account_id"].(int)
	if !ok {
		fmt.Println("session not exists", ok)
		resp.Errno = RECODE_SESSIONERR
		resp.ErrMsg = RecodeText(resp.Errno)
		return nil
	}
	address, ok := sess.Values["address"].(string)
	pass, ok := sess.Values["pass"].(string)
	//content.AccountID = sess.Values["account_id"].(int)
	//fmt.Println("account_id===", content.AccountID)

	content.Title = h.Filename //c.FormValue("title")
	//contentsql := fmt.Sprintf("insert into content(account_id,title,content,content_hash) values(%d,'%s',)")
	qrySql := "select title,content_hash from content where content_hash='" + content.Content_hash + "'"
	_, n, err := query(qrySql) //确认是否存在该hash
	if err != nil {
		fmt.Println("query content err")
		resp.Errno = RECODE_DBERR
		resp.ErrMsg = RecodeText(resp.Errno)
		return err
	}
	if n > 0 {
		fmt.Println("upload file is exists")
		resp.Errno = RECODE_EXISTSERR
		resp.ErrMsg = RecodeText(resp.Errno)
		return err
	}
	//fmt.Println("content==", content)
	id, err := content.AddContent()
	if err != nil {
		//fmt.Println("hash-copy file err", err)
		resp.Errno = RECODE_DBERR
		resp.ErrMsg = RecodeText(resp.Errno)
		return err
	}
	//还需要添加用户和资产对应关系表
	aution := &Aution{}
	aution.AccountID = accid
	aution.Content_hash = content.Content_hash
	aution.ContentID = int(id)
	aution.Percent = 100
	_, err = aution.AddAccountContent()
	if err != nil {
		fmt.Println("AddAccountContent err", err)
		resp.Errno = RECODE_DBERR
		resp.ErrMsg = RecodeText(resp.Errno)
		return err
	}

	content.ContentID = int(id)
	mapAcc := make(map[string]interface{})
	mapAcc["content_id"] = content.ContentID
	mapAcc["title"] = content.Title
	resp.Data = mapAcc
	address = string([]rune(address)[2:])
	go Pic721Token(aution.Content_hash, address, config.Eth.Contract721, pass)
	return nil
}

//查询图片信息 - 根据图片hash
func queryContent(c echo.Context) error {
	//	var resp Resp
	//	resp.Errno = RECODE_OK
	//	resp.ErrMsg = RecodeText(resp.Errno)
	//	defer ReturnData(c, &resp)

	dna := c.Param("dna")
	fmt.Println("queryContent is called,dna=", dna)
	//sql := fmt.Sprintf("select * from content where content_hash='%s'", dna)
	//fmt.Println("sql is ", sql)
	//m, n, err := query(sql)

	content := &Content{}
	content.Content_hash = dna
	err := content.getContent()
	if err != nil {
		//		resp.Errno = RECODE_NODATA
		//		resp.ErrMsg = RecodeText(resp.Errno)
		fmt.Println("queryContent query err", err)
		return err
	}
	//content := m[0]

	//发送对应文件块
	if err := SendFile(c, content.Content, content.Title); err != nil {
		//		resp.Errno = RECODE_IOERR
		//		resp.ErrMsg = RecodeText(resp.Errno)
		return err
	}
	return nil

}

//处理图片获取
func SendFile(c echo.Context, data []byte, filename string) error {

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file err")
		return nil
	}
	defer os.Remove(filename)
	num, err := file.Write(data)
	if num <= 0 || err != nil {
		fmt.Println("data write err", err)
		return err
	}
	fmt.Println("sendfile is called,num==", num, "err==", err)
	file.Close()
	http.ServeFile(c.Response(), c.Request(), filename)
	return nil
}

//判断用户是否登录
func getAccountID(c echo.Context) error {
	//通过session来判断是否可以登陆
	var resp Resp
	resp.Errno = RECODE_LOGINERR
	resp.ErrMsg = RecodeText(resp.Errno)
	defer ReturnData(c, &resp)
	sess, _ := session.Get("session", c)
	accid := sess.Values["account_id"].(int)
	if accid <= 0 {
		resp.Errno = RECODE_LOGINERR
		resp.ErrMsg = RecodeText(resp.Errno)
		return nil
	}
	resp.Errno = RECODE_OK
	resp.ErrMsg = RecodeText(resp.Errno)
	mapResp := make(map[string]interface{})
	mapResp["account_id"] = accid
	mapResp["address"] = sess.Values["address"]
	resp.Data = mapResp
	return nil
}

//退出登陆
func AccountQuit(c echo.Context) error {
	//清除session信息
	var resp Resp
	resp.Errno = RECODE_OK
	resp.ErrMsg = RecodeText(resp.Errno)
	defer ReturnData(c, &resp)
	sess, _ := session.Get("session", c)
	sess.Values["account_id"] = 0
	return nil
}

//查看当前用户所有图片 -- 接口
func getAccountContent(c echo.Context) error {
	var resp Resp
	resp.Errno = RECODE_OK
	resp.ErrMsg = RecodeText(resp.Errno)
	defer ReturnData(c, &resp)
	//查询session
	//	sess, _ := session.Get("session", c)
	//	accid := sess.Values["account_id"].(int)

	sess, _ := session.Get("session", c)
	accid, ok := sess.Values["account_id"].(int)
	if !ok {
		fmt.Println("session not exists", ok)
		resp.Errno = RECODE_SESSIONERR
		resp.ErrMsg = RecodeText(resp.Errno)
		return nil
	}

	//content.AccountID = accid
	fmt.Println("account_id===", accid)
	sql := fmt.Sprintf("select b.content_id,b.account_id,a.title,a.content_hash,b.ts,b.status from content a,account_content b where a.content_hash=b.content_hash and b.account_id=%d order by b.ts desc", accid)
	num, n, err := query(sql)
	if err != nil {
		fmt.Println("query user content err", err)
		resp.Errno = RECODE_DBERR
		resp.ErrMsg = RecodeText(resp.Errno)
		return err
	}

	total_page := int(n)/PAGE_MAX_PICTURES + 1
	current_page := 1
	mapResp := make(map[string]interface{})
	mapResp["total_page"] = total_page
	mapResp["current_page"] = current_page
	contents := make([]interface{}, 1)

	for k, v := range num {
		if k == 0 {
			contents[0] = v
		} else {
			contents = append(contents, v)
		}

	}
	mapResp["contents"] = contents
	resp.Data = mapResp

	return nil
}

//发布拍卖信息 post url=/aution  携带json信息
//account_content status=1
func AutionContent(c echo.Context) error {
	var resp Resp
	resp.Errno = RECODE_OK
	resp.ErrMsg = RecodeText(resp.Errno)
	defer ReturnData(c, &resp)
	aution := &Aution{}
	sess, _ := session.Get("session", c)
	accid, ok := sess.Values["account_id"].(int)
	if !ok {
		fmt.Println("session not exists", ok)
		resp.Errno = RECODE_SESSIONERR
		resp.ErrMsg = RecodeText(resp.Errno)
		return nil
	}
	aution.AccountID = accid

	//获取请求json数据
	if err := c.Bind(aution); err != nil {
		resp.Errno = RECODE_DATAERR
		resp.ErrMsg = RecodeText(resp.Errno)
		return err
	}
	fmt.Println("aution====", aution)
	if aution.Content_hash == "" {
		fmt.Println("request data err")
		resp.Errno = RECODE_DATAERR
		resp.ErrMsg = RecodeText(resp.Errno)
		return nil
	}
	//fmt.Println(aution)
	//拍卖资产肯定已经存在，直接修改即可
	sql := fmt.Sprintf("update account_content set sell_price=%d,sell_percent=%d,status=1 where account_id=%d and content_hash='%s'",
		aution.SellPrice, aution.SellPercent, aution.AccountID, aution.Content_hash)

	_, err := Create(sql)
	if err != nil {
		fmt.Println("update account_content err", err, "sql is ", sql)
		resp.Errno = RECODE_DBERR
		resp.ErrMsg = RecodeText(resp.Errno)
		return err
	}
	return nil
}

//查看所有拍卖中的图片
func GetAutions(c echo.Context) error {
	var resp Resp
	resp.Errno = RECODE_OK
	resp.ErrMsg = RecodeText(resp.Errno)
	defer ReturnData(c, &resp)
	//aution := &Aution{}
	//aution := &Aution{}
	sess, _ := session.Get("session", c)
	accid, ok := sess.Values["account_id"].(int)
	if !ok {
		fmt.Println("session not exists", ok)
		resp.Errno = RECODE_SESSIONERR
		resp.ErrMsg = RecodeText(resp.Errno)
		return nil
	}
	//aution.AccountID = accid

	//显示一个列表
	sql := "select a.content_hash,a.percent,a.status,a.sell_price,a.sell_percent,a.account_id from account_content a,content b where a.content_id=b.content_id and  a.status='1'"
	m, n, err := query(sql)
	if err != nil {
		fmt.Println("query account aution info err", err)
		resp.Errno = RECODE_DBERR
		resp.ErrMsg = RecodeText(resp.Errno)
		return err
	}
	total_page := int(n)/PAGE_MAX_PICTURES + 1
	current_page := 1
	mapResp := make(map[string]interface{})
	mapResp["total_page"] = total_page
	mapResp["current_page"] = current_page
	contents := make([]interface{}, 1)
	for k, v := range m {
		fmt.Println("get acc :", v["account_id"], "sessacc:", string(accid))
		if v["account_id"] == string(accid) {
			v["is_self"] = "true"
		} else {
			v["is_self"] = "false"
		}

		if k == 0 {
			contents[0] = v
		} else {
			contents = append(contents, v)
		}
	}
	mapResp["contents"] = contents
	mapResp["account_id"] = accid
	resp.Data = mapResp
	return nil
}

//购买所有权
//put url=/aution json={hash:xx,price:xx,percent:yy,account_id:1,ts:xx}

func AutionBuy(c echo.Context) error {
	var resp Resp
	resp.Errno = RECODE_OK
	resp.ErrMsg = RecodeText(resp.Errno)
	defer ReturnData(c, &resp)
	aution := &Aution{}
	//获取请求json数据
	if err := c.Bind(aution); err != nil {
		resp.Errno = RECODE_DATAERR
		resp.ErrMsg = RecodeText(resp.Errno)
		return err
	}
	sess, _ := session.Get("session", c)
	accid, ok := sess.Values["account_id"].(int)
	if !ok {
		fmt.Println("session not exists", ok)
		resp.Errno = RECODE_SESSIONERR
		resp.ErrMsg = RecodeText(resp.Errno)
		return nil
	}
	aution.AccountID = accid
	//	sess, _ := session.Get("session", c)
	//	//	from_acc := aution.AccountID // 卖方的account
	//	aution.AccountID = sess.Values["account_id"].(int)

	//需要判断智能合约账户金额是否还够购买
	//记录用户请求拍卖信息，之后会按规定的时间计算拍卖结果
	//	end_ts := aution.Ts.Add(time.Second * 86400)
	sql := fmt.Sprintf("insert into aution(content_hash,account_id,percent,price) values('%s',%d,%d,%d)",
		aution.Content_hash, aution.AccountID, aution.Percent, aution.Price)
	_, err := Create(sql)
	if err != nil {
		fmt.Println("add aution info err", err, "sql:", sql)
		resp.Errno = RECODE_DBERR
		resp.ErrMsg = RecodeText(resp.Errno)
		return err
	}
	return nil
}

//获得账户余额
func AccGetBalance(c echo.Context) error {
	var resp Resp
	resp.Errno = RECODE_OK
	resp.ErrMsg = RecodeText(resp.Errno)
	defer ReturnData(c, &resp)
	address := c.Param("address")
	balance, err := GetBalanceOf(address)
	if err != nil {
		fmt.Println("get balance err", err)
		resp.Errno = RECODE_IPCERR
		resp.ErrMsg = RecodeText(resp.Errno)
		return err
	}
	mapResp := make(map[string]interface{})
	mapResp["balance"] = balance
	resp.Data = mapResp
	return err
}

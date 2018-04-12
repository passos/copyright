package main

import (
	"fmt"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

var (
	Version   = "1.0.0"
	Commit    = ""
	BuildTime = ""
)

func installMiddleWare(e *echo.Echo, config *ServerConfig) {
	l := log.New("middleware")
	l.SetLevel(log.Level())
	l.SetHeader(config.Common.LogFormat)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	//e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
	//	TokenLookup: "header:X-XSRF-TOKEN",
	//}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://loader.cc", "https://loader.cc"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))

	// e.Use(middleware.KeyAuth(func(key string, context echo.Context) (error, bool) {
	// 	l.Debugf("Key: %s", key)
	// 	return nil, true
	// }))
}

var e *echo.Echo
var config *ServerConfig

func main() {
	config = getConfig()
	if config == nil {
		return
	}
	fmt.Println("get config ", config)
	fmt.Println("rpc==", config.Eth)
	db = initDB(config)

	log.SetLevel(log.DEBUG)
	log.SetHeader(config.Common.LogFormat)

	e = echo.New()
	installMiddleWare(e, config)

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	//静态页面设置
	e.GET("/*", staticHandler)
	e.GET("/user*", staticUserHandler)

	//路由设置
	e.GET("/ping", pingHandler)

	e.POST("/account", createAccount)
	e.POST("/login", getAccount)
	//	e.POST("/account/:id", updateAccount)
	e.GET("/account/:id", queryAccount)
	//上传图片
	e.POST("/content", UploadContent)
	//获取用户全部图片
	e.GET("/content", getAccountContent)
	//根据hash获取指定图片内容
	e.GET("/content/:dna", queryContent)
	//session处理
	e.GET("/session", getAccountID)
	e.DELETE("/session", AccountQuit)
	//图片交易
	e.POST("/aution", AutionContent)
	e.PUT("/aution", AutionBuy)
	e.GET("/aution", GetAutions)
	//查询账户pixc余额
	e.GET("/account/:address", AccGetBalance)
	//启动定时任务，每晚自动运行检查合约
	go runAtTime(time.Date(2018, 4, 11, 23, 59, 59, 0, time.Local))

	//go runAtTime(time.Date(2018, 4, 11, 16, 10, 0, 0, time.Local))
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Common.Port)))
}

package main

import (
	"fmt"
	"github.com/labstack/echo"
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

func main() {
	config := getConfig()
	if config == nil {
		return
	}

	log.SetLevel(log.DEBUG)
	log.SetHeader(config.Common.LogFormat)

	e := echo.New()
	installMiddleWare(e, config)

	e.GET("/", rootHandler)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Common.Port)))
}

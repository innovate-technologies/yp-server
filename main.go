package main

import (
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo"
)

type config struct {
	Port string `default:"80"`
}

var conf config

func main() {
	err := envconfig.Process("ypserver", &conf)
	if err != nil {
		log.Fatal(err.Error())
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Any("/icecast", icecastHandle)
	e.Logger.Fatal(e.Start(":" + conf.Port))
}

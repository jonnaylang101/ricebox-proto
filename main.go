package main

import (
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/labstack/echo/v4"
)

var (
	path = "/composer"
	dst  = "./web/dst"
)

func main() {
	e := echo.New()

	rb := rice.MustFindBox("./web/dst")

	grp := e.Group(path)

	assetHandler := http.StripPrefix(path, http.FileServer(rb.HTTPBox()))

	grp.GET("/_status", getVersion())
	grp.GET("", loadIndex(rb))
	grp.GET("/static/js/*", echo.WrapHandler(assetHandler))
	grp.Any("*", echo.WrapHandler(assetHandler))

	e.Start(":8080")
}

func loadIndex(rb *rice.Box) echo.HandlerFunc {
	return func(c echo.Context) error {

		bytes, err := rb.Bytes("index.html")
		if err != nil {
			return err
		}

		return c.HTMLBlob(http.StatusOK, bytes)
	}
}

func getVersion() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "v1.1.1")
	}
}

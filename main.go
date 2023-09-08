package main

import (
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/labstack/echo/v4"
)

var path = "/composer"

func main() {
	e := echo.New()

	rb := rice.MustFindBox("./web/dst")

	grp := e.Group("/composer")

	assetHandler := http.StripPrefix("/composer", http.FileServer(rb.HTTPBox()))

	grp.GET("", echo.WrapHandler(assetHandler))
	grp.GET("/static/js/*", echo.WrapHandler(assetHandler))
	// grp.Any("*", echo.WrapHandler(assetHandler))

	e.Start(":8080")
}

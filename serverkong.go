package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/login", func(c echo.Context) error {
		return c.String(http.StatusOK, "Login Succeeded")
	})
	e.Logger.Fatal(e.Start(":1337"))
}

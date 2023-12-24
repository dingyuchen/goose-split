package main

import (
	component "github.com/dingyuchen/goose-split/components"
	"github.com/labstack/echo/v5"
)

func index(c echo.Context) error {
	hi := component.Hello("world")
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
	return hi.Render(c.Request().Context(), c.Response())
}
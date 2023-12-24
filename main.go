package main

import (
	"log"

	"github.com/dingyuchen/goose-split/templates/components"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	// serves static files from the provided public dir (if exists)
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/", func(c echo.Context) error {
			cmp := components.Split(2)
			c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
			return cmp.Render(c.Request().Context(), c.Response())
		})
		return nil
	})

	app.RootCmd.Execute()
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

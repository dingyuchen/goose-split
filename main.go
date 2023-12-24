package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/dingyuchen/goose-split/views"
	"github.com/dingyuchen/goose-split/views/components"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func main() {
	app := pocketbase.New()

	// loosely check if it was executed using "go run"
	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		// enable auto creation of migration files when making collection changes in the Admin UI
		// (the isGoRun check is to enable it only during development)
		Automigrate: isGoRun,
	})

	// serves static files from the provided public dir (if exists)
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/", func(c echo.Context) error {
			index := views.Index()
			c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
			return index.Render(c.Request().Context(), c.Response())
		})
		e.Router.GET("/new", func(c echo.Context) error {
			cmp := views.NewParty()
			c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
			return cmp.Render(c.Request().Context(), c.Response())
		})
		e.Router.GET("/new/:person", func(c echo.Context) error {
			personCount, err := strconv.Atoi(c.PathParam("person"))
			if err != nil {
				c.NoContent(400)
			}
			cmp := components.NextPerson(personCount)
			c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
			return cmp.Render(c.Request().Context(), c.Response())
		})
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

package controllers

import (
	"bytes"
	"context"
	"fmt"
	"hyperchan/models"
	"hyperchan/views"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Serve() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Static("/static", "static")

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<marquee>HYPERCHAN LIVES!</marquee>")
	})

	e.GET("/dbversion", func(c echo.Context) error {
		v := fmt.Sprintf("<marquee>%s</marquee>", models.DBVersion())
		return c.HTML(http.StatusOK, v)
	})

	e.GET("/templtest", func(c echo.Context) error {
		h := new(bytes.Buffer)
		views.Hello("HYPERCHANNERS").Render(context.Background(), h)
		return c.HTML(http.StatusOK, h.String())
	})

	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}

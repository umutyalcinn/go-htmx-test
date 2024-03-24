package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/umutyalcinn/erp-saas/internal/render"
)

type Count struct{
    Value int
}

func main(){
    e := echo.New()
    e.Renderer = render.NewRenderer()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.Static("/public", "public")

    e.GET("/", hello)

    i := 0

    e.POST("/increment", func(c echo.Context) error {
        i = i + 1
        c.Render(http.StatusOK, "count", &Count{ Value: i})

        return nil
    })

    e.Logger.Fatal(e.Start(":3131"))
}

func hello(c echo.Context) error {
    c.Render(http.StatusOK, "home", nil)

    return nil
}

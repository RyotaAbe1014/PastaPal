package main

import (
	"net/http"

	v1Router "github.com/RyotaAbe1014/Pastapal/internal/presentation/v1"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "True")
	})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// api/v1以下のルーティングを設定
	g := e.Group("/api/v1")
	v1Router.Router(g)

	e.Logger.Fatal(e.Start(":8000"))
}

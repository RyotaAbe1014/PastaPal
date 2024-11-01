package main

import (
	"net/http"

	v1Router "github.com/RyotaAbe1014/Pastapal/internal/presentation/v1"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	// settings.CorsMiddleware(e)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"}, // フロントエンドのURLを許可
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
	}))
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

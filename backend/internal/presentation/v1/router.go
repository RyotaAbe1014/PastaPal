package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Router(g *echo.Group) {
	// TODO: 必要なapiをとりあえず定義しているので、実装時に修正する
	// echo.Groupはネストすることができるので、この関数内でさらにGroupを作成してルーティングを定義することができる
	// ミドルウェアを設定する場合など、複雑になりそうであればモジュールを分割する

	// 食材種別管理
	g.POST("/ingredient-categories", func(c echo.Context) error {
		return c.String(http.StatusOK, "ingredient-categories")
	})
	g.GET(("/ingredient-categories"), func(c echo.Context) error {
		return c.String(http.StatusOK, "ingredient-categories")
	})
	g.GET("/ingredient-categories/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "ingredient-categories")
	})

	g.PUT("/ingredient-categories/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "ingredient-categories")
	})

	g.DELETE("/ingredient-categories/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "ingredient-categories")
	})

	// 食材管理
	g.GET(("/ingredients"), func(c echo.Context) error {
		return c.String(http.StatusOK, "ingredients")
	})
	g.GET(("/ingredients/:id"), func(c echo.Context) error {
		return c.String(http.StatusOK, "ingredients")
	})

	g.POST("/ingredients", func(c echo.Context) error {
		return c.String(http.StatusOK, "ingredients")
	})

	g.PUT("/ingredients/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "ingredients")
	})

	g.DELETE("/ingredients/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "ingredients")
	})

	// レシピ管理
	g.POST("/recipes", func(c echo.Context) error {
		return c.String(http.StatusOK, "resipes")
	})
	g.GET("/recipes", func(c echo.Context) error {
		return c.String(http.StatusOK, "resipes")
	})
	g.GET("/recipes/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "resipes")
	})
	g.PUT("/recipes/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "resipes")
	})
	g.DELETE("/recipes/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "resipes")
	})
}

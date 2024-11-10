package v1

import (
	"net/http"

	"github.com/RyotaAbe1014/Pastapal/internal/infrastructure/oauth/github"
	"github.com/labstack/echo/v4"
)

func Router(g *echo.Group) {
	g.GET("/auth/github", func(c echo.Context) error {
		// 認証URLを生成し、GitHubにリダイレクト
		authURL, err := github.GenerateAuthURL("state") // TODO: stateはCSRF対策のための値、実際にはランダムな値を設定する
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.Redirect(http.StatusTemporaryRedirect, authURL)
	})

	g.GET("/auth/github/callback", func(c echo.Context) error {
		// GitHubからリダイレクトされてきたときに呼び出される
		code := c.QueryParam("code")
		if code == "" {
			return c.String(http.StatusBadRequest, "code is empty")
		}

		// 認証コードをトークンに交換
		token, err := github.ExchangeCodeForToken(c.Request().Context(), code)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, token)
	})

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
